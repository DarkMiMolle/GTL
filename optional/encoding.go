package optional

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

func (opt Value[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(opt.value)
}
func (opt *Value[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &opt.value)
}

func (opt Value[T]) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if opt.value == nil {
		return bson.TypeNull, nil, nil
	}
	return bson.MarshalValue(*opt.value)
}

func (opt *Value[T]) UnmarshalBSONValue(t bsontype.Type, data []byte) (err error) {
	if t == bson.TypeNull {
		opt.value = nil
		return nil
	}
	defer func() {
		if err != nil {
			opt.value = nil
		}
	}()
	opt.value = new(T)
	if t == bson.TypeEmbeddedDocument {
		return bson.Unmarshal(data, &opt.value)
	}
	return bson.UnmarshalValue(t, data, &opt.value)
}
