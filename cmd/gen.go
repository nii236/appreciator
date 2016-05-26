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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
