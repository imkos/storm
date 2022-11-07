package codec_test

import (
	"fmt"

	"github.com/imkos/storm/v3"
	"github.com/imkos/storm/v3/codec/gob"
	"github.com/imkos/storm/v3/codec/json"
	"github.com/imkos/storm/v3/codec/msgpack"
	"github.com/imkos/storm/v3/codec/protobuf"
	"github.com/imkos/storm/v3/codec/sereal"
)

func Example() {
	// The examples below show how to set up all the codecs shipped with Storm.
	// Proper error handling left out to make it simple.
	gobDb, _ := storm.Open("gob.db", storm.Codec(gob.Codec))
	jsonDb, _ := storm.Open("json.db", storm.Codec(json.Codec))
	msgpackDb, _ := storm.Open("msgpack.db", storm.Codec(msgpack.Codec))
	serealDb, _ := storm.Open("sereal.db", storm.Codec(sereal.Codec))
	protobufDb, _ := storm.Open("protobuf.db", storm.Codec(protobuf.Codec))

	fmt.Printf("%T\n", gobDb.Codec())
	fmt.Printf("%T\n", jsonDb.Codec())
	fmt.Printf("%T\n", msgpackDb.Codec())
	fmt.Printf("%T\n", serealDb.Codec())
	fmt.Printf("%T\n", protobufDb.Codec())

	// Output:
	// *gob.gobCodec
	// *json.jsonCodec
	// *msgpack.msgpackCodec
	// *sereal.serealCodec
	// *protobuf.protobufCodec
}
