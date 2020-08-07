/**
* @Auther:gy
* @Date:2020/8/7 17:46
 */

package sort

import (
	"encoding/json"
	"github.com/spf13/cast"
	"strings"
)

/*
type People struct {
	Age  int
	Name string
	Sex  bool
}
p1 := People{10, "zhangsan", false}
	p2 := People{20, "lisi", true}
	p3 := People{16, "wangwu", false}
	p4 := People{18, "zhaoliu", true}
	var ps []People
	ps = append(ps, p1)
	ps = append(ps, p2)
	ps = append(ps, p3)
	ps = append(ps, p4)
	SortStruct("Age", true, "int", &ps)
	fmt.Println(ps)
	SortStruct("Age", false, "int", &ps)
	fmt.Println(ps)
	SortStruct("Name", true, "string", &ps)
	fmt.Println(ps)
	SortStruct("Sex", true, "bool", &ps)
	fmt.Println(ps)
*/
//排序字段json字符串、排序方式、排序字段类型(用反射比较麻烦)、需要排序的slice传指针
//以下排序均使用冒泡排序
func SortStruct(Column string, Asc bool, Type string, Value interface{}) error {
	Type = strings.ToLower(Type)
	var ms []map[string]interface{}
	b, err := json.Marshal(Value)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &ms)
	if err != nil {
		return err
	}
	l := len(ms)
	switch Type {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64":

		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				if Asc {
					if cast.ToFloat64(ms[i][Column]) > cast.ToFloat64(ms[j][Column]) {
						a := ms[i]
						ms[i] = ms[j]
						ms[j] = a
					}
				} else {
					if cast.ToFloat64(ms[i][Column]) < cast.ToFloat64(ms[j][Column]) {
						a := ms[i]
						ms[i] = ms[j]
						ms[j] = a
					}
					//fmt.Println(ms)
				}

			}
		}
	case "string":
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				if Asc {
					if cast.ToString(ms[i][Column]) > cast.ToString(ms[j][Column]) {
						a := ms[i]
						ms[i] = ms[j]
						ms[j] = a
					}
				} else {
					if cast.ToString(ms[i][Column]) < cast.ToString(ms[j][Column]) {
						a := ms[i]
						ms[i] = ms[j]
						ms[j] = a
					}
				}
			}
		}
	case "bool":
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				if Asc {
					if !cast.ToBool(ms[i][Column]) && cast.ToBool(ms[j][Column]) {
						a := ms[i]
						ms[i] = ms[j]
						ms[j] = a
					}
				} else {
					if cast.ToBool(ms[i][Column]) && !cast.ToBool(ms[j][Column]) {
						a := ms[i]
						ms[i] = ms[j]
						ms[j] = a
					}
				}
			}
		}
	}
	b, err = json.Marshal(ms)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Value)
}
