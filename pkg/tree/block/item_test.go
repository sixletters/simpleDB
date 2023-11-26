package block

import (
	"sixletters/simple-db/pkg/consts"
	"testing"
)

func Test_item_creation(t *testing.T) {
	tests := []struct {
		key   string
		value string
	}{
		{
			key:   "key-1",
			value: "value-2",
		},
		{
			key:   "",
			value: "value-2",
		},
	}

	// Tests for newitem, setkey and setvalue
	for _, tt := range tests {
		item := newItem()
		item.setKey(tt.key)
		item.setValue(tt.value)
		if item.key != tt.key || int(item.keylen) != len(tt.key) {
			t.Errorf("key is not correct in item, expected: %s, actual: %s", tt.key, item.key)
		}
		if item.value != tt.value || int(item.valueLen) != len(tt.value) {
			t.Errorf("key is not correct in item, expected: %s, actual: %s", tt.key, item.key)
		}
	}
	// tests for NewItemWithKv
	for _, tt := range tests {
		item := NewItemWithKV(tt.key, tt.value)
		if item.key != tt.key || int(item.keylen) != len(tt.key) {
			t.Errorf("key is not correct in item, expected: %s, actual: %s", tt.key, item.key)
		}
		if item.value != tt.value || int(item.valueLen) != len(tt.value) {
			t.Errorf("key is not correct in item, expected: %s, actual: %s", tt.key, item.key)
		}
	}
}

func Test_validate(t *testing.T) {
	tests := []struct {
		key           string
		value         string
		errorReturned bool
	}{
		{
			key:           "key-1",
			value:         "value-2",
			errorReturned: false,
		},
		{
			key:           "",
			value:         "value-2",
			errorReturned: false,
		},
		{
			key:           "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			value:         "value-2",
			errorReturned: true,
		},
	}

	for _, tt := range tests {
		item := NewItemWithKV(tt.key, tt.value)
		err := item.validate()
		if (tt.errorReturned && err == nil) || (!tt.errorReturned && err != nil) {
			t.Errorf("Validation of item is not correct")
		}
	}
}

func Test_byte_conversions(t *testing.T) {
	tests := []struct {
		key           string
		value         string
		errorReturned bool
	}{
		{
			key:           "key-1",
			value:         "value-2",
			errorReturned: false,
		},
		{
			key:           "",
			value:         "value-2",
			errorReturned: false,
		},
	}

	for _, tt := range tests {
		item := NewItemWithKV(tt.key, tt.value)
		bytebuffer := item.IntoBytes()
		if len(bytebuffer) != consts.ItemSize {
			t.Errorf("length of byte buffer is note correct expected: %d, actual %d", consts.ItemSize, len(bytebuffer))
		}

		item_from_bytes := newItem().FromBytes(bytebuffer)
		if item_from_bytes.key != item.key || item_from_bytes.keylen != item.keylen {
			t.Errorf("error in conversion of key")
		}
	}
}
