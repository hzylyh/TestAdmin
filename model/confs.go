package model

import "salotto/model/InterfaceTestPartEntity"

var Models = []interface{}{
	&InterfaceTestPartEntity.TInterfaceInfo{},
	&InterfaceTestPartEntity.TItfCaseInfo{},
	&InterfaceTestPartEntity.TItfCaseStepInfo{},
	&InterfaceTestPartEntity.TProjectModule{},
	&InterfaceTestPartEntity.TCaseStepVarInfo{},
	&InterfaceTestPartEntity.TAssertInfo{},
	&InterfaceTestPartEntity.TItfCaseStepRunHis{},
	&InterfaceTestPartEntity.TInterfaceHeadersInfo{},
	&TSysUser{},
	&TProjectInfo{},
}
