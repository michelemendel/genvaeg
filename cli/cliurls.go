package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(dropTablesCmd)
	rootCmd.AddCommand(createTablesCmd)
	rootCmd.AddCommand(selectAllURLPairs)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "DB: drop and create tables",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Resetting database...")
		repo.DropTables()
		repo.CreateTables()
	},
}

var dropTablesCmd = &cobra.Command{
	Use:   "droptables",
	Short: "DB: drop tables",
	Run: func(cmd *cobra.Command, args []string) {
		repo.DropTables()
	},
}

var createTablesCmd = &cobra.Command{
	Use:   "createtables",
	Short: "DB: create tables",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating tables...")
		repo.CreateTables()
	},
}

var selectAllURLPairs = &cobra.Command{
	Use:   "urls",
	Short: "DB: list all URL pairs",
	Run: func(cmd *cobra.Command, args []string) {
		urlPairs, err := repo.GetAllURLPairs()
		if err != nil {
			fmt.Println("failed to select URLs: ", err)
			return
		}
		for _, u := range urlPairs {
			fmt.Printf("%s : %s\n", u.ShortURL, u.FullURL)
		}
	},
}
