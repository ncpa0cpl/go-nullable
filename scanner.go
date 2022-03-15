package nullable

import (
	"database/sql"
	"reflect"
)

func (v *Nullable[T]) Scan(value interface{}) error {
	switch reflect.TypeOf(v.content).String() {
	case "int":
		ns := sql.NullInt64{}
		if err := ns.Scan(value); err != nil {
			return err
		}
		if !ns.Valid {
			v.isSet = false
		} else {
			v.isSet = true
			return v.set_unsafe(int(ns.Int64))
		}
	case "int16":
		ns := sql.NullInt16{}
		if err := ns.Scan(value); err != nil {
			return err
		}
		if !ns.Valid {
			v.isSet = false
		} else {
			v.isSet = true
			return v.set_unsafe(ns.Int16)
		}
	case "int32":
		ns := sql.NullInt32{}
		if err := ns.Scan(value); err != nil {
			return err
		}
		if !ns.Valid {
			v.isSet = false
		} else {
			v.isSet = true
			return v.set_unsafe(ns.Int32)
		}
	case "int64":
		ns := sql.NullInt64{}
		if err := ns.Scan(value); err != nil {
			return err
		}
		if !ns.Valid {
			v.isSet = false
		} else {
			v.isSet = true
			return v.set_unsafe(ns.Int64)
		}
	case "float32":
		ns := sql.NullFloat64{}
		if err := ns.Scan(value); err != nil {
			return err
		}
		if !ns.Valid {
			v.isSet = false
		} else {
			v.isSet = true
			return v.set_unsafe(float32(ns.Float64))
		}
	case "float64":
		ns := sql.NullFloat64{}
		if err := ns.Scan(value); err != nil {
			return err
		}
		if !ns.Valid {
			v.isSet = false
		} else {
			v.isSet = true
			return v.set_unsafe(ns.Float64)
		}
	case "string":
		ns := sql.NullString{}
		if err := ns.Scan(value); err != nil {
			return err
		}
		if !ns.Valid {
			v.isSet = false
		} else {
			v.isSet = true
			return v.set_unsafe(ns.String)
		}
	case "bool":
		ns := sql.NullBool{}
		if err := ns.Scan(value); err != nil {
			return err
		}
		if !ns.Valid {
			v.isSet = false
		} else {
			v.isSet = true
			return v.set_unsafe(ns.Bool)
		}

	}

	return nil
}
