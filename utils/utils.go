package utils

import "reflect"

// func GetJsonTagsWithValues(obj interface{}) map[string]reflect.StructField {
// 	val := reflect.TypeOf(obj)
// 	result := make(map[string]reflect.StructField)
// 	for i := 0; i < val.NumField(); i++ {
// 		fieldJsonName := val.Field(i).Tag.Get("json")
// 		if fieldJsonName != "" {
// 			fieldName := val.Field(i).Name
// 			if fieldValue, ok := val.FieldByName(fieldName); ok {
// 				result[fieldJsonName] = fieldValue
// 			}
// 		}

// 	}
// 	return result
// }

func GetJsonTagsWithValues2(obj interface{}) map[string]interface{} {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	result := make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {
		fieldJsonName := typ.Field(i).Tag.Get("json")
		if fieldJsonName != "" {
			fieldName := typ.Field(i).Name
			fieldValue := val.FieldByName(fieldName)
			result[fieldJsonName] = fieldValue.Interface()
		}

	}
	return result
}
