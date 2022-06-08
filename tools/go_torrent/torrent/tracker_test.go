package torrent

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTracker(t *testing.T) {
	file, err := os.Open("../testfile/debian-iso.torrent")
	assert.Equal(t, err, nil)
	tf, err := ParseFile(bufio.NewReader(file))
	assert.Equal(t, err, nil)

	peers := FindPeers(tf)
	for idx, peer := range peers {
		fmt.Printf("Peer %d,Ip: %s,Port:%d \n", idx, peer.Ip, peer.Port)
	}
}
