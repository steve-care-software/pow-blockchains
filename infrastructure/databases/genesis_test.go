package databases

import (
	"os"
	"testing"

	"github.com/steve-care-software/databases/infrastructure/files"
	"github.com/steve-care-software/pow-blockchains/domain/genesis"
)

const genesisKindForTests = 0

func TestGenesis_repositoryAndService_Success(t *testing.T) {
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	readChunkSize := uint(1000000)
	defer func() {
		os.RemoveAll(dirPath)
	}()

	fileName := "my_file.db"
	dbApp := files.NewApplication(dirPath, dstExtension, bckExtension, readChunkSize)
	err := dbApp.New(fileName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pContext, err := dbApp.Open(fileName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	defer dbApp.Close(*pContext)

	// create repository:genesisKindForTests
	repositoryBuilder := NewGenesisRepositoryBuilder(dbApp)
	repository, err := repositoryBuilder.Create().WithContext(*pContext).WithKind(genesisKindForTests).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// create service:
	service, err := NewGenesisServiceBuilder(dbApp, repositoryBuilder).Create().WithContext(*pContext).WithKind(genesisKindForTests).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// create instance:
	genIns := genesis.NewGenesisForTests()

	// insert:
	err = service.Insert(genIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = dbApp.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve:
	retGenesis, err := repository.Retrieve(genIns.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare hash:
	if !genIns.Hash().Compare(retGenesis.Hash()) {
		t.Errorf("the returned genesis instance is invalid")
		return
	}

	// inserts again, returns an error:
	err = service.Insert(genIns)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
