package chains

import (
	blockchains_application "github.com/steve-care-software/blockchains/applications"
	"github.com/steve-care-software/blockchains/domain/blocks"
)

// Application represents the chain's application
type Application interface {
	CalculateNextDifficulty(block blocks.Block) (uint, error)
	Blockchain() blockchains_application.Application
}
