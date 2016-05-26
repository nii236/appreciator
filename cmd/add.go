package cmd

import (
	"fmt"
	"github.com/nii236/appreciator/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a single message to the database",
	Long:  `Adds a single message to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		mValues := models.Message{
			NamesList:      []string{"John"},
			Attended:       true,
			Gift:           "Cash",
			AdditionalNote: "Enjoy your wedding next year!",
		}

		fmt.Println(mValues)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

}
