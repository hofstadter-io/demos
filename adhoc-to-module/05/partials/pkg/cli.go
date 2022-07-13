{{ if .Config.Cli.enabled }}
func RunCLI() {
	if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
	rootCmd.AddCommand(serverCommand)
	{{ range .Datamodel.Models -}}
	rootCmd.AddCommand({{ camel .Name }}Command)
	{{ end }}
}

var rootCmd = &cobra.Command{
  Use:   "{{ .Name }}",
  Short: "{{ .Config.About }}",
  Long: `{{ .Config.Help }}`,
}

var serverCommand = &cobra.Command{
	Use: "server",
	Short: "runs the api server",
	Long: "Runs the REST server genrated as part of this binary",
  Run: func(cmd *cobra.Command, args []string) {
		if err := runServer(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
}

const rootLong = `
	{{ .Config.Help }}
`
{{ end }}
