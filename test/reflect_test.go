package test

import (
    "fmt"
    "reflect"
    "testing"
)

/*
   reflect包实现了运行时反射，允许程序操作任意类型的对象
   通过调用TypeOf获取interface的动态类型信息，返回一个Type类型值
   通过调用ValueOf函数返回一个Value类型值，该值代表运行时数据
   Zero接收一个Type类型参数并返回一个该类型零值的Value类型值
*/

func TestReflect(t *testing.T) {
    var testInt interface{} = 3
    value := reflect.ValueOf(testInt)
    // 应使用Kind函数判断该interface是什么类型，然后调用对应的方法，否者会panic
    if value.Kind() == reflect.Int {
        fmt.Println(value.Int())
    }

    var testSlice interface{} = []int{1, 2, 3}
    sT := reflect.TypeOf(testSlice)

    sT.Len()
    // Elem 返回类型的元素类型
    elem := sT.Elem()
    fmt.Println(elem)
}
