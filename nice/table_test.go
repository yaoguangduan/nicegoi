package nice

import (
	"encoding/json"
	"fmt"
	"github.com/yaoguangduan/nicegoi/nice/option/menu"
	"testing"
)

func TestData(t *testing.T) {
	data := [][]string{{"name", "age"}, {"alen", "12"}, {"john", "290825432763813"}}
	cols, datas := parseColsAndData(data)
	fmt.Println(cols, datas)

	data1 := menu.NewItem("label1", "value1")
	cols, datas = parseColsAndData(data1)
	fmt.Println(toJson(cols), toJson(datas))

	data2 := make(map[string][]interface{})
	data2["name"] = []interface{}{"jjj", "AAA", "CCC"}
	data2["label"] = []interface{}{"food", "city", "less"}
	cols, datas = parseColsAndData(data2)
	fmt.Println(toJson(cols), toJson(datas))

	data3 := []interface{}{menu.NewItem("label1", "value1"), menu.NewItem("label2", "value2")}

	cols, datas = parseColsAndData(data3)
	fmt.Println(toJson(cols), toJson(datas))
}

func toJson(obj interface{}) string {
	js, _ := json.Marshal(obj)
	return string(js)
}
