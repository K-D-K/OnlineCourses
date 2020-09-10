package validator

import (
	"errors"
	"reflect"
)

// ValidateOnCreate validates does any direct and nested record have PK id
func ValidateOnCreate(value interface{}) error {
	return validate(reflect.ValueOf(value))
}

func validate(value reflect.Value) error {
	if isIterable(value) {
		return validateArray(value)
	} else if value.Kind() == reflect.Struct {
		return validateStruct(value)
	}
	return nil
}

func validateArray(value reflect.Value) error {
	for i := 0; i < value.Len(); i++ {
		err := validate(reflect.Indirect(value.Index(i)))
		if err != nil {
			return err
		}
	}
	return nil
}

func validateStruct(value reflect.Value) error {
	err := checkPkIDOnCreate(value)
	if err == nil {
		valueType := value.Type()
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)

			/*
				As for now there is no use case to validate Anonymous fields.
				If we found any case means we need to fetch all fields of the struct including Anonymous in a straight order instead of nesting one
			*/
			if !valueType.Field(i).Anonymous {
				err = validate(reflect.ValueOf(field.Interface()))
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}

func checkPkIDOnCreate(value reflect.Value) error {
	method := value.MethodByName("GetPKID")
	if method.IsValid() {
		// If GET PK returns nil means it is new record or else record is going to update
		if !method.Call([]reflect.Value{})[0].IsNil() {
			return errors.New("PKId found on record create validation")
		}
	}
	return nil
}

// redundant function
func canValidate(value reflect.Value) bool {
	method := value.MethodByName("ValidateOnCreate")
	return method.IsValid() && method.Call([]reflect.Value{})[0].Bool()
}

func isIterable(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}
