/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"
	bookpb "gRPC_Mongo_HTTP_TestTask/proto"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new bookpost on the server through gRPC",
	Long: `Create a new bookpost on the server through gRPC.

	A book post requires an AuthorId, Name and Content`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// получаем данные с флагов
		author, err := cmd.Flags().GetString("author")
		name, err := cmd.Flags().GetString("name")
		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			return err
		}

		// сообщения для пробафера
		book := &bookpb.Book{
			AuthorId: author,
			Name:     name,
			Tag:      tag,
		}

		//вызываем RPC
		res, err := client.CreateBook(
			context.TODO(),
			&bookpb.CreateBookReq{
				Book: book,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Book created: %s\n", res.Book.Id)
		return nil
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().StringP("author", "a", "", "Add an author")
	createCmd.Flags().StringP("name", "n", "", "A name the book")
	createCmd.Flags().StringP("tag", "t", "", "The tag for the book")
	createCmd.MarkFlagRequired("author")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("tag")
	rootCmd.AddCommand(createCmd)
}
