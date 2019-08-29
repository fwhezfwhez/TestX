package main

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"shangraomajiang/control/task/taskModel"
)

var TaskMap = map[int]TaskI{
	0: CommonTask{},
}

type TaskI interface {
	AwardConfig(ut taskModel.UserTask) (TaskAward, error)
	UserJointConfig(utp taskModel.UserTaskProcedure) (TaskProcedure, error)
	FinishTask(utp *taskModel.UserTaskProcedure) error

	DefaultJointConfig() []byte
}
type CommonTask struct{}

type TaskAward struct {
	PropId   int `json:"prop_id"`   // 道具id
	PropNum  int `json:"prop_num"`  //道具数量
	ExpireIn int `json:"expire_in"` // 单位为小时，-1为无期限道具
	Point    int `json:"point"`     // 活跃度
}

type TaskProcedure struct {
	Times    int `json:"times"`     // 目标任务次数
	HasAward int `json:"has_award"` // 1-未领取，2-已领取
}

func (ct CommonTask) AwardConfig(ut taskModel.UserTask) (TaskAward, error) {
	var ta TaskAward
	if e := json.Unmarshal(ut.AwardConfigJson, &ta); e != nil {
		return TaskAward{}, errorx.Wrap(e)
	}

	return ta, nil
}
func (ct CommonTask) UserJointConfig(utp taskModel.UserTaskProcedure) (TaskProcedure, error) {
	var tp TaskProcedure

	if e := json.Unmarshal(utp.JointConfigJson, &tp); e != nil {
		return TaskProcedure{}, errorx.Wrap(e)
	}
	return tp, nil
}
func (ct CommonTask) FinishTask(utp *taskModel.UserTaskProcedure) error {
	taskProcedure, e := ct.UserJointConfig(*utp)
	if e != nil {
		return errorx.Wrap(e)
	}
	// 2- 已领取
	taskProcedure.HasAward = 2
	// 3- 已完成, 2-进行中
	utp.State = 3
	buf, e := json.Marshal(taskProcedure)
	if e != nil {
		return errorx.Wrap(e)
	}
	utp.JointConfigJson = buf
	utp.JointConfig = buf
	return nil
}
func (ct CommonTask) DefaultJointConfig() []byte {
	return []byte(`
{
    "times": 0 ,
    "has_award":1
}
`)
}
func main() {
	var tobj TaskI
	tobj = CommonTask{}
	fmt.Println(tobj.DefaultJointConfig())
}
