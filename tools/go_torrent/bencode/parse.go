package bencode

import (
	"bufio"
	"io"
)

// Parse 解析缓存读取器里面的数据
func Parse(r io.Reader) (*BObject, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}

	// 返回下一个字节的数据，不推进缓存
	b, err := br.Peek(1)
	if err != nil {
		return nil, err
	}

	var ret BObject
	// 根据buffer中的第一位判断走什么逻辑
	switch {
	case b[0] >= '0' && b[0] <= '9':
		val, errDe := DecodeString(br)
		if errDe != nil {
			return nil, errDe
		}
		ret.type_ = BSTR
		ret.val_ = val
	case b[0] == 'i':
		val, errDe := DecodeInt(br)
		if errDe != nil {
			return nil, errDe
		}
		ret.type_ = BINT
		ret.val_ = val
	case b[0] == 'l':
		// 处理掉第一位的l
		br.ReadByte()
		var list []*BObject
		for {
			if p, _ := br.Peek(1); p[0] == 'e' {
				// 处理掉最后一位e，不可能是子元素的e，因为已经在parse里处理完了
				br.ReadByte()
				break
			}
			elem, errP := Parse(br)
			if errP != nil {
				return nil, errP
			}
			list = append(list, elem)
		}
		ret.type_ = BLIST
		ret.val_ = list
	case b[0] == 'd':
		br.ReadByte()
		dict := make(map[string]*BObject)
		for {
			if p, _ := br.Peek(1); p[0] == 'e' {
				br.ReadByte()
				break
			}
			key, errDe := DecodeString(br)
			if errDe != nil {
				return nil, errDe
			}
			val, errP := Parse(br)
			if errP != nil {
				return nil, errP
			}
			dict[key] = val
		}
		ret.type_ = BDICT
		ret.val_ = dict
	default:
		return nil, ErrIvd
	}
	return &ret, nil
}
