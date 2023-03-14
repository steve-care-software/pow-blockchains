package databases

import (
	"encoding/json"

	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/pow-blockchains/domain/genesis"
	"github.com/steve-care-software/pow-blockchains/infrastructure/objects"
)

type genesisRepository struct {
	database database_application.Application
	builder  genesis.Builder
	context  uint
	kind     uint
}

func createGenesisRepository(
	database database_application.Application,
	builder genesis.Builder,
	context uint,
	kind uint,
) genesis.Repository {
	out := genesisRepository{
		database: database,
		builder:  builder,
		context:  context,
		kind:     kind,
	}

	return &out
}

// Retrieve retrieves a genesis by hash
func (app *genesisRepository) Retrieve(hash hash.Hash) (genesis.Genesis, error) {
	js, err := app.database.ReadByHash(app.context, app.kind, hash)
	if err != nil {
		return nil, err
	}

	ins := new(objects.Genesis)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithDifficulty(ins.Difficulty).
		WithReward(ins.Reward).
		WithHalving(ins.Halving).
		WithMiningValue(ins.MiningValue).
		WithMaxTrxPerBlock(ins.MaxTrxPerBlock).
		WithBlockDuration(ins.BlockDuration).
		CreatedOn(ins.CreatedOn).
		Now()
}
