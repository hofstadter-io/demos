package {{ .Name }}

import (
	"gorm.io/gorm"
)

{{ range .Types }}
{{ $TypeName := camelT .name }}

func Create{{ $TypeName }}(in *{{ $TypeName }}) error {
	res := db.Create(&{{ $TypeName }}{
		{{ range .Fields }}{{ camelT .name }}: in.{{ camelT .name }},
		{{ end }}
	})
	return res.Error
}

func List{{ $TypeName }}() ([]*{{ $TypeName }}, error) {
	out := make([]*{{ $TypeName }})
	res := db.Find(&out)
	return out, res.Error
}

func Get{{ $TypeName }}ByID(id int) (*{{ $TypeName }}, error) {	
	out := new({{ $TypeName }})
	res := db.First(&out, id)
	return out, res.Error
}

func Update{{ $TypeName }}(id int, up *{{ $TypeName }}) error {
	res := db.Model(up).Update(&{{ $TypeName }}{
		ID: id,
		{{ range .Fields }}{{ camelT .name }}: in.{{ camelT .name }},
		{{ end }}
	})
	return res.Error
}

func Delete{{ $TypeName }}(id int) error {
	res := db.Model(up).Delete(&{{ $TypeName }}{
		ID: id,
	})

	return res.Error
}

{{ end }}
