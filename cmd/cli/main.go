package main

import (
	"github.com/michelemendel/genvaeg/cli"
	"github.com/michelemendel/genvaeg/repository"
)

func main() {
	repo := repository.NewRepo()
	cli.Execute(repo)
}
