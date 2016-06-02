package cmd

import (
	"fmt"
	"log"

	"github.com/nii236/appreciator/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a single message to the database",
	Long:  `Adds a single message to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		names, err := cmd.Flags().GetStringSlice("names")
		if err != nil {
			fmt.Println(err)
			return
		}
		from, err := cmd.Flags().GetString("from")
		fmt.Println("Names:", names)
		if len(names) == 0 {
			log.Fatal("ERROR: Please enter recipients")
		}

		if from == "" {
			log.Fatal("ERROR: Please enter your name under flag --from")
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		mValues := models.Message{
			NamesList:      names,
			Attended:       true,
			Gift:           "Cash",
			AdditionalNote: "Enjoy your wedding next year!",
			From:           from,
		}

		mValues.Process()

		fmt.Printf("%+v", mValues)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().StringSliceP("names", "n", []string{}, "Enter a list of names here")
	addCmd.Flags().StringP("from", "f", "", "Enter your name")

}
