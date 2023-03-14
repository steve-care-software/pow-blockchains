package miners

import (
	"math/big"

	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents the miner application
type Application interface {
	Execute(hash hash.Hash, miningValue uint8, difficulty uint) (*big.Int, error)
}
