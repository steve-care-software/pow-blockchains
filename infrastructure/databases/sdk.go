package databases

import (
	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/pow-blockchains/domain/genesis"
)

// NewGenesisServiceBuilder creates a new genesis service builder
func NewGenesisServiceBuilder(
	database database_application.Application,
	repositoryBuilder genesis.RepositoryBuilder,
) genesis.ServiceBuilder {
	return createGenesisServiceBuilder(database, repositoryBuilder)
}

// NewGenesisRepositoryBuilder creates a new genesis repository builder
func NewGenesisRepositoryBuilder(
	database database_application.Application,
) genesis.RepositoryBuilder {
	builder := genesis.NewBuilder()
	return createGenesisRepositoryBuilder(database, builder)
}
