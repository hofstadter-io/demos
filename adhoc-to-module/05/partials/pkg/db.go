var db *gorm.DB

func initDB() (err error) {
	// Create db connection
	db, err = gorm.Open(sqlite.Open("{{ .Name }}.db"), &gorm.Config{})
  if err != nil {
		return fmt.Errorf("failed to connect database:\n%s", err)
  }

  // Migrate the schema
	{{ range .Datamodel.Models -}}
  err = db.AutoMigrate(&{{ .Name }}{})
  if err != nil {
		return fmt.Errorf("failed to migrate database:\n%s", err)
  }
  {{ end }}

	return nil
}


type seedData struct {
	{{ range .Datamodel.Models -}}
	{{ .PluralName }} []{{ .Name }}
	{{ end }}
}

func seedDB(filename string) error {
	// order is semi-important here
	// we want to do as much work
	// without modifying the database
	// as possible, incase there are errors

	var data seedData

	// read file
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	jsonByte, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	// decode data
	json.Unmarshal([]byte(jsonByte), &data)
	// setup db
	if err := initDB(); err != nil {
		return err
	}

	// seed db
	{{ range .Datamodel.Models -}}
	if len(data.{{ .PluralName }}) == 0 {
		fmt.Println("No {{ .PluralName }} found in", filename)
	} else {
		fmt.Println("Seeding {{ .PluralName }} from", filename)
	}
	for _, d := range data.{{ .PluralName }} {
		D := d // local reference
		// fmt.Printf("adding: %#v\n\n", D)
		res := db.Create(&D)
		if res.Error != nil {
			fmt.Println("while seeding {{ .Name }}:", D)
			return res.Error
		}
		res = db.Save(&D)
		if res.Error != nil {
			fmt.Println("while seeding {{ .Name }}:", D)
			return res.Error
		}
	}
	{{ end }}

	return nil
}

var seedCommand = &cobra.Command{
	Use: "seed users.json data.json...",
	Short: "seeds the database",
	Long: seedLongHelp,
  RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no seed json files provided")
		}
		for _, arg := range args {
			if err := seedDB(arg); err != nil {
				return err
			}
		}

		return nil
  },
}

const seedLongHelp = `
Seeds the database, talking directly to it.

Seed data schema, owned models will be created as well

{{ range .Datamodel.Models -}}
{{ .PluralName }}: [...{
  // fields
  {{ range .Fields -}}
  {{ .Name }}: {{ .Type }}
  {{ end }}
  // relations
  {{ range .Reln -}}
  {{- if eq .Type "OwnedBy" -}}
  {{ .Name }}ID: uint
  {{- end }}
  {{- if eq .Type "HasOne" -}}
  {{ .Name }}: {{ .Type }}
  {{- end }}
  {{- if eq .Type "HasMany" "ManyToMany" -}}
  {{ .PluralName }}: [...{{ .Name }}]
  {{- end }}
  {{- end }}
}]

{{ end }}
`
