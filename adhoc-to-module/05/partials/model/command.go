{{ $ModelName := camelT .Name }}
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
				if len(args) < 2 {
					return fmt.Errorf("missing args for create")
				}
				dargs := args[1:]
				dmap := make(map[string]any)
				for _, darg := range dargs {
					parts := strings.Split(darg, "=")
					dmap[parts[0]] = parts[1]
				}
				
				results, err := {{ $ModelName }}Create(dmap)
				if err != nil {
					return err
				}
				if results == nil {
					fmt.Println("{}")
					return nil
				}
				data, err := json.Marshal(results)
				if err != nil {
					return err
				}
				fmt.Println(string(data))
				return nil

			case "update":
				if len(args) < 3 {
					return fmt.Errorf("missing id arg and anything to update")
				}
				dargs := args[2:]
				dmap := make(map[string]any)
				for _, darg := range dargs {
					parts := strings.Split(darg, "=")
					dmap[parts[0]] = parts[1]
				}
				
				results, err := {{ $ModelName }}Update(args[1], dmap)
				if err != nil {
					return err
				}
				if results == nil {
					fmt.Println("{}")
					return nil
				}
				data, err := json.Marshal(results)
				if err != nil {
					return err
				}
				fmt.Println(string(data))
				return nil

			case "list":
				results, err := {{ $ModelName }}List()
				if err != nil {
					return err
				}
				if results == nil {
					fmt.Println("[]")
					return nil
				}
				data, err := json.Marshal(results)
				if err != nil {
					return err
				}
				fmt.Println(string(data))
				return nil

			case "get":
				if len(args) < 2 {
					return fmt.Errorf("missing id arg")
				}
				results, err := {{ $ModelName }}GetByID(args[1])
				if err != nil {
					return err
				}
				if results == nil {
					fmt.Println("{}")
					return nil
				}
				data, err := json.Marshal(results)
				if err != nil {
					return err
				}
				fmt.Println(string(data))
				return nil

			case "delete":
				if len(args) < 2 {
					return fmt.Errorf("missing id arg")
				}
				err := {{ $ModelName }}Delete(args[1])
				if err != nil {
					return err
				}
				fmt.Println("done")
				return nil


			default:
			if len(args) < 1 {
				return fmt.Errorf("unknown first arg 'verb' as '%s'", verb)
			}
		}

		fmt.Println("{{ .Name }}", args)

		return nil
		
	},
}
