package attr

import (
	"fmt"
	"strconv"
	"strings"
)

type Value struct {
	v interface{}
}

func NewValue(v interface{}) *Value {
	return &Value{v: v}
}

func (v *Value) Set(val interface{}) {
	v.v = val
}

func (v *Value) Interface() interface{} {
	return v.v
}

func (v *Value) Type() string {
	return fmt.Sprintf("%T", v.v)
}

func (v *Value) String() string {
	if vv, ok := v.v.(Attribute); ok {
		return vv.Value().String()
	}

	if vv, ok := v.v.(fmt.Stringer); ok {
		return vv.String()
	}

	switch vv := v.v.(type) {
	case string:
		return vv
	case []string:
		return strings.Join(vv, ",")
	case int:
		return strconv.Itoa(vv)
	case []int:
		var ss []string
		for _, i := range vv {
			ss = append(ss, strconv.Itoa(i))
		}

		return strings.Join(ss, ",")
	case bool:
		return strconv.FormatBool(vv)
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	case []float64:
		var ss []string
		for _, f := range vv {
			ss = append(ss, strconv.FormatFloat(f, 'f', -1, 64))
		}

		return strings.Join(ss, ",")
	default:
		return fmt.Sprintf("%v", vv)
	}
}
