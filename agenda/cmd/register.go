/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Used to register an account",
	Long:  `User can register an account by enter username,password,email and phone`,
	Run: func(cmd *cobra.Command, args []string) {
		uname, _ := cmd.Flags().GetString("username")
		pwd, _ := cmd.Flags().GetString("password")
		mail, _ := cmd.Flags().GetString("email")
		phone_num, _ := cmd.Flags().GetString("phone")

		var flag = true
		if uname == "" || pwd == "" || mail == "" || phone_num == "" {
			fmt.Println("Missing necessary user infomation")
			flag = false
		}

		if flag {
			group := userAccount{
				Username: uname,
				Password: pwd,
				Email:    mail,
				Phone:    phone_num,
			}

			var filename = "user.json"
			var f *os.File
			var err error

			if checkFileIsExist(filename) {
				f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)

			} else {
				f, err = os.Create(filename)

			}

			encoder := json.NewEncoder(f)
			err = encoder.Encode(group)

			if err != nil {
				fmt.Println("Register failed", err.Error())
			} else {
				fmt.Println("Register succeed")
			}
			f.Close()
		}

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("username", "u", "", "/")
	registerCmd.Flags().StringP("password", "p", "", "/")
	registerCmd.Flags().StringP("email", "m", "", "/")
	registerCmd.Flags().StringP("phone", "t", "", "/")
}
