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
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

type userAccount struct {
	Username string
	Password string
	Email    string
	Phone    string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list user infomation",
	Long:  `list the first user infomation stored in user.json`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open("user.json")
		if err != nil {
			fmt.Println("Open file failed")
			return
		}
		var user interface{}
		decoder := json.NewDecoder(f)
		for {

			err = decoder.Decode(&user)
			if err != nil {
				fmt.Println(err.Error())
				break
			} else {
				fmt.Println(user)
			}
		}

		f.Close()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
