package catav15

import (
	"reflect"
	"strings"
)

// getFieldType allows us to dynamically take the database column name
// a user is trying to manipulate and find out what it's real value is, which
// means we can then type cast what the user gives us
func getFieldType(field string, table interface{}) reflect.Type {
	fields := reflect.VisibleFields(reflect.TypeOf(table))
	for _, f := range fields {

		if f.Name == field {
			return f.Type
		}
	}

	return nil
}

func columnKeyToStructKey(key string) string {
	s := strings.Split(key, "_")
	for si, ss := range s {
		s[si] = strings.Title(ss)
	}
	return strings.Join(s, "")
}
