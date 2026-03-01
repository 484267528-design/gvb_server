package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetValidMsg(err error, obj any) string {
	get0bj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exits := get0bj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
