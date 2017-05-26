package msg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/wlibo666/json"
)

var (
	ERR_ELEM_CANNOT_SET = errors.New("Elem can not set")
)

type BaseMsg struct {
	Str_1  string `json:"str_1,omitempty"`
	Str_2  string `json:"str_2,omitempty"`
	Str_3  string `json:"str_3,omitempty"`
	Str_4  string `json:"str_4,omitempty"`
	Str_5  string `json:"str_5,omitempty"`
	Str_6  string `json:"str_6,omitempty"`
	Str_7  string `json:"str_7,omitempty"`
	Str_8  string `json:"str_8,omitempty"`
	Str_9  string `json:"str_9,omitempty"`
	Str_10 string `json:"str_10,omitempty"`
	Str_11 string `json:"str_11,omitempty"`
	Str_12 string `json:"str_12,omitempty"`
	Str_13 string `json:"str_13,omitempty"`
	Str_14 string `json:"str_14,omitempty"`
	Str_15 string `json:"str_15,omitempty"`
	Str_16 string `json:"str_16,omitempty"`
	Str_17 string `json:"str_17,omitempty"`
	Str_18 string `json:"str_18,omitempty"`
	Str_19 string `json:"str_19,omitempty"`
	Str_20 string `json:"str_20,omitempty"`
	Str_21 string `json:"str_21,omitempty"`
	Str_22 string `json:"str_22,omitempty"`
	Str_23 string `json:"str_23,omitempty"`
	Str_24 string `json:"str_24,omitempty"`
	Str_25 string `json:"str_25,omitempty"`
	Str_26 string `json:"str_26,omitempty"`
	Str_27 string `json:"str_27,omitempty"`
	Str_28 string `json:"str_28,omitempty"`
	Str_29 string `json:"str_29,omitempty"`
	Str_30 string `json:"str_30,omitempty"`

	Int64_1 int64 `json:"int64_1,omitempty"`
	Int64_2 int64 `json:"int64_2,omitempty"`
	Int64_3 int64 `json:"int64_3,omitempty"`
	Int64_4 int64 `json:"int64_4,omitempty"`
	Int64_5 int64 `json:"int64_5,omitempty"`

	Int_1 int `json:"int_1,omitempty"`
	Int_2 int `json:"int_2,omitempty"`
	Int_3 int `json:"int_3,omitempty"`
	Int_4 int `json:"int_4,omitempty"`
	Int_5 int `json:"int_5,omitempty"`

	Int32_1 int32 `json:"int32_1,omitempty"`
	Int32_2 int32 `json:"int32_2,omitempty"`
	Int32_3 int32 `json:"int32_3,omitempty"`
	Int32_4 int32 `json:"int32_4,omitempty"`
	Int32_5 int32 `json:"int32_5,omitempty"`
}

func setValue(p interface{}, fieldNameValues map[string]string) error {
	tmpValue := reflect.ValueOf(p)
	if !tmpValue.Elem().CanSet() {
		return ERR_ELEM_CANNOT_SET
	}
	tmpValue = tmpValue.Elem()
	for k, v := range fieldNameValues {
		tmpV := tmpValue.FieldByName(k)
		if tmpV.Kind() == reflect.Invalid || !tmpV.IsValid() || !tmpV.CanSet() {
			continue
		}

		switch tmpV.Kind() {
		case reflect.String:
			tmpV.SetString(v)
		case reflect.Int64, reflect.Int, reflect.Int32:
			var vInt64 int64
			fmt.Sscanf(v, "%d", &vInt64)
			tmpV.SetInt(vInt64)
		case reflect.Bool:
			if v == "true" || v == "1" {
				tmpV.SetBool(true)
			} else {
				tmpV.SetBool(false)
			}
		}
	}
	return nil
}

func (p *BaseMsg) Marshal(fieldNameAlias ...map[string]string) ([]byte, error) {
	if len(fieldNameAlias) > 0 {
		tmpNameAlias := make(map[string]string)
		for k, v := range fieldNameAlias[0] {
			tmpNameAlias[strings.ToLower(k)] = v
		}
		return json.Marshal(p, tmpNameAlias)
	}
	return json.Marshal(p)
}

func (p *BaseMsg) SetValue(fieldNameValues map[string]string) error {
	return setValue(p, fieldNameValues)
}
