package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/mitchellh/reflectwalk"
	"reflect"
)

type Values struct {
	Extra map[string]string
}

type WalkMap struct {
	MapVal reflect.Value
	Keys   map[string]bool
	Values map[string]bool
}

func (t *WalkMap) Map(m reflect.Value) error {
	t.MapVal = m
	return nil
}

func (t *WalkMap) MapElem(m, k, v reflect.Value) error {
	if t.Keys == nil {
		t.Keys = make(map[string]bool)
		t.Values = make(map[string]bool)
	}

	t.Keys[k.Interface().(string)] = true
	t.Values[v.Interface().(string)] = true
	return nil
}

func main() {
	w := new(WalkMap)

	type S struct {
		Map map[string]string
	}

	data := &S{
		Map: map[string]string{
			"v1": "1v",
			"v2": "2v",
			"v3": "3v",
			"v4": "4v",
		},
	}

	err := reflectwalk.Walk(data, w)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(w.MapVal)

	var result Values
	if err := mapstructure.Decode(w.MapVal, &result); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", result)
}
