package {{ .Datamodel.Name }}

import (
	"gorm.io/gorm"
)

{{ range .Datamodel.Models }}
{{ $ModelName := camelT .Name }}
// helper functions for {{ $ModelName }}
func create{{ $ModelName }}(in *{{ .Name }}) error {
	res := db.Create(&{{ .Name }}{
		{{ range .Fields }}{{ .Name }}: in.{{ .Name }},
		{{ end }}
	})
	return res.Error
}

func list{{ $ModelName }}() ([]*{{ .Name }}, error) {
	out := make([]*{{ .Name }})
	res := db.Find(&out)
	return out, res.Error
}

func get{{ $ModelName }}ByID(id int) (*{{ .Name }}, error) {	
	out := new({{ $ModelName }})
	res := db.First(&out, id)
	return out, res.Error
}

func update{{ $ModelName }}(id int, up *{{ .Name }}) error {
	res := db.Model(up).Update(&{{ $ModelName }}{
		ID: id,
		{{ range .Fields }}{{ .Name }}: in.{{ .Name }},
		{{ end }}
	})
	return res.Error
}

func delete{{ $ModelName }}(id int) error {
	res := db.Model(up).Delete(&{{ .Name }}{
		ID: id,
	})

	return res.Error
}

{{ end }}
