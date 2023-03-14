package applications

import (
	"github.com/steve-care-software/pow-blockchains/applications/chains"
	"github.com/steve-care-software/pow-blockchains/applications/genesis"
	"github.com/steve-care-software/pow-blockchains/applications/miners"
)

// Application represents the pow blockchain application
type Application interface {
	Miner() miners.Application
	Genesis() genesis.Application
	Chain() chains.Application
}
