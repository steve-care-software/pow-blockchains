package chains

import (
	"github.com/steve-care-software/blockchains/domain/blocks"
)

// Application represents the chain's application
type Application interface {
	CalculateDifficulty(block blocks.Block) (uint, error)
}
