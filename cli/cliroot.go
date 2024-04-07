/*
Copyright © 2023 Michele Mendel <michelemendel@blackflatcap.com>
*/
package cli

import (
	"log/slog"
	"os"

	"github.com/michelemendel/genvaeg/repository"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "genvaeg cli",
}

var repo *repository.Repo

func Execute(dbRepo *repository.Repo) {
	repo = dbRepo
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("problems starting the CLI", "err", err)
		os.Exit(1)
	}
}
