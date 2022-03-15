package nullable

import (
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func (v Nullable[T]) GormDataType() string {
	switch reflect.TypeOf(v.content).String() {
	case "int":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		return "int"
	case "float32":
		fallthrough
	case "float64":
		return "float"
	case "bool":
		return "bool"
	case "string":
		return "string"
	}

	return "bytes"
}

func (v Nullable[T]) GormDBDataType(db *gorm.DB, f *schema.Field) string {
	switch reflect.TypeOf(v.content).String() {
	case "int":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		return "int"
	case "float32":
		fallthrough
	case "float64":
		return "float"
	case "bool":
		return "bool"
	case "string":
		return "string"
	}

	return "bytes"
}
