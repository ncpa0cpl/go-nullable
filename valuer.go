package nullable

import (
	"database/sql/driver"
	"errors"
	"reflect"
)

func (v Nullable[T]) Value() (driver.Value, error) {
	if !v.isSet {
		return nil, nil
	}

	t := reflect.TypeOf(v.content).String()

	if t == "int" || t == "int16" || t == "int32" || t == "int64" {
		var i64 int64
		var err error = nil
		var ok bool

		switch t {
		case "int":
			var i int
			i, ok = v.asInterface().(int)
			i64 = int64(i)
		case "int16":
			var i int16
			i, ok = v.asInterface().(int16)
			i64 = int64(i)
		case "int32":
			var i int32
			i, ok = v.asInterface().(int32)
			i64 = int64(i)
		case "int64":
			i64, ok = v.asInterface().(int64)
		}

		if !ok {
			err = errors.New("Invalid type conversion.")
		}

		return i64, err
	} else if t == "float32" || t == "float64" {
		var f64 float64
		var err error = nil
		var ok bool

		switch t {
		case "float32":
			var f float32
			f, ok = v.asInterface().(float32)
			f64 = float64(f)
		case "float64":
			f64, ok = v.asInterface().(float64)
		}

		if !ok {
			err = errors.New("Invalid type conversion.")
		}

		return f64, err
	} else if t == "bool" {
		var err error = nil

		b, ok := v.asInterface().(bool)

		if !ok {
			err = errors.New("Invalid type conversion.")
		}

		return b, err
	} else if t == "string" {
		var err error = nil

		str, ok := v.asInterface().(string)

		if !ok {
			err = errors.New("Invalid type conversion.")
		}

		return str, err
	}

	return nil, errors.New("Unrecognized type.")
}
