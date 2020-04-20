package sequenceutils

import (
	"reflect"
	"strings"
)

// IsIn 判断val是否在obj中，obj可以是map，array，struct；如果obj是map或struct，则val为key；如果obj是array，则val为元素。
//  - obj: map、array或者struct类型的引用
//  - val: 带判断的元素或者key
func IsIn(obj interface{}, val interface{}) (isIn bool) {
	pField := reflect.ValueOf(obj)
	vField := pField.Elem()
	vKind := vField.Kind()
	if vKind == reflect.Array || vKind == reflect.Slice {
		for i := 0; i < vField.Len(); i++ {
			if vField.Index(i).Interface() == val {
				isIn = true
				return
			}
		}
		return
	}
	if vKind == reflect.Struct {
		fieldName := val.(string)
		if _, isIn = vField.Type().FieldByName(fieldName); isIn {
			return
		}
		if _, isIn = vField.Type().MethodByName(fieldName); isIn {
			return
		}
		return
	}
	if vKind == reflect.Map {
		for _, v := range vField.MapKeys() {
			if v.Interface() == val {
				isIn = true
				return
			}
		}
		return
	}
	if vKind == reflect.String {
		return strings.Contains(obj.(string), val.(string))
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
		length := vField.Len()
		index := key.(int)
		if index < 0 {
			index += length
		}
		if index < 0 || index+1 > length {
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
		length := vField.Len()
		if index < 0 {
			index += length
		}
		if index < 0 || index+1 > length {
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

// Slice 有序变量切片函数，支持按步长进行切片，切片逆序，从末尾开始计算下下标等操作。
//  - seq: 有序变量，string、array或slice类型
//  - params: 切片操作参数，第一个值为首位置索引，第二个值为末位置索引，第三个值为步长。
func Slice(seq interface{}, params ...int) (result interface{}) {
	var startIdx, endIdx, step, length int
	paramSize := len(params)
	if paramSize == 0 {
		return seq
	}
	vField := reflect.ValueOf(seq)
	vKind := vField.Kind()
	// 处理下标及步长
	if vKind == reflect.Array || vKind == reflect.Slice || vKind == reflect.String {
		length = vField.Len()
		switch paramSize {
		case 1:
			startIdx = params[0]
			endIdx = length
			step = 1
		case 2:
			startIdx = params[0]
			endIdx = params[1]
			step = 1
		default:
			startIdx = params[0]
			endIdx = params[1]
			step = params[2]
			if step == 0 {
				return seq
			}
		}
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
	// 处理切片及数组
	if vKind == reflect.Array || vKind == reflect.Slice {
		// 原切片长度小于等于0，或者截取长度小于等于0，则直接返回空切片
		tmpResult := reflect.New(reflect.TypeOf(seq)).Elem()
		if length <= 0 || endIdx-startIdx <= 0 {
			return tmpResult.Interface()
		}
		switch {
		case step == 1:
			tmpResult = vField.Slice(startIdx, endIdx)
		case step < 0:
			for i := endIdx - 1; i >= startIdx; i += step {
				tmpResult = reflect.Append(tmpResult, vField.Index(i))
			}
		default:
			for i := startIdx; i < endIdx; i += step {
				tmpResult = reflect.Append(tmpResult, vField.Index(i))
			}
		}
		return tmpResult.Interface()
	}
	// 处理字符串
	if vKind == reflect.String {
		var tmpResult string
		if length <= 0 || endIdx-startIdx <= 0 {
			return tmpResult
		}
		seqStr := seq.(string)
		switch {
		case step == 1:
			tmpResult = seqStr[startIdx:endIdx]
		case step < 0:
			for i := endIdx - 1; i >= startIdx; i += step {
				tmpResult += string(seqStr[i])
			}
		default:
			for i := startIdx; i < endIdx; i += step {
				tmpResult += string(seqStr[i])
			}
		}
		result = tmpResult
		return
	}
	return seq
}

// IfElse 三元表达函数
//  - condition: 条件
//  - a: condition为true时返回此值
//  - b: condition为false时返回此值
func IfElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
