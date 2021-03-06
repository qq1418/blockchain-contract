package expressionutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	common0 "unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/data"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/function"
	"unicontract/src/core/engine/execengine/inf"
	"unicontract/src/core/engine/execengine/task"
	"unicontract/src/core/engine/gRPCClient"

	"gopkg.in/Knetic/govaluate.v2"
)

type ExpressionParseEngine struct {
	FunctionEngine *function.FunctionParseEngine
	Contract       inf.ICognitiveContract
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//类构造方法
func NewExpressionParseEngine() *ExpressionParseEngine {
	ep := &ExpressionParseEngine{}
	return ep
}

func (ep *ExpressionParseEngine) SetFunctionEngine(p_func_engine *function.FunctionParseEngine) {
	ep.FunctionEngine = p_func_engine
}

func (ep *ExpressionParseEngine) SetContract(p_contract inf.ICognitiveContract) {
	ep.Contract = p_contract
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
// 针对不同类型的表达式，解析表达式的值
// 表达式分为4大类：
//   1. 常量表达式   ExpressionType[Expression_Constant]
//   2. 变量表达式   ExpressionType[Expression_Variable]
//   3. 条件表达式   ExpressionType[Expression_Condition]
//   4. 函数表达式   ExpressionType[Expression_Function]
//   5. 决策表达式   ExpressionType[Expression_Candidate]
//Have read
func (ep *ExpressionParseEngine) EvaluateExpressionValue(p_exprtype string, p_expression string) (interface{}, error) {
	var v_return interface{} = nil
	var v_err error = nil
	//表达式类型不录入的话，需要动态识别表达式类型：常量、变量、函数
	if p_exprtype == "" || p_exprtype == constdef.ExpressionType[constdef.Expression_Unknown] {
		if ep.IsExprConst(p_expression) {
			v_return, v_err = ep.EvaluateExpressionConstant(p_expression)
		} else if ep.IsExprVariable(p_expression) {
			v_return, v_err = ep.EvaluateExpressionVariable(p_expression)
		} else if ep.IsExprFunction(p_expression) {
			v_func_return, v_err := ep.EvaluateExpressionFunction(p_expression)
			if v_err == nil {
				v_return = v_func_return.GetData()
			}
		}
		return v_return, v_err
	}
	if p_exprtype == constdef.ExpressionType[constdef.Expression_Constant] {
		v_return, v_err = ep.EvaluateExpressionConstant(p_expression)
	} else if p_exprtype == constdef.ExpressionType[constdef.Expression_Variable] {
		v_return, v_err = ep.EvaluateExpressionVariable(p_expression)
	} else if p_exprtype == constdef.ExpressionType[constdef.Expression_Condition] {
		v_return, v_err = ep.EvaluateExpressionCondition(p_expression)
	} else if p_exprtype == constdef.ExpressionType[constdef.Expression_Function] {
		v_return, v_err = ep.EvaluateExpressionFunction(p_expression)
	} else if p_exprtype == constdef.ExpressionType[constdef.Expression_Candidate] {
		v_return, v_err = ep.EvaluateExpressionCandidate(p_expression)
	}
	return v_return, v_err
}

//解析常量表达式，并求表达式的值
//常量表达式分类：
//   1.纯数字      => 直接返回该表达式值 int64
//   2.纯浮点数    => 直接返回该表达式值 float64
//   3.纯bool值   => 直接返回该表达式值 bool
//   4.纯字符串    => 直接返回该表达式值 string
//   5.纯日期串    => 转化为时间戳再返回 int64
//   6.纯数组串    => 转化为数组返回     []interface{}
//Have read
func (ep *ExpressionParseEngine) EvaluateExpressionConstant(p_expression string) (interface{}, error) {
	var v_return interface{} = nil
	var v_err error = nil

	if p_expression == "" {
		errMsg := "EvaluateExpressionConstant Param[p_expression] is nil!"
		uniledgerlog.Warn(errMsg)
		return "", fmt.Errorf(errMsg)
	}
	var v_classify string = ep.ParseExpressionClassify(p_expression)
	switch v_classify {
	case constdef.ExpressionClassify[constdef.Expr_Num]:
		v_return, v_err = ep.ParseExprNumValue(p_expression)
	case constdef.ExpressionClassify[constdef.Expr_Float]:
		v_return, v_err = ep.ParseExprFloatValue(p_expression)
	case constdef.ExpressionClassify[constdef.Expr_Bool]:
		v_return, v_err = ep.ParseExprBoolValue(p_expression)
	case constdef.ExpressionClassify[constdef.Expr_String]:
		v_return, v_err = ep.ParseExprStringValue(p_expression)
	case constdef.ExpressionClassify[constdef.Expr_Date]: // TODO 现在返回的是字符串
		v_return, v_err = ep.ParseExprDateValue(p_expression)
	case constdef.ExpressionClassify[constdef.Expr_Array]: // TODO 不只是识别纯数数组，任何数组均可
		v_return, v_err = ep.ParseExprArrayValue(p_expression)
	default:
		v_return = p_expression
	}
	return v_return, v_err
}

//解析变量表达式，并求表达式的值
//变量表达式分类：
//   变量表达式  => 解析变量表达式，并返回变量表达式的值
//Have read
func (ep *ExpressionParseEngine) EvaluateExpressionVariable(p_expression string) (interface{}, error) {
	var v_return interface{} = nil
	var v_err error = nil

	if p_expression == "" {
		errMsg := "EvaluateExpressionVariable Param[p_expression] is nil!"
		uniledgerlog.Warn(errMsg)
		return v_return, fmt.Errorf(errMsg)
	}
	if ep.IsExprVariable(p_expression) { // TODO : add else branch
		v_return, v_err = ep.ParseExprVariableValue(p_expression)
		if v_err != nil {
			var r_buf bytes.Buffer = bytes.Buffer{}
			r_buf.WriteString("[Result]:EvaluateExpressionVariable fail;")
			r_buf.WriteString("[Error]:" + v_err.Error())
			uniledgerlog.Warn(r_buf.String())
			return nil, v_err
		}
	}
	return v_return, v_err
}

//解析条件表达式，并求表达式的值
//条件表达式分类：
//   1）纯bool值  2）函数bool值  3）逻辑bool值
//Have read
func (ep *ExpressionParseEngine) EvaluateExpressionCondition(p_expression string) (bool, error) {
	var v_return bool = false
	var v_err error = nil

	if p_expression == "" {
		errMsg := "EvaluateExpressionCondition Param[p_expression] is nil!"
		uniledgerlog.Warn(errMsg)
		return v_return, fmt.Errorf(errMsg)
	}
	if ep.IsExprBool(p_expression) {
		v_return, v_err = ep.ParseExprBoolValue(p_expression)
	} else if ep.IsExprCondition(p_expression) {
		v_return, v_err = ep.ParseExprConditionValue(p_expression)
	} else if ep.IsExprFunction(p_expression) {
		var v_common_result common.OperateResult
		v_common_result, v_err = ep.ParseExprFunctionValue(p_expression)
		if v_common_result.GetCode() != 200 {
			uniledgerlog.Warn("[Result]:EvaluateExpressionCondition fail(Code != 200);")
			return v_return, v_err
		}
		str, ok := v_common_result.GetData().(string)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
			return v_return, fmt.Errorf("assert error")
		}
		v_return, v_err = strconv.ParseBool(str)
	} else {
		uniledgerlog.Error("type if %s is unknown", p_expression)
	}
	if v_err != nil {
		var r_buf bytes.Buffer = bytes.Buffer{}
		r_buf.WriteString("[Result]:EvaluateExpressionCondition fail;")
		r_buf.WriteString("[Error]:" + v_err.Error())
		uniledgerlog.Warn(r_buf.String())
	}
	return v_return, v_err
}

//解析函数表达式，并求表达式的值
//函数表达式分类：
//   函数表达式  => 解析函数表达式，并返回表达式的值
//Have read
func (ep *ExpressionParseEngine) EvaluateExpressionFunction(p_expression string) (common.OperateResult, error) {
	var v_return common.OperateResult = common.OperateResult{}
	var v_err error = nil
	if p_expression == "" {
		errMsg := "EvaluateExpressionFunction Param[p_expression] is nil!"
		uniledgerlog.Warn(errMsg)
		return v_return, fmt.Errorf(errMsg)
	}
	if ep.IsExprFunction(p_expression) { // TODO : add else branch
		v_return, v_err = ep.ParseExprFunctionValue(p_expression)
	}
	if v_err != nil {
		var r_buf bytes.Buffer = bytes.Buffer{}
		r_buf.WriteString("[Result]:EvaluateExpressionFunction fail;")
		r_buf.WriteString("[Error]:" + v_err.Error())
		uniledgerlog.Warn(r_buf.String())
	}
	return v_return, v_err
}

//解析决策候选者表达式，并求表达式的值
//决策候选者表达式分类：
//Have read
func (ep *ExpressionParseEngine) EvaluateExpressionCandidate(p_expression string) (interface{}, error) {
	var v_return interface{} = nil
	var v_err error = nil
	if p_expression == "" {
		errMsg := "EvaluateExpressionCandidate Param[p_expression] is nil!"
		uniledgerlog.Warn(errMsg)
		return v_return, fmt.Errorf(errMsg)
	}
	//TODO 待补充
	return v_return, v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//解析表达式所属分类
//   1.纯数字      => 直接返回该表达式  Expr_Num
//   2.纯浮点数    => 直接返回该表达式  Expr_Float
//   3.纯bool值    => 直接返回该表达式  Expr_Bool
//   4.纯字符串    => 直接返回该表达式  Expr_String
//   5.纯日期串    => 转化为时间戳返回  Expr_Date
//   6.纯数组串    => 转化为数组返回    Expr_Array
//   7.条件表达式  => 解析条件表达式，并返回表达式的值      Expr_Condition
//   8.函数表达式  => 解析函数表达式，并返回表达式的值      Expr_Function
//   9.变量表达式  => 解析变量表达式，并返回变量表达式的值  Expr_Variable
//Have read
func (ep *ExpressionParseEngine) ParseExpressionClassify(p_expression string) string {
	var v_classify string = ""

	if ep.IsExprNum(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_Num]
	} else if ep.IsExprFloat(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_Float]
	} else if ep.IsExprBool(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_Bool]
	} else if ep.IsExprString(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_String]
	} else if ep.IsExprDate(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_Date]
	} else if ep.IsExprArray(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_Array]
	} else if ep.IsExprCondition(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_Condition]
	} else if ep.IsExprFunction(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_Function]
	} else if ep.IsExprVariable(p_expression) {
		return constdef.ExpressionClassify[constdef.Expr_Variable]
	}
	return v_classify
}

//解析字符串是否为 单个 单词串
func (ep *ExpressionParseEngine) IsSingleWord(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Signal])
}

//解析字符串是否为 纯数字值
//Have read
func (ep *ExpressionParseEngine) IsExprNum(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Num])
}

//解析字符串是否为 纯浮点值
//Have read
func (ep *ExpressionParseEngine) IsExprFloat(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Float])
}

//解析字符串是否为 纯Bool值
//Have read
func (ep *ExpressionParseEngine) IsExprBool(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Bool])
}

//解析字符串是否为 纯字符串值
//Have read
func (ep *ExpressionParseEngine) IsExprString(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_String])
}

//解析字符串是否为 纯日期类型
//Have read
func (ep *ExpressionParseEngine) IsExprDate(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Date])
}

//解析字符串是否为 纯数组值
//Have read
func (ep *ExpressionParseEngine) IsExprArray(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Array])
}

//解析字符串是否为 条件表达式
//Have read
func (ep *ExpressionParseEngine) IsExprCondition(p_expression string) bool {
	if strings.HasPrefix(p_expression, "Func") {
		return false
	} else {
		return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Condition])
	}
}

//解析字符串是否为 函数串表达式
//Have read
func (ep *ExpressionParseEngine) IsExprFunction(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Func])
}

//解析字符串是否为 常量表达式
//Have read
func (ep *ExpressionParseEngine) IsExprConst(p_expression string) bool {
	return ep.IsExprFloat(p_expression) || ep.IsExprBool(p_expression) || ep.IsExprDate(p_expression) || ep.IsExprString(p_expression) || ep.IsExprNum(p_expression) || ep.IsExprArray(p_expression)
}

//解析字符串是否为 变量表达式
//Have read
func (ep *ExpressionParseEngine) IsExprVariable(p_expression string) bool {
	return ep.IsMatchRegexp(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name])
}

//---------------------------------------------------------------------------
//解析变量名是否为：合约名称[1 合约层]
//Have read
func (ep *ExpressionParseEngine) IsNameContract(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Contract])
}

//解析变量名是否为： 查询组件名称[2 Task组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameTaskEnquiry(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Task_Enquiry])
}

//解析变量名是否为 动作组件名称[2 Task组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameTaskAction(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Task_Action])
}

//解析变量名是否为 决策组件名称[2  Task组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameTaskDecision(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Task_Decision])
}

//解析变量名是否为 计划组件名称[2 Task组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameTaskPlan(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Task_Plan])
}

//解析变量名是否为 候选组件名称[2 Task组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameTaskCandidate(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Task_Candidate])
}

//解析变量名是否为 IntData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataInt(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_Int])
}

//解析变量名是否为 UintData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataUint(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_Uint])
}

//解析变量名是否为 FloatData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataFloat(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_Float])
}

//解析变量名是否为 TextData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataText(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_Text])
}

//解析变量名是否为 DateData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataDate(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_Date])
}

//解析变量名是否为 ArrayData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataArray(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_Array])
}

//解析变量名是否为 MatrixData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataMatrix(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_Matrix])
}

//解析变量名是否为 CompoundData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataCompound(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_Compound])
}

//解析变量名是否为 OperateResultData组件名称[3 Data组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameDataOperateResult(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Data_OperateResult])
}

//解析变量名是否为 FunctionExpression组件名称[4 Expression组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameExprFunc(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Expr_Func])
}

//解析变量名是否为 LogicArgumentExpression组件名称[4 Expression组件层]
//Have read
func (ep *ExpressionParseEngine) IsNameExprArgu(p_expression string) bool {
	return strings.HasPrefix(p_expression, constdef.ExpressionRegexp[constdef.Regexp_Name_Expr_Argu])
}

//---------------------------------------------------------------------------
//解析字符串是否完全匹配正则表达式
//Args: p_expression  => 待匹配字符串
//      p_regstr      => 正则表达式
func (ep *ExpressionParseEngine) IsMatchRegexp(p_expression string, p_regstr string) bool {
	var v_bool bool = false
	if p_expression == "" {
		v_bool = true
	} else {
		p_expression = strings.TrimSpace(p_expression)
		var v_len int = len(p_expression)
		reg := regexp.MustCompile(p_regstr)
		v_str := reg.FindString(p_expression)
		if len(v_str) == v_len {
			v_bool = true
		} else {
			v_bool = false
		}
	}
	return v_bool
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//解析表达式的值 纯数字值
//Have read
func (ep *ExpressionParseEngine) ParseExprNumValue(p_expression string) (int64, error) {
	var v_return int64
	var v_err error = nil

	v_return, v_err = strconv.ParseInt(strings.TrimSpace(p_expression), 10, 64)
	return v_return, v_err
}

//解析表达式的值 纯浮点数值
//Have read
func (ep *ExpressionParseEngine) ParseExprFloatValue(p_expression string) (float64, error) {
	var v_return float64
	var v_err error = nil
	v_return, v_err = strconv.ParseFloat(strings.TrimSpace(p_expression), 64)
	return v_return, v_err
}

//解析表达式的值 纯Bool值
//Have read
func (ep *ExpressionParseEngine) ParseExprBoolValue(p_expression string) (bool, error) {
	var v_return bool
	var v_err error = nil
	v_return, v_err = strconv.ParseBool(strings.TrimSpace(p_expression))
	return v_return, v_err
}

//解析表达式的值 纯字符串值
//Have read
func (ep *ExpressionParseEngine) ParseExprStringValue(p_expression string) (string, error) {
	var v_return string
	var v_err error = nil
	v_return = strings.TrimSpace(p_expression)
	return v_return, v_err
}

//解析表达式的值 纯日期类型
//Have read
func (ep *ExpressionParseEngine) ParseExprDateValue(p_expression string) (string, error) {
	var v_return string
	var v_err error = nil
	v_return, v_err = common.GenSpecialTimestamp(strings.TrimSpace(p_expression))
	return v_return, v_err
}

//解析表达式的值 数组类型
//Have read
func (ep *ExpressionParseEngine) ParseExprArrayValue(p_expression string) ([]interface{}, error) {
	var v_return []interface{}
	var v_err error = nil
	//过滤字符串两边的空格
	p_expression = strings.TrimSpace(p_expression)
	//过滤数组表达式两边的方括号
	p_expression = strings.TrimLeft(p_expression, "[")
	p_expression = strings.TrimRight(p_expression, "]")
	//按，分隔符，获取数组元素
	v_expression_array := strings.Split(p_expression, ",")
	for _, v_expr := range v_expression_array {
		v_return = append(v_return, strings.TrimSpace(v_expr))
	}
	return v_return, v_err
}

//解析表达式的值 纯条件表达式值
//重点：
//Have read
func (ep *ExpressionParseEngine) ParseExprConditionValue(p_expression string) (bool, error) {
	var v_err error = nil
	//识别条件表达式中的变量 map[string]string
	v_variables, v_err := ep.ParseVariablesInExprCondition(p_expression)
	if v_err != nil {
		uniledgerlog.Warn("ParseVariablesInExprCondition fail(" + v_err.Error() + ")")
		return false, v_err
	}

	//识别并设置表达中的变量的值
	v_parameters := make(map[string]interface{}, len(v_variables))
	for _, v_param := range v_variables {
		uniledgerlog.Debug("======variable: " + v_param)
		if ep.IsExprFunction(v_param) { //函数变量表达式
			v_result, v_err := ep.ParseExprFunctionValue(v_param)
			if v_err != nil {
				uniledgerlog.Warn("ParseExprFunctionValue fail(" + v_err.Error() + ")")
				return false, v_err
			}
			_, ok := v_result.GetData().(string)
			if !ok {
				v_err = fmt.Errorf("ParseExprFunctionValue[" + v_param + "] fail!")
				uniledgerlog.Warn("ParseExprFunctionValue fail(" + v_err.Error() + ")!")
				return false, v_err
			}
			v_parameters[v_param] = v_result.GetData().(string)
			p_expression = strings.Replace(p_expression, v_param, v_result.GetData().(string), -1)
		} else if ep.IsExprVariable(v_param) { // 变量表达式
			v_result, v_err := ep.ParseExprVariableValue(v_param)
			if v_err != nil || v_result == nil {
				uniledgerlog.Warn("ParseExprVariableValue fail(" + v_err.Error() + ")")
				return false, v_err
			}
			v_parameters[v_param] = v_result
			var str_result string = ""

			switch v_result.(type) {
			// TODO : 类型还要根据实际增加
			case int:
				{
					uniledgerlog.Debug("v_result.(type) is int")
					str_result = strconv.Itoa(v_result.(int))
				}
			case string:
				{
					uniledgerlog.Debug("v_result.(type) is string")

					// 以下针对data_date 日期格式数据做处理
					tmp_str := v_result.(string)
					if str_time, err := TryConvertToDate(tmp_str); err == nil {
						tmp_str = str_time
					}

					str_result = "\"" + tmp_str + "\""
				}
			case reflect.Value:
				{
					uniledgerlog.Debug("v_result.(type) is reflect.Value")
					switch v_result.(reflect.Value).Type().Kind() {
					// TODO : 类型还要根据实际增加
					case reflect.Bool:
						{
							uniledgerlog.Debug("v_result.(reflect.Value).Type().Kind() is bool")
							str_result = common.TernaryOperator(v_result.(reflect.Value).Bool(),
								"true", "false").(string)
						}
					case reflect.String:
						{
							uniledgerlog.Debug("v_result.(reflect.Value).Type().Kind() is string")
							str_result = "\"" + v_result.(reflect.Value).String() + "\""
						}
					default:
						uniledgerlog.Warn("default : v_result.(reflect.Value).Type().Kind() is %+v",
							v_result.(reflect.Value).Type().Kind())
					}
				}
			default:
				uniledgerlog.Warn("default : v_result.(type) is %+v", reflect.TypeOf(v_result))
			}
			uniledgerlog.Debug("variable: ", v_param, "  value: ", str_result)
			p_expression = strings.Replace(p_expression, v_param, str_result, -1)
		}
	}
	uniledgerlog.Debug("expression is: ", p_expression)
	//Eval 条件表达式的值
	v_expression, v_err := govaluate.NewEvaluableExpression(p_expression)
	if v_err != nil {
		uniledgerlog.Warn("govaluate.NewEvaluableExpression fail(" + v_err.Error() + ")")
		return false, v_err
	}
	//Eval 表达式的值
	//v_result, v_err := v_expression.Evaluate(v_parameters)
	v_result, v_err := v_expression.Evaluate(nil)
	if v_err != nil {
		uniledgerlog.Warn("expression.Evaluate fail(" + v_err.Error() + ")")
		return false, v_err
	}
	b, ok := v_result.(bool)
	if !ok {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		return false, fmt.Errorf("assert error")
	}
	uniledgerlog.Debug("expression evaluate result is : ", b)
	return b, v_err
}

//解析表达式的值 纯函数值
//重点：
//Have read
func (ep *ExpressionParseEngine) ParseExprFunctionValue(p_expression string) (common.OperateResult, error) {
	var v_return common.OperateResult = common.OperateResult{}
	var v_err error = nil
	v_return, v_err = ep.RunFunction(p_expression)
	if v_err != nil {
		uniledgerlog.Error("RunFunction(" + p_expression + ") fail(" + v_err.Error() + ")")
		//return v_return, v_err
	}
	return v_return, v_err
}

//解析表达式的值 纯变量值（默认变量都是component_table, property_table中存储的变量）
//     约定component， property规则
//变量举例：concract_xxxx._ContractAssetsList.assert_xxxx._MetaData.meta_xxxxxx
//          concract_xxxx._ContractState
//          enquiry_xxxxx._DataList.
//重点：
//Have read
func (ep *ExpressionParseEngine) ParseExprVariableValue(p_expression string) (interface{}, error) {
	var v_err error = nil
	//过滤字符串两边的空格
	p_expression = strings.TrimSpace(p_expression)
	//按分隔符“.”将字符串分割成变量数组
	v_variable_array := strings.Split(p_expression, ".")
	//获取变量的层数
	//  第一层  ：component_name
	//  第二层  ：property_name
	//  第三层...第N层： 子属性变量值
	//  最后一层：Value
	v_variable_count := len(v_variable_array)
	uniledgerlog.Info("###### v_variable_count is ", v_variable_count)
	uniledgerlog.Info("###### v_variable_array is ", v_variable_array)
	//注意：变量最少需要涵盖两层：component层（v_variable_array[0]）， property层（v_variable_array[1]）
	//识别第一层：component_name from component_table
	var parse_component inf.IComponent
	v_component := ep.Contract.GetComponentTtem(v_variable_array[0])
	if v_component == nil {
		uniledgerlog.Info("###### v_component is nil")
	}
	parse_component = ep.ReflectComponent(v_component, v_variable_array[0])
	//识别第二层： property from property_table
	var v_component_object reflect.Value = reflect.Value{}
	var v_property_field reflect.Value = reflect.Value{}

	v_component_object = reflect.ValueOf(parse_component).Elem()
	fmt.Println("=================================================================")
	uniledgerlog.Debug("component - v_variable_array[0] : ")
	uniledgerlog.Debug(v_variable_array[0])
	uniledgerlog.Debug("component - v_component_object.Kind() : ")
	uniledgerlog.Debug(v_component_object.Kind())
	uniledgerlog.Debug("component - v_component_object.Type() : ")
	uniledgerlog.Debug(v_component_object.Type())
	uniledgerlog.Debug("component - v_component_object.Interface() : ")
	uniledgerlog.Debug("%+v", v_component_object.Interface())
	fmt.Println("=================================================================")
	//两层达到时，直接返回值
	if v_variable_count == 2 {
		v_property_field = v_component_object.FieldByName(v_variable_array[1])
		if !v_property_field.IsValid() {
			v_err = fmt.Errorf("field[" + v_variable_array[1] + "] is not valid!")
			return nil, v_err
		}
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		uniledgerlog.Debug("field - v_variable_array[1] :")
		uniledgerlog.Debug(v_variable_array[1])
		uniledgerlog.Debug("field - v_property_field.Kind() :")
		uniledgerlog.Debug(v_property_field.Kind())
		uniledgerlog.Debug("field - v_property_field.String():")
		uniledgerlog.Debug(v_property_field.String())
		uniledgerlog.Debug("field - v_property_field.IsValid() :")
		uniledgerlog.Debug(v_property_field.IsValid())
		uniledgerlog.Debug("field - v_property_field.Interface() : ")
		uniledgerlog.Debug("%+v", v_property_field.Interface())
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		return v_property_field.Interface(), v_err
	}

	//识别第三层以后的：. subItem from array, map, and other
	v_idx := 1
	for v_idx < v_variable_count {
		v_property_field = v_component_object.FieldByName(v_variable_array[v_idx])
		uniledgerlog.Debug("======field: ", v_variable_array[v_idx], v_property_field.Kind(), " ",
			v_property_field.String(), " ", v_property_field.IsValid())

		switch v_property_field.Kind() {
		case reflect.Map:
			uniledgerlog.Debug("Map")
			v_idx = v_idx + 1
			if v_idx >= v_variable_count {
				break
			}

			v_component_object = v_property_field.MapIndex(reflect.ValueOf(v_variable_array[v_idx]))
			v_property_field = reflect.ValueOf(v_component_object)
			uniledgerlog.Debug("======field: ", v_variable_array[v_idx], v_property_field.Kind(), " ",
				v_property_field.String(), " ", v_property_field.IsValid())
		case reflect.Slice:
			uniledgerlog.Debug("Slice")
			v_idx = v_idx + 1
			if v_idx >= v_variable_count {
				break
			}
			v_arr_idx, _ := strconv.Atoi(v_variable_array[v_idx])
			//TODO: 随着字段类型的增加，需要补充支持
			switch v_property_field.Interface().(type) {
			case []string:
				data_arr := v_property_field.Interface().([]string)
				v_component_object = reflect.ValueOf(data_arr[v_arr_idx])
			case []interface{}:
				data_arr := v_property_field.Interface().([]interface{})
				v_component_object = reflect.ValueOf(data_arr[v_arr_idx])
			case []task.DecisionCandidate:
				data_arr := v_property_field.Interface().([]task.DecisionCandidate)
				v_component_object = reflect.ValueOf(data_arr[v_arr_idx])
			default:
				panic("v_property_field.Interface().(type) unknown")
			}

			v_property_field = reflect.ValueOf(v_component_object)
			uniledgerlog.Debug("======field: ", v_variable_array[v_idx], v_property_field.Kind(), " ",
				v_property_field.String(), " ", v_property_field.IsValid())
		case reflect.Array:
			uniledgerlog.Debug("Array")
			v_idx = v_idx + 1
			if v_idx >= v_variable_count {
				break
			}
			v_arr_idx, _ := strconv.Atoi(v_variable_array[v_idx])
			//TODO：随着字段类型的增加，需要补充支持
			switch v_property_field.Interface().(type) {
			case [2]int:
				data_arr := v_property_field.Interface().([2]int)
				v_component_object = reflect.ValueOf(data_arr[v_arr_idx])
			default:
				v_component_object = reflect.ValueOf(v_property_field.Interface())
			}

			v_property_field = reflect.ValueOf(v_component_object)
			uniledgerlog.Debug("======field: ", v_variable_array[v_idx], v_property_field.Kind(), " ",
				v_property_field.String(), " ", v_property_field.IsValid())
		case reflect.Struct:
			v_struct_property := v_property_field.Interface()
			v_component_object = reflect.ValueOf(v_struct_property)
			uniledgerlog.Debug("Struct")
		default:
			uniledgerlog.Debug("default")
		}

		v_idx = v_idx + 1
	}
	uniledgerlog.Debug("======field: ", v_property_field.Interface())
	return v_property_field.Interface(), v_err
}

//TODO： 可视化设计中组件的定义规则
//   合约组件：
//   任务组件：
//   数据组件：
//   描述组件：
//Have read
func (ep *ExpressionParseEngine) ReflectComponent(p_component interface{}, p_variable string) inf.IComponent {
	var parse_component inf.IComponent
	ok := false
	if ep.IsNameContract(p_variable) {
		parse_component, ok = p_component.(inf.ICognitiveContract)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameTaskEnquiry(p_variable) {
		parse_component, ok = p_component.(*task.Enquiry)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameTaskAction(p_variable) {
		parse_component, ok = p_component.(*task.Action)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameTaskDecision(p_variable) {
		parse_component, ok = p_component.(*task.Decision)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameTaskPlan(p_variable) {
		parse_component, ok = p_component.(*task.Plan)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataInt(p_variable) {
		parse_component, ok = p_component.(*data.IntData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataUint(p_variable) {
		parse_component, ok = p_component.(*data.UintData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataFloat(p_variable) {
		parse_component, ok = p_component.(*data.FloatData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataText(p_variable) {
		parse_component, ok = p_component.(*data.TextData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataDate(p_variable) {
		parse_component, ok = p_component.(*data.DateData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataArray(p_variable) {
		parse_component, ok = p_component.(*data.ArrayData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataMatrix(p_variable) {
		parse_component, ok = p_component.(*data.MatrixData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataCompound(p_variable) {
		parse_component, ok = p_component.(*data.CompoundData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameDataOperateResult(p_variable) {
		parse_component, ok = p_component.(*data.OperateResultData)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameExprFunc(p_variable) {
		parse_component, ok = p_component.(*expression.Function)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	} else if ep.IsNameExprArgu(p_variable) {
		parse_component, ok = p_component.(*expression.LogicArgument)
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.ASSERT_ERROR, ""))
		}
	}
	return parse_component
}

//---------------------------------------------------------------------------
//截取字符串
//Args:  str   待截断字符串
//       start 字符串起点
//       end   字符串终点
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		panic("start is wrong")
	}
	if end < 0 || end > length {
		panic("end is wrong")
	}
	return string(rs[start:end])
}

//按指定字符切割字符串
//Args: 待切割字符串 p_str string
//      切割符号集   p_reg string
func SplitString(p_str string, p_reg string) map[string]string {
	var map_result map[string]string = make(map[string]string, 0)
	if p_str == "" {
		return map_result
	}
	reg := regexp.MustCompile(constdef.ExpressionTagString)
	dataSlice := reg.FindAllIndex([]byte(p_str), -1)
	var int_begin int = 0
	var int_end int = 0
	var int_temp int = 0
	for idx, v_array := range dataSlice {
		if idx == 0 {
			int_begin = 0
		} else {
			int_begin = int_temp
		}
		int_end = v_array[0]
		int_temp = v_array[1]
		map_result[Substr2(p_str, int_begin, int_end)] = Substr2(p_str, int_begin, int_end)
	}
	map_result[Substr2(p_str, int_temp, len(p_str))] = Substr2(p_str, int_temp, len(p_str))
	uniledgerlog.Debug(map_result)
	return map_result
}

//识别条件表达式中变量数组
//Have read
func (ep *ExpressionParseEngine) ParseVariablesInExprCondition(p_expression string) (map[string]string, error) {
	var v_return map[string]string = nil
	var v_err error = nil
	if p_expression == "" {
		uniledgerlog.Warn("ParseVariablesInExprCondition param is nil!")
		return v_return, v_err
	}
	//初始化返回结果
	v_return = make(map[string]string, 0)
	//获取分隔符字符串
	v_variable_arr := SplitString(p_expression, constdef.ExpressionTagString) // TODO : ?
	for _, v_variable := range v_variable_arr {
		v_variable = strings.TrimSpace(v_variable)
		if v_variable == "" {
			continue
		} else if ep.IsExprNum(v_variable) {
			continue
		} else if ep.IsExprFloat(v_variable) {
			continue
		} else if ep.IsExprBool(v_variable) {
			continue
		} else if ep.IsExprString(v_variable) {
			continue
		} else if ep.IsExprArray(v_variable) {
			continue
		} else if ep.IsExprDate(v_variable) {
			continue
		} else if ep.IsExprFunction(v_variable) {
			v_return[v_variable] = v_variable
			continue
		} else {
			//带半边括号的变量，去掉两边的括号
			v_variable = strings.TrimLeft(v_variable, "(")
			v_variable = strings.TrimRight(v_variable, ")")
			v_variable = strings.TrimSpace(v_variable)
			v_return[v_variable] = v_variable
		}
	}
	return v_return, v_err
}

//---------------------------------------------------------------------------
//---------------------------------------------------------------------------
//函数运行
//Have read
func (ep *ExpressionParseEngine) RunFunction(p_function string) (common.OperateResult, error) {
	var v_err error = nil
	var v_result common.OperateResult = common.OperateResult{}
	var r_buf bytes.Buffer = bytes.Buffer{}
	if p_function == "" {
		r_buf.WriteString("[Result]: RunFunction(" + p_function + ") fail;")
		r_buf.WriteString("[Error]: param[p_function] is nil!")
		uniledgerlog.Warn(r_buf.String())
		v_err = fmt.Errorf(" param[p_function] is nil!")
		v_result = common.OperateResult{Code: 400, Message: r_buf.String()}
		return v_result, v_err
	}
	//校验参数格式 xxxx(xxx, xxx, xxx)
	if v_bool, v_err := regexp.MatchString(`[a-zA-Z_0-9]+\(.*\)`, p_function); !v_bool || v_err != nil {
		r_buf.WriteString("[Result]: RunFunction(" + p_function + ") fail;")
		r_buf.WriteString("[Error]: param[p_function] format error!")
		uniledgerlog.Warn(r_buf.String())
		v_err = fmt.Errorf(" param[p_function] format error!")
		v_result = common.OperateResult{Code: 400, Message: r_buf.String()}
		return v_result, v_err
	}

	//正则匹配函数名
	name_reg := regexp.MustCompile(`\s*([^\(]+)`)
	func_name := strings.TrimSpace(name_reg.FindString(p_function))
	uniledgerlog.Debug("=======func_name:", func_name)

	//正则匹配函数的参数变量
	param_reg := regexp.MustCompile(`\((.*)\)`)
	func_param_str := strings.Trim(param_reg.FindString(p_function), "(|)")
	uniledgerlog.Debug("=======func_params:", func_param_str)

	//函数调用
	var func_run reflect.Value
	var func_params []reflect.Value
	var func_params_ map[string]interface{}
	count := 1
	if gRPCClient.On {
		func_params_ = make(map[string]interface{}, 10)
		_ = func_run
		_ = func_params
	} else {
		func_run = reflect.ValueOf(ep.FunctionEngine.ContractFunctions[func_name])
		_ = count
		_ = func_params_
	}

	if func_param_str != "" {
		//分割匹配的函数参数列表
		//TODO 大参数解析（比如json串解析）
		var func_param_array []string
		var flag bool
		var funcType int32
		if !gRPCClient.On {
			if func_name == "FuncTransferAsset" {
				flag = true
			}
		} else {
			funcType, v_err = gRPCClient.QueryFuncType(func_name)
			if v_err != nil {
				return v_result, v_err
			}
			if funcType == 2 {
				flag = true
			}
		}
		if flag {
			if !gRPCClient.On {
				func_param_array = make([]string, 8)
				first_param_array := strings.Split(func_param_str, "@")
				first_idx := 0
				for t_idx, t_param := range strings.Split(first_param_array[0], ",") {
					func_param_array[first_idx+t_idx] = t_param
				}
				func_param_array[3] = first_param_array[1]
				second_idx := 4
				for v_idx, v_param := range strings.Split(first_param_array[2], ",") {
					func_param_array[second_idx+v_idx] = v_param
				}
			} else {
				var func_param_tmp []string
				first_param_array := strings.Split(func_param_str, "@")
				for _, t_param := range strings.Split(first_param_array[0], ",") {
					func_param_tmp = append(func_param_tmp, t_param)
				}
				first_param_array[1] = strings.Replace(first_param_array[1], "\\\"", "\"", -1)
				first_param_array[1] = strings.Replace(first_param_array[1], "\\\\", "\\", -1)
				func_param_tmp = append(func_param_tmp, first_param_array[1])
				for _, v_param := range strings.Split(first_param_array[2], ",") {
					func_param_tmp = append(func_param_tmp, v_param)
				}

				num := len(func_param_tmp) - 6
				func_param_array = func_param_tmp[num:]
				for index := 0; index < num; index++ {
					func_param_array = append(func_param_array, func_param_tmp[index])
				}
			}
		} else {
			func_param_array = strings.Split(func_param_str, ",")
		}

		if !gRPCClient.On {
			func_params = make([]reflect.Value, len(func_param_array))
		}
		for index, v_args := range func_param_array {
			isConstant := func(v string) bool {
				return ep.IsExprNum(v) || ep.IsExprBool(v) || ep.IsExprFloat(v) || ep.IsExprString(v)
			}(v_args)

			//识别字符串，获取参数的值
			if isConstant {
				v_args = strings.TrimLeft(v_args, "\\\"")
				v_args = strings.TrimRight(v_args, "\\\"")
				//v_args = v_args.re
				if gRPCClient.On {
					func_params_[fmt.Sprintf("Param%02d", count)] = v_args
					count++
					_ = index
				} else {
					func_params[index] = reflect.ValueOf(v_args)
				}
			} else {
				v_arg_value, err := ep.EvaluateExpressionVariable(v_args)
				if err != nil {
					r_buf.WriteString("[Result]: RunFunction(" + p_function + ") fail;")
					r_buf.WriteString("[Error]: function args(" + v_args + ") parse error!")
					uniledgerlog.Warn(r_buf.String())
					v_err = fmt.Errorf(" function args(" + v_args + ") parse error!")
					v_result = common.OperateResult{Code: 400, Message: r_buf.String()}
					return v_result, v_err
				}

				if gRPCClient.On {
					var inter interface{}
					switch v_arg_value.(type) {
					case reflect.Value:
						{
							uniledgerlog.Debug("v_arg_value.(type) is reflect.Value")
							switch v_arg_value.(reflect.Value).Type().Kind() {
							// TODO : 类型还要根据实际增加
							case reflect.Int:
								{
									uniledgerlog.Debug("v_arg_value.(reflect.Value).Type().Kind() is int")
									inter = v_arg_value.(reflect.Value).Int()
								}
							case reflect.Bool:
								{
									uniledgerlog.Debug("v_arg_value.(reflect.Value).Type().Kind() is bool")
									inter = v_arg_value.(reflect.Value).Bool()
								}
							case reflect.String:
								{
									uniledgerlog.Debug("v_arg_value.(reflect.Value).Type().Kind() is string")
									inter = v_arg_value.(reflect.Value).String()
								}
							default:
								uniledgerlog.Warn("default : v_arg_value.(reflect.Value).Type().Kind() is %+v",
									v_arg_value.(reflect.Value).Type().Kind())
								inter = v_arg_value
							}
						}
					default:
						{
							uniledgerlog.Debug("v_arg_value.(type) is %+v", reflect.TypeOf(v_arg_value))
							inter = v_arg_value
						}
					}

					func_params_[fmt.Sprintf("Param%02d", count)] = inter
					count++
					_ = index
				} else {
					func_params[index] = reflect.ValueOf(v_arg_value)
				}
			}
		}
	}

	if gRPCClient.On {
		slData, v_err := json.Marshal(func_params_)
		if v_err != nil {
			r_buf.WriteString("[Result]: RunFunction(" + p_function + ") fail;")
			r_buf.WriteString("[Error]: " + v_err.Error() + "!")
			uniledgerlog.Error(r_buf.String())
			v_err = fmt.Errorf(r_buf.String())
			v_result = common.OperateResult{Code: 400, Message: r_buf.String()}
			return v_result, v_err
		}

		hostname, _ := os.Hostname()
		v_result, v_err = gRPCClient.FunctionRun(hostname+"|"+common0.GenTimestamp(), func_name, string(slData))
		return v_result, v_err
	} else {
		func_result_arr := func_run.Call(func_params)
		r_buf.WriteString("[Result]: RunFunction(" + p_function + ") success;")
		uniledgerlog.Info(r_buf.String())

		retResult, ok1 := func_result_arr[0].Interface().(common.OperateResult)
		if !ok1 {
			retResult = common.OperateResult{}
		}
		retErr, ok2 := func_result_arr[1].Interface().(error)
		if !ok2 {
			retErr = nil
		}
		uniledgerlog.Info(retResult, retErr)
		return retResult, retErr
	}
}

// 尝试转换成时间函数
func TryConvertToDate(str_timeStamp string) (string, error) {
	timeStamp, err := strconv.Atoi(str_timeStamp)
	if err != nil {
		return "", err
	}
	tm := time.Unix(int64(timeStamp)/1000, 0)
	format := "2006-01-02 15:04:05"
	timeFormat, err := time.Parse(format, tm.Format(format))
	if err != nil {
		return "", err
	}
	return timeFormat.Format(format), err
}
