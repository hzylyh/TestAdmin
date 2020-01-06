package model

import "salotto/model/InterfaceTestPartEntity"

var Models = []interface{}{
	&Schedule{}, &ScheduleJob{}, &TProjectInfo{},
	&InterfaceTestPartEntity.TInterfaceInfo{},
	&InterfaceTestPartEntity.TItfCaseInfo{},
	&InterfaceTestPartEntity.TItfCaseStepInfo{},
	&InterfaceTestPartEntity.TProjectModule{},
	&InterfaceTestPartEntity.TCaseStepVarInfo{},
	&InterfaceTestPartEntity.TAssertInfo{},
}
