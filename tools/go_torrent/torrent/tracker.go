package torrent

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"sharp/tools/go_torrent/bencode"
	"strconv"
	"time"
)

const (
	PeerPort = 6666
	IpLen    = 4
	PortLen  = 2
	PeerLen  = IpLen + PortLen
)

type PeerInfo struct {
	Ip   net.IP
	Port uint16
}

type TrackerResp struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func buildUrl(tf *TorrentFile) (string, error) {
	var peerId [20]byte
	_, err := rand.Read(peerId[:])
	if err != nil {
		return "", err
	}
	base, err := url.Parse(tf.Announce)
	if err != nil {
		fmt.Println("Announce Error: " + tf.Announce)
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(tf.InfoSHA[:])},
		"peer_id":    []string{string(peerId[:])},
		"port":       []string{strconv.Itoa(PeerPort)},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"0"},
		"left":       []string{strconv.Itoa(tf.FileLen)},
	}

	base.RawQuery = params.Encode()
	return base.String(), nil
}

func buildPeerInfo(peers []byte) []PeerInfo {
	num := len(peers) / PeerLen
	if len(peers)%PeerLen != 0 {
		fmt.Println("Received malformed peers")
		return nil
	}
	infos := make([]PeerInfo, num)
	for i := 0; i < num; i++ {
		offset := i * PeerLen
		infos[i].Ip = net.IP(peers[offset : offset+IpLen])
		infos[i].Port = binary.BigEndian.Uint16(peers[offset+IpLen : offset+PortLen])
	}
	return infos
}

func FindPeers(tf *TorrentFile) []PeerInfo {
	url, err := buildUrl(tf)
	if err != nil {
		fmt.Println("Build Tracker Url Error: " + err.Error())
		return nil
	}
	cli := &http.Client{Timeout: 15 * time.Second}
	resp, err := cli.Get(url)
	if err != nil {
		fmt.Println("Fail to connect to Tracker: " + err.Error())
		return nil
	}
	defer resp.Body.Close()

	trackResp := new(TrackerResp)
	err = bencode.Unmarshal(resp.Body, trackResp)
	if err != nil {
		fmt.Println("Tracker Response Error: " + err.Error())
		return nil
	}
	return buildPeerInfo([]byte(trackResp.Peers))
}
