package parse

import (
	"fmt"
	"salotto/utils/stringext"
	"strings"
)

//func main() {
//	var properties = map[string]string{"name": "houzheyu", "age": "2222"}
//	var Ihandler TokenHandler
//	Ihandler = &VariableTokenHandler{
//		variables: properties,
//	}
//	a := GenericTokenParser{
//		openToken:  "${",
//		closeToken: "}",
//		handler:    Ihandler,
//	}
//	parse := a.Parse("my name is ${name}, my age is ${age}")
//	fmt.Println(parse)
//	//index := strings.Index("my name is ${name}, my age is ${age}", "${")
//	//fmt.Println(index)
//
//	//end := indexOf("my name is ${name}", "}", 17)
//	//fmt.Println(end)
//}

type GenericTokenParser struct {
	OpenToken  string
	CloseToken string
	Handler    TokenHandler
}

func (gtp *GenericTokenParser) Parse(text string) string {
	var (
		start   int
		end     int
		src     []string
		offset  int = 0
		builder StringBuilder
		//expression StringBuilder
		expression string
	)

	if text == "" {
		return ""
	}

	// 查找有没有openToken，如果没有，直接返回
	//start = strings.Index(text, gtp.OpenToken)
	// 解决中文索引位置错误问题
	start = stringext.UnicodeIndex(text, gtp.OpenToken)
	if start == -1 {
		return text
	}

	// 将整个字符串转换为单个字符的切片
	src = func(text string) []string {
		var res []string
		for _, char := range text {
			res = append(res, string(char))
		}
		return res
	}(text)

	for start > -1 {
		if start > 0 && src[start-1] == "\\" {
			fmt.Printf("准备舍弃")
		} else {
			builder.append(src, offset, start-offset) //
			offset = start + len(gtp.OpenToken)
			end = indexOf(text, gtp.CloseToken, offset)
			for end > -1 {
				if end > offset && src[end-1] == "\\" {
					fmt.Printf("准备舍弃")
				} else {
					//expression.append(src, offset, end - offset)
					expression = strings.Join(src[offset:offset+(end-offset)], "")
					offset = end + len(gtp.CloseToken)
					break
				}
			}
			if end == -1 {
				// 结束标志未发现
				builder.append(src, start, len(src)-start)
				offset = len(src)
			} else {
				//builder.appendString(gtp.handler.handleToken(expression.toString()))
				builder.appendString(gtp.Handler.handleToken(expression))
				offset = end + len(gtp.CloseToken)
			}
		}
		start = indexOf(text, gtp.OpenToken, offset)
	}

	if offset < len(src) {
		builder.append(src, offset, len(src)-offset)
	}

	return builder.toString()
}

type TokenHandler interface {
	handleToken(content string) string
}

type VariableTokenHandler struct {
	Variables map[string]string
}

func (vth *VariableTokenHandler) handleToken(content string) string {
	//fmt.Println("拼接后的字符串", content)
	return vth.Variables[content]
}

type StringBuilder struct {
	value []string
	count int
}

func (sb *StringBuilder) append(str []string, offset, len int) {
	appendStr := str[offset : offset+len]
	sb.value = append(sb.value, appendStr...)
}

func (sb *StringBuilder) appendString(str string) {
	var sliceStr []string
	for _, eachStr := range str {
		sliceStr = append(sliceStr, string(eachStr))
	}
	sb.value = append(sb.value, sliceStr...)
}

func (sb *StringBuilder) toString() string {
	return strings.Join(sb.value, "")
}

func indexOf(src, target string, offset int) int {
	//for i, char := range src {
	//	if i < offset {
	//		continue
	//	}
	//	if string(char) == target {
	//		return i
	//	}
	//}
	//return -1
	//subString := src[offset:]
	newText := []rune(src)
	subString := string(newText[offset:])
	//subStrIndex := strings.Index(subString, target)
	subStrIndex := stringext.UnicodeIndex(subString, target)
	if subStrIndex == -1 {
		return -1
	} else {
		return offset + subStrIndex
	}
}
