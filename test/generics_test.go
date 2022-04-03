package test

import (
    "fmt"
    "strconv"
    "testing"
)

/*
   golang 1.18 泛型测试
   自带泛型：
   1.any 表示任何类型，即interface
   2.comparable 表示可以被比较的类型
       comparable is an interface that is implemented by all comparable types
       (booleans, numbers, strings, pointers, channels, arrays of comparable types,
       structs whose fields are all comparable types).
       The comparable interface may only be used as a type parameter constraint,
       not as the type of a variable.
*/

// 1. 泛型的类型限制，在函数上直接申明该函数支持的多个类型
func AddElem[T int | string](params []T) (sum T) {
    for _, elem := range params {
        sum += elem
    }
    return
}

func TestGenerics_AddElem(t *testing.T) {

    // 1. 在函数上声明泛型支持的多个类型
    // 1.1 传入支持的int
    intSum := AddElem([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
    t.Logf("测试1.1: 类型=%T，val=%+v", intSum, intSum)

    // 1.2 传入支持的string
    strSum := AddElem([]string{"静", "以", "修", "身", "，", "俭", "以", "养", "德"})
    t.Logf("测试1.2: 类型=%T，val=%+v", strSum, strSum)

    // 1.3 传入不支持的类型  报错./generics_test.go:29:24: float64 does not implement int|string
    //floatSum := AddElem([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
    //t.Logf("测试1.3: 类型=%T，val=%+v", floatSum, floatSum)
}

// 2 泛型的类型限制，声明一个interface，包括所有需要支持的类型
// ~int 表示底层数据是int
type NumStr interface {
    ~int | ~uint | ~float64 | ~string
}

type MyInt int

func AddNumStr[T NumStr](params []T) (sum T) {
    for _, param := range params {
        sum += param
    }
    return
}

func TestGenerics_AddNumStr(t *testing.T) {
    // 2.1 支持的int
    intSum := AddNumStr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
    t.Logf("测试2.1: 类型=%T，val=%+v", intSum, intSum)

    // 2.2 传入支持的string
    strSum := AddNumStr([]string{"风", "平", "浪", "静"})
    t.Logf("测试2.2: 类型=%T，val=%+v", strSum, strSum)

    // 2.3 传入自定义int
    myIntSum := AddNumStr([]MyInt{1, 2, 3, 4, 5, 6, 7, 8, 9})
    t.Logf("测试2.3: 类型=%T，val=%+v", myIntSum, myIntSum)
}

// 3. 泛型切片
// any为泛型自带的一种类型，即interface
type Vector[T any] []T
type NumSlice[T int | float64] []T

func TestGenerics_Slice(t *testing.T) {
    // 测试3.1
    v := Vector[string]{"z", "x", "c"}
    t.Logf("测试3.1: 类型=%T，val=%+v", v, v)

    // 测试3.2
    ns := NumSlice[int]{1, 2, 3, 4, 5, 6}
    t.Logf("测试3.2: 类型=%T，val=%+v", ns, ns)

    // 测试3.3
    sum := AddElem(ns)
    t.Logf("测试3.3: 类型=%T，val=%+v", sum, sum)
}

// 4. 泛型map
type M[K string, V any] map[K]V

func TestGenerics_Map(t *testing.T) {
    m := M[string, int]{
        "zx": 123,
        "as": 456,
        "qw": 789,
    }
    t.Logf("测试4.1: 类型=%T，val=%+v", m, m)
}

// 5. 泛型通道
type Ch[T any] chan T

func TestGenerics_Chan(t *testing.T) {
    ch := make(Ch[int], 1)
    ch <- 10

    res := <-ch
    t.Logf("测试5.1: 类型=%T，val=%+v", res, res)
    t.Logf("测试5.2: 类型=%T，val=%+v", ch, ch)
}

// 6. 方法约束
type FlyAnimal interface {
    ToString() string
}

// Dragon实现了FlyAnimal
type Dragon int

func (d Dragon) ToString() string {
    return "string_" + strconv.Itoa(int(d))
}

// Tiger没有实现flyAnimal
type Tiger int

func PrintStr[T FlyAnimal](params ...T) {
    for _, param := range params {
        fmt.Println(param.ToString())
    }
}

func TestGenerics_Method_limit(t *testing.T) {
    // 测试6.1 传入实现了方法的类型
    dragon := Dragon(1)
    PrintStr(dragon)

    // 测试6.2 传入未实现对应方法的类型  ./generics_test.go:136:13: Tiger does not implement FlyAnimal (missing ToString method)
    //tiger := Tiger(100)
    //PrintStr(tiger)
}

// 7. 类型+方法的约束
type CanSpeak interface {
    ~int | ~int32 | ~int64 | ~float32 | ~float64
    Speak() string
}

type Mouth int32

func (m Mouth) Speak() string {
    return fmt.Sprintf("speak %v", m)
}

type Nose string

func (n Nose) Speak() string {
    return fmt.Sprintf("speak %v", n)
}

type Ear int

func SpeakLoudly[T CanSpeak](params []T) {
    for _, param := range params {
        fmt.Println(param.Speak())
    }
}

func TestGenerics_Type_Method_Limit(t *testing.T) {
    // 7.1 测试类型与方法均符合
    SpeakLoudly([]Mouth{1, 2, 3, 4, 5, 6})

    // 7.2 测试类型符合   ./generics_test.go:172:16: Nose does not implement CanSpeak
    //SpeakLoudly([]Nose{"z", "x", "c"})

    // 7.3 测试方法符合   ./generics_test.go:175:16: Ear does not implement CanSpeak (missing Speak method)
    //SpeakLoudly([]Ear{1, 2, 3, 4, 5, 6})
}
