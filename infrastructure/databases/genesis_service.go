package databases

import (
	"encoding/json"
	"errors"
	"fmt"

	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/pow-blockchains/domain/genesis"
	"github.com/steve-care-software/pow-blockchains/infrastructure/objects"
)

type genesisService struct {
	database   database_application.Application
	repository genesis.Repository
	context    uint
	kind       uint
}

func createGenesisService(
	database database_application.Application,
	repository genesis.Repository,
	context uint,
	kind uint,
) genesis.Service {
	out := genesisService{
		database:   database,
		repository: repository,
		context:    context,
		kind:       kind,
	}

	return &out
}

// Insert inserts a genesis instance
func (app *genesisService) Insert(genesis genesis.Genesis) error {
	hash := genesis.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		str := fmt.Sprintf("the Genesis (hash: %s) already exists", hash.String())
		return errors.New(str)
	}

	ins := objects.Genesis{
		Difficulty:     genesis.Difficulty(),
		Reward:         genesis.Reward(),
		Halving:        genesis.Halving(),
		MiningValue:    genesis.MiningValue(),
		MaxTrxPerBlock: genesis.MaxTrxPerBlock(),
		BlockDuration:  genesis.BlockDuration(),
		CreatedOn:      genesis.CreatedOn(),
	}

	js, err := json.Marshal(ins)
	if err != nil {
		return err
	}

	return app.database.Write(
		app.context,
		app.kind,
		hash,
		js,
	)
}
