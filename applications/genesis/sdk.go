package genesis

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/pow-blockchains/domain/genesis"
)

// Application represents the genesis application
type Application interface {
	Retrieve(hash hash.Hash) (genesis.Genesis, error)
	Insert(genesis genesis.Genesis) error
	Delete(hash hash.Hash) error
}
