package main

import (
	"fmt"
	tmpMsg "github.com/wlibo666/json/example/msg"
)

func TestBaseMsg() {
	msg := &tmpMsg.BaseMsg{}

	fieldNameValues := make(map[string]string)
	fieldNameValues["Str_1"] = "wangchunyan"
	fieldNameValues["Int_1"] = "29"

	err := msg.SetValue(fieldNameValues)
	if err != nil {
		fmt.Printf("set value failed:%s", err.Error())
		return
	}

	data, err := msg.Marshal()
	if err != nil {
		fmt.Printf("marshal not map failed,err:%s", err.Error())
		return
	}
	fmt.Printf("not map data:%s\n", string(data))

	fieldNameAlias := make(map[string]string)
	fieldNameAlias["Str_1"] = "name"
	fieldNameAlias["Int_1"] = "age"

	data, err = msg.Marshal(fieldNameAlias)
	if err != nil {
		fmt.Printf("marshal map failed,err:%s", err.Error())
		return
	}
	fmt.Printf("map data:%s\n", string(data))
}

func main() {
	TestBaseMsg()
}
