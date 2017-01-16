package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"encoding/hex"
)

type Id interface {
	SetId(id string)
	GetId() string
}

type Handle func(result interface{}) error

func hexifyId(handler Handle, result Id) error {
	if err := handler(&result); err != nil {
		return err
	}
	var bytes = []byte(result.GetId())
	objectId := bson.ObjectIdHex(hex.EncodeToString(bytes))
	result.SetId(objectId.Hex())
	return nil
}