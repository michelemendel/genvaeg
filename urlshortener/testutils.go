package urlshortener

import (
	"fmt"
	"os"
	"path"

	"github.com/michelemendel/genvaeg/constants"
	"github.com/michelemendel/genvaeg/repository"
	"github.com/michelemendel/genvaeg/util"
)

func InitTest() *repository.Repo {
	rootDir := util.GetRootDir()
	dbDir := path.Join(rootDir, os.Getenv(constants.ENV_DB_DIR_KEY))

	os.Setenv(constants.ENV_DB_NAME_KEY, "test.db")
	os.Setenv(constants.ENV_DB_DIR_KEY, dbDir)

	fmt.Println("DB test file", path.Join(dbDir, os.Getenv(constants.ENV_DB_NAME_KEY)))

	repo := repository.NewRepo()

	repo.DropTables()
	repo.CreateTables()
	repo.CreateIndexes()

	return repo
}
