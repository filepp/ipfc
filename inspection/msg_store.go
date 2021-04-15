package inspection

import (
	"bytes"
	"encoding/gob"
	ds2 "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-ipfs/miner/proto"
)

const (
	windowPostMessagePrefix = "wndpost/"
)

func saveWindowPostMessage(localStore ds2.Datastore, minerId, msgId string, msg proto.WindowPostReq) error {
	return localStore.Put(windowPostKey(minerId, msgId), encodeMessage(msg))
}

func getWindowPostMessage(localStore ds2.Datastore, minerId, msgId string) (proto.WindowPostReq, error) {
	data, err := localStore.Get(windowPostKey(minerId, msgId))
	if err != nil {
		return proto.WindowPostReq{}, err
	}
	return decodeMessage(data)
}

func deleteWindowPostMessage(localStore ds2.Datastore, minerId, msgId string) error {
	return localStore.Delete(windowPostKey(minerId, msgId))
}

func windowPostKey(minerId, msgId string) ds2.Key {
	return ds2.NewKey(windowPostMessagePrefix + minerId + "/" + msgId)
}

func encodeMessage(msg proto.WindowPostReq) []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(msg)
	if err != nil {
		return []byte{}
	}
	return buffer.Bytes()
}

func decodeMessage(data []byte) (proto.WindowPostReq, error) {
	dec := gob.NewDecoder(bytes.NewReader(data))
	var v proto.WindowPostReq
	err := dec.Decode(&v)
	if err != nil {
		return proto.WindowPostReq{}, err
	}
	return v, nil
}
