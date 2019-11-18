package model

import "salotto/model/InterfaceTestPartEntity"

var Models = []interface{}{
	&Schedule{}, &ScheduleJob{}, &ProjectInfo{},
	&InterfaceTestPartEntity.InterfaceInfo{},
	&InterfaceTestPartEntity.ItfCaseInfo{},
	&InterfaceTestPartEntity.ItfCaseStepInfo{},
	&InterfaceTestPartEntity.ProjectModule{},
}
