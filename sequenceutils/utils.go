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
	pField := reflect.ValueOf(container)
	vField := pField.Elem()
	if vField.Kind() == reflect.Slice || vField.Kind() == reflect.Array {
		index := key.(int)
		if index < 0 || index+1 > vField.Len() {
			return
		}
		result := reflect.AppendSlice(vField.Slice(0, index), vField.Slice(index+1, vField.Len()))
		vField.Set(result)
		return
	}
	if vField.Kind() == reflect.Map {
		for _, mapKey := range vField.MapKeys() {
			if mapKey.Interface() == key {
				vField.SetMapIndex(mapKey, reflect.Value{})
			}
		}
		return
	}
	if vField.Kind() == reflect.String {
		index := key.(int)
		if index < 0 || index+1 > vField.Len() {
			return
		}
		ps := container.(*string)
		result := (*ps)[:index] + (*ps)[index+1:]
		vField.SetString(result)
		return
	}
}

// SliceSeq 对有序变量进行切片
//  - seq: 有序变量，string、array或slice类型
//  - startIdx: 首位置索引
//  - endIdx: 末位置索引
func SliceSeq(seq interface{}, startIdx, endIdx int) (result interface{}) {
	vField := reflect.ValueOf(seq)
	var length int
	vKind := vField.Kind()
	if vKind == reflect.Array || vKind == reflect.Slice || vKind == reflect.String {
		length = vField.Len()
		if startIdx < 0 {
			startIdx = length + startIdx
		}
		if endIdx < 0 {
			endIdx = length + endIdx
		}
		if startIdx < 0 {
			startIdx = 0
		}
		if endIdx > length {
			endIdx = length
		}
	}
	if vKind == reflect.Array || vKind == reflect.Slice {
		// 原切片长度小于等于0，或者截取长度小于等于0，则直接返回空切片
		if length <= 0 || endIdx-startIdx <= 0 {
			result = reflect.New(reflect.TypeOf(seq)).Elem().Interface()
			return
		}
		result = vField.Slice(startIdx, endIdx).Interface()
		return
	}
	if vKind == reflect.String {
		if length <= 0 || endIdx-startIdx <= 0 {
			result = ""
			return
		}
		seqStr := seq.(string)
		result = seqStr[startIdx:endIdx]
		return
	}
	return
}
