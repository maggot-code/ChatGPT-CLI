/*
 * @FilePath: /ChatGPT-CLI/cmd/key.go
 * @Author: maggot-code
 * @Date: 2023-02-28 20:34:44
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-02-28 22:12:48
 * @Description:
 */
/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"maggot-code/ChatGPT-CLI/services"

	"github.com/spf13/cobra"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "OpenAI API Key",
	Run: func(cmd *cobra.Command, args []string) {
		services := services.DefineKeyService()
		if cmd.Flags().Changed("set") {
			if len(args) == 0 {
				cmd.Println("Please enter your OpenAI API Key")
				cmd.Help()
				return
			}

			services.SetKey(args[0])
			return
		}
		if cmd.Flags().Changed("get") {
			key, err := services.GetKey()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(key)
			return
		}
		if cmd.Flags().Changed("clear") {
			services.ClearKey()
			return
		}

		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	keyCmd.Flags().BoolP("set", "s", false, "Set OpenAI API Key")
	keyCmd.Flags().BoolP("get", "g", false, "Get OpenAI API Key")
	keyCmd.Flags().BoolP("clear", "c", false, "Clear OpenAI API Key")
}
