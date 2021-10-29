package blockchain

import (
	"fmt"
	"sync"

	"githun.com/heuristicwave/buildingBlockchain/db"
	"githun.com/heuristicwave/buildingBlockchain/utils"
)

type blockchain struct {
	NewestHash string `json:"newesthash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// Singleton : 변수를 외부와 공유하지 않고 아래 함수를 통해 노출
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			// search for checkpoint on the db
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				// restore b from bytes
				b.restore(checkpoint)
			}

		})
	}
	fmt.Println(b.NewestHash)
	return b
}
