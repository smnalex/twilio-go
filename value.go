package twilio

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// Values converts a struct into url.Values
func Values(v interface{}) url.Values {
	values := make(url.Values)
	if v == nil {
		return values
	}
	rval := reflect.ValueOf(v)
	if rval.Kind() != reflect.Struct {
		return nil
	}
	for rval.Kind() == reflect.Ptr {
		if rval.IsNil() {
			return values
		}
		rval = rval.Elem()
	}

	parseValues(values, rval, "")
	return values
}

func parseValues(values url.Values, val reflect.Value, scope string) {
	var embedded []reflect.Value

	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		sv := val.Field(i)
		tag := sf.Tag.Get("url")
		if tag == "-" {
			continue
		}
		name, opts := parseTag(tag)
		if name == "" {
			if sf.Anonymous && sv.Kind() == reflect.Struct {
				embedded = append(embedded, sv)
				continue
			}

			name = sf.Name
		}

		if scope != "" {
			name = scope + "." + name
		}

		if opts.Contains("omitempty") && isEmptyValue(sv) {
			continue
		}

		if sv.Kind() == reflect.Slice || sv.Kind() == reflect.Array {
			for i := 0; i < sv.Len(); i++ {
				values.Add(name, valueString(sv.Index(i), opts))
			}
			continue
		}

		for sv.Kind() == reflect.Ptr {
			if sv.IsNil() {
				break
			}
			sv = sv.Elem()
		}

		if sv.Kind() == reflect.Struct {
			parseValues(values, sv, name)
			continue
		}

		values.Add(name, valueString(sv, opts))
	}

	for _, f := range embedded {
		parseValues(values, f, scope)
	}
}

func valueString(v reflect.Value, opts tagOptions) string {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	return fmt.Sprint(v.Interface())
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return false
}

type tagOptions []string

func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

func (o tagOptions) Contains(option string) bool {
	for _, s := range o {
		if s == option {
			return true
		}
	}
	return false
}
