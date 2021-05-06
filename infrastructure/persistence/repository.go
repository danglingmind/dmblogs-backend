package persistence

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Row standard query result format
type Row map[string]interface{}

// All datastores will implement Database interface
type Database interface {
	Open(ctx context.Context, host, username, password, dbname string, port int) error
	Close(ctx context.Context) error
	QueryRow(ctx context.Context, q string, params interface{}) (Row, error)
	Query(ctx context.Context, q string, params ...interface{}) ([]Row, error)
	Save(ctx context.Context, q string, params ...interface{}) error
}

func (qr Row) Serialize(s interface{}) error {
	for key, value := range qr {
		// inside the loop because after setting up every field the layout of the bits
		// will be changed
		structValue := reflect.ValueOf(s)
		structFieldValue := structValue.FieldByName(key)
		fmt.Println(structFieldValue.Interface())

		if !structFieldValue.IsValid() {
			return fmt.Errorf("no such field in struct: %s", key)
		}
		if !structFieldValue.CanSet() {
			return fmt.Errorf("cannot set %s field value", key)
		}
		structFieldType := structFieldValue.Type()
		val := reflect.ValueOf(value)
		fmt.Println(structFieldType)
		fmt.Println(val.Type())
		if structFieldType != val.Type() {
			return errors.New("value type do not match for field:" + key)
		}

		structFieldValue.Set(val)
	}
	return nil
}

func (qr Row) Serialize2(s interface{}) error {
	for key, value := range qr {
		// inside the loop because after setting up every field the layout of the bits
		// will be changed
		valStruct := reflect.ValueOf(s).Elem()
		// typeStruct := reflect.TypeOf(s)

		// get the value of field from Row
		v := reflect.ValueOf(value)

		if valStruct.Kind() == reflect.Struct {
			field := valStruct.FieldByName(key)

			if !field.IsValid() {
				return fmt.Errorf("serialize: field %s is not valid", key)
			}
			if !field.CanSet() {
				return fmt.Errorf("serialize: cannot set %s field value", key)
			}

			// get the type of the field
			typeOfField := field.Type().String()

			// convert on the basis of type of the field
			switch typeOfField {
			case "string":
				valueToAssign := string(v.Interface().([]uint8))
				field.Set(reflect.ValueOf(valueToAssign))
				break
			case "int":
				valueToAssign := int(v.Interface().(int64))
				field.Set(reflect.ValueOf(valueToAssign))
				break
			case "int64":
				valueToAssign := int64(v.Interface().(int64))
				field.Set(reflect.ValueOf(valueToAssign))
				break
			case "bool":
				valueToAssign := int(v.Interface().(int64))
				if valueToAssign == 1 {
					field.Set(reflect.ValueOf(true))
				} else {
					field.Set(reflect.ValueOf(false))
				}
				break
			case "time.Time":
				valueToAssign := time.Time(v.Interface().(time.Time))
				field.Set(reflect.ValueOf(valueToAssign))
			}
		} else {
			return fmt.Errorf("not a struct")
		}
	}
	return nil
}
