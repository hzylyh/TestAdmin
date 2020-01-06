package hassert

import (
	"fmt"
	"github.com/tidwall/gjson"
	"reflect"
)

type HAssert struct {
}

func (ha *HAssert) AssertEqual(exp, act string) {

}

//func (ha *HAssert)  {
//
//}

func MulAssert(exp, act string, verifyCols []string) {
	defer recoverAssert()
	for _, verifyCol := range verifyCols {
		exp := gjson.Get(exp, verifyCol)
		act := gjson.Get(act, verifyCol)

		// 根据jsonpath取出的字段，进行校验，后续考虑panic跟error结合使用
		assertCore(exp.Value(), act.Value())
	}

}

func recoverAssert() {
	// 后续写日志表日库逻辑，可扩展成接口，对接不同记录模式
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	} else {
		fmt.Println("我是正常的")
	}
}

func assertCore(exp, act interface{}) {
	if exp == nil {
		assertEqual(exp, act)
	} else {
		expIn := reflect.TypeOf(exp)
		//actIn := reflect.TypeOf(act)
		//fmt.Println(expIn.Name(), expIn.Kind())

		if expIn.Kind() == reflect.Slice {
			expList, ok1 := exp.([]interface{})
			actList, ok2 := act.([]interface{})
			if ok1 && ok2 {
				for i := 0; i < len(expList); i++ {
					assertCore(expList[i], actList[i])
				}
			} else {
				//fmt.Println("slice类型转换异常")
				panic(fmt.Sprintf("%v !== %v", expList, actList))
			}
		} else if expIn.Kind() == reflect.Map {
			expMap, expOk := exp.(map[string]interface{})
			actMap, actOk := act.(map[string]interface{})
			if expOk && actOk {
				for k := range expMap {
					assertCore(expMap[k], actMap[k])
				}
			} else {
				//fmt.Println("map类型转换异常")
				panic(fmt.Sprintf("%v !== %v", expMap, actMap))
			}

		} else if expIn.Kind() == reflect.String {
			//fmt.Println("String")
			expString, expOk := exp.(string)
			actString, actOk := act.(string)
			if expOk && actOk {
				assertEqual(expString, actString)
			}
		} else {
			assertEqual(exp, act)
		}
	}
}

func assertEqual(exp, act interface{}) {
	if exp == act {
		fmt.Printf("%v == %v\n", exp, act)
	} else {
		panic(fmt.Sprintf("%v != %v\n", exp, act))
	}
}
