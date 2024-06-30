package migrations

import (
	"reflect"
	"strings"
)

func createTableSQL(model interface{}) string {
	v := reflect.ValueOf(model).Elem()
	t := v.Type()
	tableName := strings.ToLower(t.Name())

	var columns []string


}