package cli

import (
	"fmt"

	"github.com/michelemendel/genvaeg/entity"
	"github.com/michelemendel/genvaeg/util"
	"github.com/spf13/cobra"
)

var email string
var name string
var pw string

func init() {
	rootCmd.AddCommand(readUserCmd)

	rootCmd.AddCommand(createUserCmd)
	createUserCmd.Flags().StringVarP(&name, "name", "n", "", "name")
	createUserCmd.Flags().StringVarP(&pw, "pw", "p", "", "password")

	rootCmd.AddCommand(users)
}

var readUserCmd = &cobra.Command{
	Use:   "user <uid>",
	Short: "DB: Read user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("name required")
			return
		}
		name := args[0]
		u, err := repo.GetUserByName(name)
		if err != nil {
			fmt.Printf("error reading user %s. %s\n", name, err)
			return
		}
		util.PP(u)
	},
}

var createUserCmd = &cobra.Command{
	Use:   "usercreate",
	Short: "DB: Create user",
	Run: func(cmd *cobra.Command, args []string) {
		hashedPw, _ := util.HashPassword(pw)
		u := entity.NewUser(name, hashedPw)
		err := repo.CreateUser(u)
		if err != nil {
			fmt.Println("user create err", err)
		}
		fmt.Println("created user", u.UUID)
	},
}

var users = &cobra.Command{
	Use:   "users",
	Short: "DB: List all users",
	Run: func(cmd *cobra.Command, args []string) {
		users := repo.GetAllUsers()
		for _, u := range users {
			fmt.Printf("%s : %s\n", u.UUID, u.Name)
		}
	},
}
