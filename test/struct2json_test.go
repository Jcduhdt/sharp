package test

import (
    "encoding/json"
    "fmt"
    "testing"
)

type School struct {
    Name       string    `json:"name,omitempty"` // omitempty,当值为该类型的默认值时会被忽略掉，不会出现在json结果中
    Principal  *Teacher  `json:"principal,omitempty"`
    Teachers   []Teacher `json:"teachers,omitempty"`
    Students   []Student `json:"students,omitempty"`
    BuildNum   *int      `json:"build_num,omitempty"`   // 只要指针不是nil，即便值为默认值0也会出现在json结果中，所以当某种情况下需要忽略该参数，但是有不能忽略默认值时使用
    CollegeNum int       `json:"college_num,omitempty"` // 转json的时候会忽略该类型的默认值，int即为0
    Area       int       `json:"-"`                     // 转成json的时候永远忽略这个值，无论是什么类型，在任何情况下使用整个对象的jspon时都不需要该字段是可使用
}

type Student struct {
    Name string `json:"name,omitempty"`
    Age  int    `json:"age,omitempty"`
}

type Teacher struct {
    Name  string `json:"name,omitempty"`
    Age   int    `json:"age,omitempty"`
    Major string `json:"major,omitempty"`
}

func TestStruct2Json(t *testing.T) {
    students := []Student{
        {
            Name: "张三",
            Age:  15,
        },
        {
            Name: "李四",
            Age:  16,
        },
    }

    principal := Teacher{
        Name:  "校长",
        Age:   56,
        Major: "manage",
    }

    teachers := []Teacher{{
        Name: "teacher wang",
    }}

    buildNum, collegeNum, area := 0, 0, 123456

    schoolInfo := School{
        Name:       "星光",
        Principal:  &principal,
        Teachers:   teachers,
        Students:   students,
        BuildNum:   &buildNum,
        CollegeNum: collegeNum,
        Area:       area,
    }

    bytes, _ := json.Marshal(schoolInfo)

    fmt.Println(string(bytes))
    // json结果
    //{"name":"星光","principal":{"name":"校长","age":56,"major":"manage"},"teachers":[{"name":"teacher wang"}],"students":[{"name":"张三","age":15},{"name":"李四","age":16}],"build_num":0}

}
