package {{ camel .Datamodel.Name }}

import (
	"gorm.io/gorm"
)

{{ range .Datamodel.Models }}
{{ $ModelName := camelT .name }}

func Create{{ $ModelName }}(in *{{ $ModelName }}) error {
	res := db.Create(&{{ $ModelName }}{
		{{ range .Fields }}{{ camelT .name }}: in.{{ camelT .name }},
		{{ end }}
	})
	return res.Error
}

func List{{ $ModelName }}() ([]*{{ $ModelName }}, error) {
	out := make([]*{{ $ModelName }})
	res := db.Find(&out)
	return out, res.Error
}

func Get{{ $ModelName }}ByID(id int) (*{{ $ModelName }}, error) {	
	out := new({{ $ModelName }})
	res := db.First(&out, id)
	return out, res.Error
}

func Update{{ $ModelName }}(id int, up *{{ $ModelName }}) error {
	res := db.Model(up).Update(&{{ $ModelName }}{
		ID: id,
		{{ range .Fields }}{{ camelT .name }}: in.{{ camelT .name }},
		{{ end }}
	})
	return res.Error
}

func Delete{{ $ModelName }}(id int) error {
	res := db.Model(up).Delete(&{{ $ModelName }}{
		ID: id,
	})

	return res.Error
}

{{ end }}
