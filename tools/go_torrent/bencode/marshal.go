package bencode

import (
    "errors"
    "io"
    "reflect"
    "strings"
)

func Unmarshal(r io.Reader, s interface{}) error {
    o, err := Parse(r)
    if err != nil {
        return err
    }

    p := reflect.ValueOf(s)
    if p.Kind() != reflect.Ptr {
        return errors.New("dest must be a pointer")
    }

    switch o.type_ {
    case BLIST:
        list, _ := o.List()
        // 外面传进来指针指向的slice可能是一个空的slice或者长度与预期的不一致，所以需要内部make，set一下
        l := reflect.MakeSlice(p.Elem().Type(), len(list), len(list))
        p.Elem().Set(l)
        err = unmarshalList(p, list)
        if err != nil {
            return err
        }
    case BDICT:
        dict, _ := o.Dict()
        // 可以直接set
        err = unmarshalDict(p, dict)
        if err != nil {
            return err
        }
    default:
        return errors.New("src code must be struct or slice")
    }
    return nil
}

func unmarshalList(p reflect.Value, list []*BObject) error {
    // 保证传入的是指针，类型是slice
    if p.Kind() != reflect.Ptr || p.Elem().Type().Kind() != reflect.Slice {
        return errors.New("dest must be pointer to slice")
    }
    // 获取指针里的元素，slice
    v := p.Elem()
    if len(list) == 0 {
        return nil
    }
    // bencode的list只能是同类型的？
    switch list[0].type_ {
    case BSTR:
        for i, o := range list {
            val, err := o.Str()
            if err != nil {
                return err
            }
            v.Index(i).SetString(val)
        }
    case BINT:
        for i, o := range list {
            val, err := o.Int()
            if err != nil {
                return err
            }
            v.Index(i).SetInt(int64(val))
        }
    case BLIST:
        for i, o := range list {
            val, err := o.List()
            if err != nil {
                return err
            }
            if v.Type().Elem().Kind() != reflect.Slice {
                return ErrTyp
            }
            lp := reflect.New(v.Type().Elem())
            // make一个长度和类型与原slice一样的slice
            ls := reflect.MakeSlice(v.Type().Elem(), len(val), len(val))
            lp.Elem().Set(ls)
            err = unmarshalList(lp, val)
            if err != nil {
                return err
            }
            v.Index(i).Set(lp.Elem())
        }
    case BDICT:
        for i, o := range list {
            val, err := o.Dict()
            if err != nil {
                return err
            }
            // Kind 类型的种类
            if v.Type().Elem().Kind() != reflect.Struct {
                return ErrTyp
            }
            dp := reflect.New(v.Type().Elem())
            err = unmarshalDict(dp, val)
            if err != nil {
                return err
            }
            v.Index(i).Set(dp.Elem())
        }
    }
    return nil
}

func unmarshalDict(p reflect.Value, dict map[string]*BObject) error {
    if p.Kind() != reflect.Ptr || p.Elem().Type().Kind() != reflect.Struct {
        return errors.New("dest must be pointer")
    }
    // 获取元素，struct
    v := p.Elem()
    // 遍历所有字段
    for i, n := 0, v.NumField(); i < n; i++ {
        fv := v.Field(i)
        if !fv.CanSet() {
            continue
        }
        ft := v.Type().Field(i)
        key := ft.Tag.Get("bencode")
        if key == "" {
            // 保证代码顺利执行
            key = strings.ToLower(ft.Name)
        }
        // 根据上面获得的key获取数据
        fo := dict[key]
        if fo == nil {
            continue
        }
        switch fo.type_ {
        case BSTR:
            if ft.Type.Kind() != reflect.String {
                break
            }
            val, _ := fo.Str()
            fv.SetString(val)
        case BINT:
            if ft.Type.Kind() != reflect.Int {
                break
            }
            val, _ := fo.Int()
            fv.SetInt(int64(val))
        case BLIST:
            if ft.Type.Kind() != reflect.Slice {
                break
            }
            list, _ := fo.List()
            lp := reflect.New(ft.Type)
            ls := reflect.MakeSlice(ft.Type, len(list), len(list))
            lp.Elem().Set(ls)
            err := unmarshalList(lp, list)
            if err != nil {
                return err
            }
            fv.Set(lp.Elem())
        case BDICT:

            if v.Type().Elem().Kind() != reflect.Struct {
                return ErrTyp
            }
            dp := reflect.New(ft.Type)

            dictK, _ := fo.Dict()
            err := unmarshalDict(dp, dictK)
            if err != nil {
                return err
            }
            fv.Set(dp.Elem())
        }
    }
    return nil
}

func marshalValue(w io.Writer, v reflect.Value) int {
    length := 0
    switch v.Kind() {
    case reflect.String:
        length += EncodeString(w, v.String())
    case reflect.Int:
        length += EncodeInt(w, int(v.Int()))
    case reflect.Slice:
        length += marshalList(w, v)
    case reflect.Struct:
        length += marshalDict(w, v)
    }
    return length
}

func marshalList(w io.Writer, vl reflect.Value) int {
    length := 2
    w.Write([]byte{'l'})
    for i := 0; i < vl.Len(); i++ {
        ev := vl.Index(i)
        length += marshalValue(w, ev)
    }
    w.Write([]byte{'e'})
    return length
}

func marshalDict(w io.Writer, vd reflect.Value) int {
    length := 2
    w.Write([]byte{'d'})
    for i := 0; i < vd.NumField(); i++ {
        fv := vd.Field(i)
        ft := vd.Type().Field(i)
        key := ft.Tag.Get("bencode")
        if key == "" {
            key = strings.ToLower(ft.Name)
        }
        length += EncodeString(w, key)
        length += marshalValue(w, fv)
    }
    w.Write([]byte{'e'})
    return length
}

func Marshal(w io.Writer, s interface{}) int {
    v := reflect.ValueOf(s)
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    return marshalValue(w, v)
}
