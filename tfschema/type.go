package tfschema

import (
	"fmt"
	"reflect"
)

type Type struct {
	// T is an instance of github.com/hashicorp/terraform/vendor/github.com/zclconf/go-cty.Type
	ctyType interface{}
}

func NewType(t interface{}) *Type {
	return &Type{
		ctyType: t,
	}
}

func (t *Type) MarshalJSON() ([]byte, error) {
	v := reflect.ValueOf(t.ctyType).MethodByName("FriendlyName")
	if !v.IsValid() {
		return nil, fmt.Errorf("Faild to find FriendlyName(): %#v", t)
	}

	nv := v.Call([]reflect.Value{})
	if len(nv) == 0 {
		return nil, fmt.Errorf("Faild to call FriendlyName(): %#v", v)
	}

	friendlyName := nv[0].String()

	return []byte(`"` + friendlyName + `"`), nil
}
