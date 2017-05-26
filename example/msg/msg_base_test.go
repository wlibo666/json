package msg

import (
	"testing"
)

func TestBaseMsg(t *testing.T) {
	msg := &BaseMsg{}

	fieldNameValues := make(map[string]string)
	fieldNameValues["Str_1"] = "wangchunyan"
	fieldNameValues["Int_1"] = "29"

	err := msg.SetValue(fieldNameValues)
	if err != nil {
		t.Fatalf("set value failed:%s", err.Error())
	}

	data, err := msg.Marshal()
	if err != nil {
		t.Fatalf("marshal not map failed,err:%s", err.Error())
	}
	t.Logf("not map data:%s", string(data))

	fieldNameAlias := make(map[string]string)
	fieldNameAlias["Str_1"] = "name"
	fieldNameAlias["Int_1"] = "age"

	data, err = msg.Marshal(fieldNameAlias)
	if err != nil {
		t.Fatalf("marshal map failed,err:%s", err.Error())
	}
	t.Logf("map data:%s", string(data))
}
