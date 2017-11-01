// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/jack-cheng/CLI-agenda/entity"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "for guest to login",
	Long:  `for guests to enter correct username and password to login `,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if entity.Login(args[0], args[1]) {
			fmt.Println("user:" + args[0] + " log in successfully")
		} else {
			fmt.Println("failed to log in!")
		}
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout",
	Long:  `log out as a guest`,
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if entity.GetCurrentUser().Logout() {
			fmt.Println("now you are a guest")
		}
	},
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "to register a new user",
	Long: `register a new user with this command ,followed by a unique username and
	password. If the user by the given username is already existed, the register
	action will fail.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if entity.Register(args[0], args[1]) {
			fmt.Println("user:" + args[0] + " created successfully!")
		} else {
			fmt.Println("fail to create a new user")
		}
	},
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "list all users",
	Long:  `list all users' information`,
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		entity.GetCurrentUser().LookupAllUser()
	},
}

var cancelUserCmd = &cobra.Command{
	Use:   "cancelUser",
	Short: "remove an account from users",
	Long:  `remove an account from users`,
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		entity.GetCurrentUser().CancelAccount()
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "n", "", "the name of user to log in")
	loginCmd.Flags().StringP("password", "p", "", "the password of user to log in")

	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "the name of new user to be created")
	registerCmd.Flags().StringP("password", "p", "", "the password of user to be created")

	RootCmd.AddCommand(logoutCmd)
	RootCmd.AddCommand(usersCmd)
	RootCmd.AddCommand(cancelUserCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCommandsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCommandsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
