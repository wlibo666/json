package main

import (
	"fmt"
	"github.com/wlibo666/json"
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

type B struct {
	Name string `json:"name"`
	A01  int    `json:"a01"`
	A10  string `json:"a10"`
}

type A struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	T    B      `json:"b"`
}

func sortMarshal() {

	a := A{
		Name: "name",
		Age:  30,
		T: B{
			Name: "b",
			A01:  1,
			A10:  "a10",
		},
	}
	//json.FieldSortType = json.FieldSortDesc
	json.FieldSortType = json.FieldSortAsc
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("marshal failed,err:%s\n", err.Error())
		return
	}
	fmt.Printf("data:%s\n", data)
}

func main() {
	sortMarshal()
	TestBaseMsg()
}
