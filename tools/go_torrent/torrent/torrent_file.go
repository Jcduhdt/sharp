package torrent

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"sharp/tools/go_torrent/bencode"
)

type rawInfo struct {
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	PieceLength int    `bencode:"piece length"`
	pieces      string `bencode:"pieces"`
}

type rawFile struct {
	Announce string  `bencode:"announce"`
	Info     rawInfo `bencode:"info"`
}

const SHALEN = 20

// 将torrentFile打平
type TorrentFile struct {
	Announce  string
	InfoSHA   [SHALEN]byte // 唯一标识
	FileName  string
	FileLen   int
	PieceLen  int
	PiecesSHA [][SHALEN]byte
}

func ParseFile(r io.Reader) (*TorrentFile, error) {
	// new得到的就是一个指针，指向该类型零值得指针
	raw := new(rawFile)
	err := bencode.Unmarshal(r, raw)
	if err != nil {
		fmt.Println("Fail to parse torrent file")
		return nil, err
	}
	ret := new(TorrentFile)
	ret.Announce = raw.Announce
	ret.FileName = raw.Info.Name
	ret.FileLen = raw.Info.Length
	ret.PieceLen = raw.Info.PieceLength

	// calculate info SHA
	buf := new(bytes.Buffer)
	// todo 作者说Marshal没有用tag，使用的是filedName，后面待修改
	wLen := bencode.Marshal(buf, raw.Info)
	if wLen == 0 {
		fmt.Println("raw file info error")
	}
	ret.InfoSHA = sha1.Sum(buf.Bytes())

	// calculate pieces SHA
	bys := []byte(raw.Info.pieces)
	cnt := len(bys) / SHALEN
	hashes := make([][SHALEN]byte, cnt)
	for i := 0; i < cnt; i++ {
		copy(hashes[i][:], bys[i*SHALEN:(i+1)*SHALEN])
	}
	ret.PiecesSHA = hashes
	return ret, nil
}
