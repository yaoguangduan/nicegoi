package ui

import (
	"fmt"
	"reflect"
	"strings"
)

type Table struct {
	*valuedWidget
}

func NewTable(data interface{}) *Table {
	cols, realData := parseColsAndData(data)
	if cols == nil || realData == nil {
		panic("invalid data format")
	}
	t := &Table{newValuedWidget("table", realData)}
	t.e.Set("columns", cols)
	return t
}

func parseColsAndData(data interface{}) ([]map[string]string, interface{}) {
	var dt = reflect.ValueOf(data)
	if dt.Kind() == reflect.Interface {
		dt = dt.Elem()
	}
	if dt.Kind() == reflect.Ptr {
		dt = dt.Elem()
	}
	if dt.Kind() == reflect.Struct {
		var typ = reflect.TypeOf(data)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		return getColFromStruct(typ), []interface{}{data}
	} else if dt.Kind() == reflect.Slice {
		first := dt.Index(0)
		if first.Kind() == reflect.Ptr {
			first = first.Elem()
		}
		if first.Kind() == reflect.Interface {
			first = first.Elem()
		}
		if first.Kind() == reflect.Ptr {
			first = first.Elem()
		}
		if first.Kind() == reflect.Struct {
			return getColFromStruct(first.Type()), data
		} else if first.Kind() == reflect.Slice {
			cols := make([]map[string]string, 0)
			for i := 0; i < first.Len(); i++ {
				cell := fmt.Sprintf("%s", first.Index(i))
				cols = append(cols, map[string]string{"colKey": cell, "title": cell})
			}
			realData := make([]map[string]string, 0)
			for i := 1; i < dt.Len(); i++ {
				row := dt.Index(i)
				rowData := make(map[string]string)
				for j := 0; j < row.Len(); j++ {
					if first.Len() > j {
						rowData[fmt.Sprintf("%s", first.Index(j))] = fmt.Sprintf("%s", row.Index(j))
					}
				}
				realData = append(realData, rowData)
			}
			return cols, realData
		} else if first.Kind() == reflect.Map {
			cols := make([]map[string]string, 0)
			for _, key := range first.MapKeys() {
				cell := fmt.Sprintf("%s", key.Interface())
				cols = append(cols, map[string]string{"colKey": cell, "title": cell})
			}
			return cols, data
		}
	} else if dt.Kind() == reflect.Map {
		mv := dt.MapIndex(dt.MapKeys()[0])
		if mv.Kind() == reflect.Ptr {
			mv = mv.Elem()
		}
		if mv.Kind() == reflect.Slice {
			cols := make([]map[string]string, 0)
			realData := make([]map[string]string, 0)
			var maxLen = -1
			for _, k := range dt.MapKeys() {
				cell := fmt.Sprintf("%s", k)
				cols = append(cols, map[string]string{"colKey": cell, "title": cell})
				maxLen = dt.MapIndex(k).Len()
			}
			for i := 0; i < maxLen; i++ {
				row := make(map[string]string)
				for _, k := range dt.MapKeys() {
					valSlice := dt.MapIndex(k)
					cell := fmt.Sprintf("%s", k)
					if i < valSlice.Len() {
						v := fmt.Sprintf("%s", valSlice.Index(i))
						row[cell] = v
					} else {
						row[cell] = ""
					}
				}
				realData = append(realData, row)
			}
			return cols, realData

		}
	}
	return nil, nil
}

func getColFromStruct(typ reflect.Type) []map[string]string {
	cols := make([]map[string]string, 0)
	for i := range typ.NumField() {
		f := typ.Field(i)
		col := f.Tag.Get("json")
		if strings.Contains(col, ",") {
			col = strings.Split(col, ",")[0]
		}
		cols = append(cols, map[string]string{"colKey": col, "title": col})
	}
	return cols
}
