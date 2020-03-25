package stime

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(GetCurTime())
}

func GetCurTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
