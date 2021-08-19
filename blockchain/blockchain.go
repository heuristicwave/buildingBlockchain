package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain

// Singleton : 13L의 변수를 외부와 공유하지 않고 아래 함수를 통해 노출
func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}
