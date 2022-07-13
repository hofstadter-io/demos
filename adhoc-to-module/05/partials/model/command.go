var {{camel .Name}}Command = &cobra.Command{
  Use:   "{{ camel .Name }}",
	Short: "manage {{ camel .Name }} resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing required first arg 'verb'")
		}
		verb := args[0]

		switch strings.ToLower(verb) {
			case "create":
				fmt.Println("creating a {{ .Name }} ... tbd")
			case "list":
				fmt.Println("listing {{ .Name }} ... tbd")
			case "get":
				fmt.Println("getting a {{ .Name }} ... tbd")
			case "update":
				fmt.Println("updating a {{ .Name }} ... tbd")
			case "delete":
				fmt.Println("deleting a {{ .Name }} ... tbd")

			default:
			if len(args) < 1 {
				return fmt.Errorf("unknown first arg 'verb' as '%s'", verb)
			}
		}

		fmt.Println("{{ .Name }}", args)

		return nil
		
	},
}
