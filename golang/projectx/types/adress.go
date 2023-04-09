package types

import (
	"encoding/hex"
	"fmt"
)

type Adress [20]uint8

func (a Adress) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

func (a Adress) String() string {
	return hex.EncodeToString(a.ToSlice())
}

func AdressFromBytes(b []byte) Adress {
	if len(b) != 20 {
		msg := fmt.Sprintf("given bytes with lengths %d should be 20", len(b))
		panic(msg)
	}

	var value [20]uint8
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}

	return Adress(value)

}
