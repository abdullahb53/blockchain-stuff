package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/abdullahb53/blockchain-stuff/projectx/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     9893394,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}

	h.EncodeBinary(buf)
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     9893394,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecoce := &Block{}
	assert.Nil(t, b.DecodeBinary(buf))
	assert.Equal(t, b, bDecoce)
	// fmt.Printf("%+v", bDecoce)
}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
		},
		Transactions: []Transaction{},
	}
	h := b.Hash()
	fmt.Println(h)
	assert.False(t, h.IsZero())
}
