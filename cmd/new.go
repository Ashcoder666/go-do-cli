/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "New is used create new todo",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
		type NewToDo struct {
			Title   string
			Status  bool
			DueDate string
			uuid    uuid.UUID
		}

		currentUser, err := user.Current()
		if err != nil {
			fmt.Println("Error getting current user:", err)
			return
		}

		var taskTitle string
		fmt.Println("Enter Task Title")
		fmt.Scan(&taskTitle)
		var taskDate string
		fmt.Println("Enter Task Due Date")
		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')

		taskDate = strings.TrimSpace(input)

		// check if that file exists false then create this else get the value

		filePath := fmt.Sprintf("/home/%v/go/token.txt", currentUser.Username)

		file, err := os.Open(filePath)

		if err == nil {
			content, err := io.ReadAll(file)

			if err != nil {
				fmt.Println("Error reading the file:", err)
				return
			}
			fmt.Println(string(content))
			newtodo := NewToDo{Title: taskTitle, Status: false, DueDate: taskDate, uuid: uuid.UUID(content)}
			fmt.Println(newtodo)
		} else {
			if err != nil {
				fmt.Println("Error creating the file:", err)
				return
			}
			newuuid, _ := uuid.NewUUID()
			// store that in filepath

			file, err := os.Create(filePath)
			content, err := io.WriteString(file, "kkk")

			if err != nil {
				fmt.Println("Error reading the file:", err)
				return
			}

			fmt.Println(content)

			newtodo := NewToDo{Title: taskTitle, Status: false, DueDate: taskDate, uuid: newuuid}
			fmt.Println(newtodo)

		}

		defer file.Close()

		// fmt.Println(newtodo)

		// generate a new uild

		// store in a file and save in go folder // if itsnt exxisnt already
		// using that create and get list

	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
