package comm

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type ImageHash string

type NullableImageHash struct {
	hash *ImageHash
}

func (n NullableImageHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.hash)
}

func (n NullableImageHash) UnmarshalJSON(bytes []byte) error {
	var tmp string
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	n.hash = (*ImageHash)(&tmp)

	return nil
}

func (n NullableImageHash) IsNull() bool {
	return n.hash == nil
}

func (n NullableImageHash) IsNotNull() bool {
	return n.hash != nil
}

func (n *NullableImageHash) SetNull() discord.NullableField {
	n.hash = nil
	return n
}

func (n NullableImageHash) Get() ImageHash {
	return *n.hash
}

func (n *NullableImageHash) Set(hash ImageHash) {
	n.hash = &hash
}

