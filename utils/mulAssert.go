package utils

import (
	"fmt"
	"github.com/tidwall/gjson"
	"reflect"
)

func MulAssert(exp, act string, verifyCols []string) {

	for _, verifyCol := range verifyCols {
		exp := gjson.Get(exp, verifyCol)
		act := gjson.Get(act, verifyCol)

		// 根据jsonpath取出的字段，进行校验
		assertCore(exp.Value(), act.Value())
	}

}

func assertCore(exp, act interface{}) {
	expIn := reflect.TypeOf(exp)
	//actIn := reflect.TypeOf(act)
	//fmt.Println(expIn.Name(), expIn.Kind())

	if expIn.Kind() == reflect.Slice {
		expList := exp.([]interface{})
		actList := act.([]interface{})
		//fmt.Println("transformList", expList)
		//fmt.Println("transformList", actList)
		for i := 0; i < len(expList); i++ {
			assertCore(expList[i], actList[i])
		}
	} else if expIn.Kind() == reflect.Map {
		expMap := exp.(map[string]interface{})
		actMap := act.(map[string]interface{})
		//fmt.Println("transformMap", expMap)
		for k := range expMap {
			assertCore(expMap[k], actMap[k])
		}
	} else if expIn.Kind() == reflect.String {
		//fmt.Println("String")
		expString := exp.(string)
		actString := act.(string)
		assertEqual(expString, actString)
	} else {
		assertEqual(exp, act)
	}
}

func assertEqual(exp, act interface{}) {
	if exp == act {
		fmt.Printf("%v == %v\n", exp, act)
	} else {
		fmt.Printf("%v != %v\n", exp, act)
	}
}
