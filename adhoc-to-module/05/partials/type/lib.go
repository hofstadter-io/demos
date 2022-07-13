{{ $ModelName := camelT .name }}
func Create{{ $ModelName }}(in *{{ $ModelName }}) (*{{ $ModelName }}, error) {
	t := &{{ $ModelName }}{
		{{ range .Fields }}{{ camelT .name }}: in.{{ camelT .name }},
		{{ end }}
	} 
	res := db.Create(t)
	if res.Error != nil {
		return nil, res.Error
	}
	return t, nil
}

func List{{ $ModelName }}() ([]*{{ $ModelName }}, error) {
	out := make([]*{{ $ModelName }},0)
	res := db.Find(&out)
	return out, res.Error
}

func Get{{ $ModelName }}ByID(id string) (*{{ $ModelName }}, error) {	
	out := new({{ $ModelName }})
	res := db.First(&out, id)
	return out, res.Error
}

func Update{{ $ModelName }}(id string, up *{{ $ModelName }}) error {
	t := &{{ $ModelName }}{
		{{ range .Fields }}{{ camelT .name }}: up.{{ camelT .name }},
		{{ end }}
	}
	res := db.Model(up).Where("ID = ?", id).Updates(t)
	return res.Error
}

func Delete{{ $ModelName }}(id string) error {
	res := db.Delete(&{{ $ModelName }}{}, id)
	return res.Error
}

