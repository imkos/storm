package msgpack

import (
	"testing"

	"github.com/imkos/storm/v3/codec/internal"
)

func TestMsgpack(t *testing.T) {
	internal.RoundtripTester(t, Codec)
}
