package cmd

import (
	"fmt"
	"github.com/nii236/appreciator/models"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generates the letter template",
	Long: `Gen creates a letter to a reciepient and accepts flags including the
names and gifts received, and whether or not they attended the event you hosted.

The event is currently hardcoded to a wedding, but will later be changed to
accept an input, so users can enter their own, such as a birthday or party.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
		names, err := cmd.Flags().GetStringSlice("names")
		from, err := cmd.Flags().GetString("from")
		fmt.Println("Names:", names)
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

		t := template.Must(template.New("letter").Parse(messageTemplate))
		t.Execute(os.Stdout, mValues)

	},
}

func init() {
	RootCmd.AddCommand(genCmd)
	genCmd.Flags().StringSliceP("names", "n", []string{}, "Enter a list of names here")
	genCmd.Flags().StringP("from", "f", "", "Enter your name")
}

const messageTemplate = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
{{.AdditionalNote}}

Best wishes,
{{.From}}
`
