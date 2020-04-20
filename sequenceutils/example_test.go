package sequenceutils

import "fmt"

func ExampleIsIn() {
	testArray := []int{1, 2, 3, 4, 5}
	IsIn(&testArray, 3)   // true
	IsIn(&testArray, "3") // false
	IsIn(&testArray, 9)   // false
	testStruct := struct {
		Name string
	}{}
	IsIn(&testStruct, "Name") // true
	IsIn(&testStruct, "node") // false
}

func ExampleDelFrom() {
	testArray := []int{1, 2, 3, 4, 5}
	DelFrom(&testArray, 3) // []int{1, 2, 3, 5}
	testMap := map[string]interface{}{
		"k1": "v1",
		"k2": 2,
		"k3": 3.3,
	}
	DelFrom(&testMap, "k2") // map[string]interface{}{"k1": "v1", "k3": 3.3}
}

func ExampleSlice() {
	testStr := "123456789"
	// 效果等同于Python的  testStr[1:-2][::-2]
	result := Slice(testStr, 1, -2, -2).(string)
	println(result) // [7, 5, 3]
	testArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 效果等同于Python的 testArray[:8][::3]
	resultArray := Slice(testArray, 0, 8, 3).([]int)
	println(resultArray) // [1, 4, 7]
}

func ExampleIfElse() {
	a := map[string]interface{}{
		"name": "zhangsan",
		"id":   1,
	}
	b := map[string]interface{}{
		"name": "wangwu",
		"id":   2,
	}
	// c = a[id]>b[id]?a:b
	c := IfElse(a["id"].(int) > b["id"].(int), a, b).(map[string]interface{})
	fmt.Printf("%#v", c)
}
