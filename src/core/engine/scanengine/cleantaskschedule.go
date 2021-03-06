/*************************************************************************
  > File Name: cleantaskschedule.go
  > Module:
  > Function: 清理数据表（TaskSchedule）中的过期或者已经执行成功的任务
  > Author: wangyp
  > Company:
  > Department:
  > Mail: wangyepeng87@163.com
  > Created Time: 2017.05.08
 ************************************************************************/
package scanengine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

import (
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	"unicontract/src/core/engine/common"
	"unicontract/src/core/engine/common/db"
)

//---------------------------------------------------------------------------
func _CleanTaskSchedule() {
	ticker := time.NewTicker(time.Minute * (time.Duration)(scanEngineConf["clean_time"].(int)))
	for _ = range ticker.C {
		uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "query all success task"))
		strSuccessTask, err := common.GetTaskSchedulesSuccess(config.Config.Keypair.PublicKey)
		if err != nil {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
			continue
		}

		if len(strSuccessTask) == 0 {
			uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "success task is null"))
			continue
		}

		var slTasks []db.TaskSchedule
		json.Unmarshal([]byte(strSuccessTask), &slTasks)

		uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "success task filter"))
		slID := _TaskFilter(slTasks)

		if len(slID) == 0 {
			uniledgerlog.Debug(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "_TaskFilter return is null"))
			continue
		}

		uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "success task delete"))
		deleteNum, err := common.DeleteTaskSchedules(slID)
		if deleteNum != len(slID) {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
		}
	}
	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func _TaskFilter(slTasks []db.TaskSchedule) []interface{} {
	var slID []interface{}

	// 过滤时间点
	cleanTimePoint := time.Now().
		Add(-time.Hour*24*(time.Duration)(scanEngineConf["clean_data_time"].(int))).
		UnixNano() / 1000000
	for index, value := range slTasks {
		nTimePoint, err := strconv.Atoi(value.LastExecuteTime)
		if err != nil {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
			continue
		}
		if int64(nTimePoint) < cleanTimePoint {
			slID = append(slID, slTasks[index].Id)
		}
	}

	return slID
}

//---------------------------------------------------------------------------
