package cmd

import "github.com/spf13/cobra"

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates the letter template",
	Long: `Gen creates a letter to a recipient and accepts flags including the
names and gifts received, and whether or not they attended the event you hosted.

The event is currently hardcoded to a wedding, but will later be changed to
accept an input, so users can enter their own, such as a birthday or party.`,
	Run: func(cmd *cobra.Command, args []string) {

		// t := template.Must(template.New("letter").Parse(messageTemplate))
		// t.Execute(os.Stdout, mValues)

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
