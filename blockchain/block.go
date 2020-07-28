package blockchain

type Block struct {
	timestamp string
	lastHash  string
	hash      string
	data      string
}

func NewBlock(timestamp, lastHash, hash, data string) *Block {

	block := &Block{
		timestamp: timestamp,
		lastHash:  lastHash,
		hash:      hash,
		data:      data,
	}

	return block
}
