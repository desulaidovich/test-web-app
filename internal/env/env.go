package env

import (
	"fmt"
	"iter"
	"os"
	"reflect"
	"strconv"
)

const (
	tagName = "env"
	// defaultTagName = "env-default"
)

func Read[T any](t *T) error {
	for field, tag := range fields(t) {
		env := os.Getenv(tag)

		switch field.Kind() {
		case reflect.String:
			field.SetString(env)
		case reflect.Bool:
			value, err := strconv.ParseBool(env)
			if err != nil {
				return err
			}

			field.SetBool(value)
		case
			reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			value, err := strconv.ParseInt(env, 10, field.Type().Bits())
			if err != nil {
				return err
			}

			field.SetInt(value)
		case
			reflect.Float32,
			reflect.Float64:
			value, err := strconv.ParseFloat(env, field.Type().Bits())
			if err != nil {
				return err
			}

			field.SetFloat(value)
		default:
			return fmt.Errorf("unsupported type: %s", field.Kind())
		}
	}
	return nil
}

func fields[T any](t *T) iter.Seq2[reflect.Value, string] {
	return func(yield func(reflect.Value, string) bool) {
		e := reflect.ValueOf(t).Elem()

		for num := range e.NumField() {
			f := e.Field(num)

			if !f.CanSet() {
				continue
			}

			if !f.IsValid() {
				continue
			}

			tag := e.Type().Field(num).Tag.Get(tagName)

			if !yield(f, tag) {
				return
			}
		}
	}
}
