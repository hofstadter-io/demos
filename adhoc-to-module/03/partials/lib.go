{{ $TypeName := camelT .name }}
func Create{{ $TypeName }}(in *{{ $TypeName }}) (*{{ $TypeName }}, error) {
	t := &{{ $TypeName }}{
		{{ range .Fields }}{{ camelT .name }}: in.{{ camelT .name }},
		{{ end }}
	} 
	res := db.Create(t)
	if res.Error != nil {
		return nil, res.Error
	}
	return t, nil
}

func List{{ $TypeName }}() ([]*{{ $TypeName }}, error) {
	out := make([]*{{ $TypeName }},0)
	res := db.Find(&out)
	return out, res.Error
}

func Get{{ $TypeName }}ByID(id string) (*{{ $TypeName }}, error) {	
	out := new({{ $TypeName }})
	res := db.First(&out, id)
	return out, res.Error
}

func Update{{ $TypeName }}(id string, up *{{ $TypeName }}) error {
	t := &{{ $TypeName }}{
		{{ range .Fields }}{{ camelT .name }}: up.{{ camelT .name }},
		{{ end }}
	}
	res := db.Model(up).Where("ID = ?", id).Updates(t)
	return res.Error
}

func Delete{{ $TypeName }}(id string) error {
	res := db.Delete(&{{ $TypeName }}{}, id)
	return res.Error
}

