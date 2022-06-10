package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var listFolderSize = &cobra.Command{
	Use:   "listdirsize",
	Short: "returns public ip address",
	Run: func(cmd *cobra.Command, args []string) {
		var size int64

		path := "/home/janbodnar/Documents/prog/golang/"

		err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !info.IsDir() {

				size += info.Size()
			}

			return err
		})

		if err != nil {

			log.Println(err)
		}

		fmt.Printf("The directory size is: %d\n", size/10000)
	},
}

func init() {
	rootCmd.AddCommand(listFolderSize)
}
