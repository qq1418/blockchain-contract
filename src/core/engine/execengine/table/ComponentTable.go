package table

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"unicontract/src/core/engine/execengine/constdef"
	"unicontract/src/core/engine/execengine/expression"
	"unicontract/src/core/engine/execengine/inf"
)

type ComponentTable struct {
	//TODO: need sort struct
	//type: map[string][]property.PropertyT
	//        type:  Unknown, Data, Task, Expression
	CompTable map[string][]map[string]interface{} `json:"CompTable"`
}

func NewComponentTable() *ComponentTable {
	ct := &ComponentTable{}
	return ct
}

func (ct *ComponentTable) getComponentType(p_component interface{}) (string, string) {
	var r_type string = ""
	var r_name string = ""
	if p_component == nil {
		r_type = constdef.ComponentType[constdef.Component_Unknown]
		ccom, ok := p_component.(inf.IComponent)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = ccom.GetName()
		return r_type, r_name
	}
	switch p_component.(type) {
	case *inf.IData:
		r_type = constdef.ComponentType[constdef.Component_Data]
		ddata, ok := p_component.(inf.IData)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = ddata.GetName()
	case inf.IData:
		r_type = constdef.ComponentType[constdef.Component_Data]
		ddata, ok := p_component.(inf.IData)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = ddata.GetName()
	case *inf.ITask:
		r_type = constdef.ComponentType[constdef.Component_Task]
		ttask, ok := p_component.(inf.ITask)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = ttask.GetName()
	case inf.ITask:
		r_type = constdef.ComponentType[constdef.Component_Task]
		ttask, ok := p_component.(inf.ITask)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = ttask.GetName()
	case *inf.IExpression:
		r_type = constdef.ComponentType[constdef.Component_Expression]
		eexp, ok := p_component.(inf.IExpression)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = eexp.GetName()
	case *expression.LogicArgument:
		r_type = constdef.ComponentType[constdef.Component_Expression]
		eexp, ok := p_component.(inf.IExpression)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = eexp.GetName()
	case *expression.Function:
		r_type = constdef.ComponentType[constdef.Component_Expression]
		eexp, ok := p_component.(inf.IExpression)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = eexp.GetName()
	case inf.IExpression:
		r_type = constdef.ComponentType[constdef.Component_Expression]
		eexp, ok := p_component.(inf.IExpression)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = eexp.GetName()
	case inf.ICognitiveContract:
		r_type = constdef.ComponentType[constdef.Component_Contract]
		eexp, ok := p_component.(inf.ICognitiveContract)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = eexp.GetName()
	case *inf.ICognitiveContract:
		r_type = constdef.ComponentType[constdef.Component_Contract]
		eexp, ok := p_component.(inf.ICognitiveContract)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = eexp.GetName()
	default:
		fmt.Println(reflect.ValueOf(p_component).Type())
		r_type = constdef.ComponentType[constdef.Component_Unknown]
		ccom, ok := p_component.(inf.IComponent)
		if !ok {
			logs.Error("assert error")
			return "", ""
		}
		r_name = ccom.GetName()
	}
	return r_type, r_name
}

func (ct *ComponentTable) AddComponent(p_component interface{}) {
	if ct.CompTable == nil {
		ct.CompTable = make(map[string][]map[string]interface{})
	}
	c_type, c_name := ct.getComponentType(p_component)
	if _, ok := ct.CompTable[c_type]; !ok {
		ct.CompTable[c_type] = make([]map[string]interface{}, 0)
	}
	v_component := make(map[string]interface{})
	v_component[c_name] = p_component
	ct.CompTable[c_type] = append(ct.CompTable[c_type], v_component)
}

func (ct *ComponentTable) GetComponent(cname string, ctype string) interface{} {
	var g_component interface{}
	if ctype == "" {
		//总：map[string][]map[string]interface()
		for _, ct_value := range ct.CompTable {
			for _, ct_com := range ct_value {
				if _, ok := ct_com[cname]; ok {
					g_component = ct_com[cname]
				}
			}
		}
	} else {
		//总：map[string][]map[string]interface()
		//for: map[string]
		if v_comp_type, ok := ct.CompTable[ctype]; ok {
			//for: []
			for _, v_comp := range v_comp_type {
				//map[string]interface()
				if r_res, ok := v_comp[cname]; ok {
					g_component = r_res
				}
			}
		}
	}
	return g_component
}

func (ct *ComponentTable) GetTaskByID(cid string, ctype string) interface{} {
	var g_component interface{}
	if ctype == constdef.ComponentType[constdef.Component_Task] {
		//总：map[string][]map[string]interface()
		//for: map[string]
		if v_comp_type, ok := ct.CompTable[ctype]; ok {
			//for: []
			for _, v_comp := range v_comp_type {
				//map[string]interface()
				for _, v_value := range v_comp {
					if v_value == nil {
						continue
					}
					ttask, ok := v_value.(inf.ITask)
					if !ok {
						logs.Error("assert error")
						return nil
					}
					if ttask.GetTaskId() == cid {
						g_component = v_value
						break
					}
				}
			}
		}
	}
	return g_component
}

func (ct *ComponentTable) GetComponentByType(c_type string) []map[string]interface{} {
	if c_type == "" {
		return nil
	}
	if _, ok := ct.CompTable[c_type]; !ok {
		return nil
	}
	return ct.CompTable[c_type]
}

//更新PropertyTable中的元素
//Args: c_type string 组件类型
//      c_name string 组件名称
//      c_component interface{} 组件
func (ct *ComponentTable) UpdateComponent(c_type string, c_name string, c_component interface{}) error {
	var err error = nil
	//获取ctype对应的组件数组
	task_component_array := ct.CompTable[c_type]
	var new_task_component_array []map[string]interface{} = make([]map[string]interface{}, len(task_component_array))
	for v_idx, v_component_map := range task_component_array {
		//替换组件数组中对应的组件
		for v_key, v_value := range v_component_map {
			if v_key == c_name {
				v_component_map[v_key] = c_component
			} else {
				v_component_map[v_key] = v_value
			}
		}
		new_task_component_array[v_idx] = v_component_map
	}
	//更新ctype对应的组件数组
	ct.CompTable[c_type] = new_task_component_array
	return err
}
