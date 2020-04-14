package sequenceutils

import "reflect"

// IsIn 判断val是否在obj中，obj可以是map，array，struct；如果obj是map或struct，则val为key；如果obj是array，则val为元素。
//  - obj: map、array或者struct类型的引用
//  - val: 带判断的元素或者key
func IsIn(obj interface{}, val interface{}) (isIn bool) {
	pField := reflect.ValueOf(obj)
	vField := pField.Elem()
	if vField.Kind() == reflect.Array || vField.Kind() == reflect.Slice {
		for i := 0; i < vField.Len(); i++ {
			if vField.Index(i).Interface() == val {
				isIn = true
				return
			}
		}
		return
	}
	if vField.Kind() == reflect.Struct {
		fieldName := val.(string)
		if _, isIn = vField.Type().FieldByName(fieldName); isIn {
			return
		}
		if _, isIn = vField.Type().MethodByName(fieldName); isIn {
			return
		}
		return
	}
	if vField.Kind() == reflect.Map {
		for _, v := range vField.MapKeys() {
			if v.Interface() == val {
				isIn = true
				return
			}
		}
		return
	}
	return
}

// DelFrom 通过索引从数组中删除元素
//  - container: 容器型变量的引用，array、slice、map或者string类型
//  - key: 待删除元素的索引
func DelFrom(container interface{}, key interface{}) {
	vField := reflect.ValueOf(container)
	value := vField.Elem()
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		index := key.(int)
		if index < 0 || index+1 > value.Len() {
			return
		}
		result := reflect.AppendSlice(value.Slice(0, index), value.Slice(index+1, value.Len()))
		value.Set(result)
		return
	}
	if value.Kind() == reflect.Map {
		for _, mapKey := range value.MapKeys() {
			if mapKey.Interface() == key {
				value.SetMapIndex(mapKey, reflect.Value{})
			}
		}
		return
	}
	if value.Kind() == reflect.String {
		index := key.(int)
		if index < 0 || index+1 > value.Len() {
			return
		}
		ps := container.(*string)
		result := (*ps)[:index] + (*ps)[index+1:]
		value.SetString(result)
		return
	}
}
