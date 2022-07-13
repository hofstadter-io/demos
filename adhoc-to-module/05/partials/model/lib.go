{{ $ModelName := camelT .Name }}
func create{{ $ModelName }}(in *{{ .Name }}) (*{{ .Name }}, error) {
	t := &{{ .Name }}{
		{{ range .Fields }}{{ .Name }}: in.{{ .Name }},
		{{ end }}
	} 
	res := db.Create(t)
	if res.Error != nil {
		return nil, res.Error
	}
	return t, nil
}

func list{{ $ModelName }}() ([]*{{ .Name }}, error) {
	out := make([]*{{ .Name }},0)
	res := db.Find(&out)
	return out, res.Error
}

func get{{ $ModelName }}ByID(id string) (*{{ .Name }}, error) {	
	out := new({{ .Name }})
	res := db.First(&out, id)
	return out, res.Error
}

func update{{ $ModelName }}(id string, up *{{ .Name }}) error {
	t := &{{ .Name }}{
		{{ range .Fields }}{{ .Name }}: up.{{ .Name }},
		{{ end }}
	}
	res := db.Model(up).Where("ID = ?", id).Updates(t)
	return res.Error
}

func delete{{ $ModelName }}(id string) error {
	res := db.Delete(&{{ .Name }}{}, id)
	return res.Error
}
