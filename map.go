package nullable

import (
	"reflect"
)

type MaybeNullable interface {
	isNullableStruct() bool
}

func CreateEntityMap(entity interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(entity)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	maybeNullable := reflect.TypeOf((*MaybeNullable)(nil)).Elem()

	entityType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldRValue := v.Field(i)
		fieldRType := entityType.Field(i)
		field := v.Field(i).Interface()

		if fieldRValue.Type().Implements(maybeNullable) && fieldRValue.FieldByName("isSet").Bool() {
			result[fieldRType.Name] = field
		}
	}

	return result
}
