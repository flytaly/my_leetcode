package serializeanddeserializebinarytree

import (
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	t.Run("serialize-deserialize", func(t *testing.T) {
		ser := Constructor()
		deser := Constructor()
		// data := `1,2,3,null,null,4,5`
		// data := `1,2`
		data := `4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2`
		root := deser.deserialize(data)
		got := ser.serialize(root)

		if data != got {
			t.Errorf("Codec.deserialize() = %v, want %v", got, data)
		}
	})
}
