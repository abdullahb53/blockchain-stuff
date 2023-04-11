package core

import (
	"crypto/sha256"

	"github.com/abdullahb53/blockchain-stuff/projectx/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(b *Block) types.Hash {

	h := sha256.Sum256(b.HeaderData())
	return types.Hash(h)
}

// ABSTRACTION IS HERE. U SHOULD NOT DELETE such a blockhasher types.
// So that types give the abstraction to implement other Hash Types.
// type xyzHasher struct {
// }

// func (xyzHasher) Hash(b *Block) types.Hash {
// 	// DO OTHER STUFFS.
// 	h := sha256.Sum256(b.HeaderData())
// 	return types.Hash(h)
// }
