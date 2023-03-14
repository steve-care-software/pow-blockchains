package databases

import (
	"errors"

	database_application "github.com/steve-care-software/databases/applications"
	"github.com/steve-care-software/pow-blockchains/domain/genesis"
)

type genesisServiceBuilder struct {
	database          database_application.Application
	repositoryBuilder genesis.RepositoryBuilder
	pContext          *uint
	pKind             *uint
}

func createGenesisServiceBuilder(
	database database_application.Application,
	repositoryBuilder genesis.RepositoryBuilder,
) genesis.ServiceBuilder {
	out := genesisServiceBuilder{
		database:          database,
		repositoryBuilder: repositoryBuilder,
		pContext:          nil,
		pKind:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *genesisServiceBuilder) Create() genesis.ServiceBuilder {
	return createGenesisServiceBuilder(app.database, app.repositoryBuilder)
}

// WithContext adds a context to the builder
func (app *genesisServiceBuilder) WithContext(context uint) genesis.ServiceBuilder {
	app.pContext = &context
	return app
}

// WithKind adds a kind to the builder
func (app *genesisServiceBuilder) WithKind(kind uint) genesis.ServiceBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Service instance
func (app *genesisServiceBuilder) Now() (genesis.Service, error) {
	if app.pContext == nil {
		return nil, errors.New("the context is mandatory in order to build a genesis Service instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a genesis Service instance")
	}

	repository, err := app.repositoryBuilder.Create().WithContext(*app.pContext).WithKind(*app.pKind).Now()
	if err != nil {
		return nil, err
	}

	return createGenesisService(
		app.database,
		repository,
		*app.pContext,
		*app.pKind,
	), nil
}
