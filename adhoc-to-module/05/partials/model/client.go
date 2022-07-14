// Client functions
{{ $ModelName := camelT .Name }}
func {{ $ModelName }}Create(input map[string]any) (*{{ .Name }}, error ) {
	data := new({{.Name}})
	url := host + "/{{ kebab .Name }}"
	
	client := req.C()
	_, err := client.R().
		SetBody(input).
		SetResult(&data).
		Post(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func {{ $ModelName }}List() ([]*{{ .Name }}, error ) {
	data := make([]*{{.Name}}, 0)
	url := host + "/{{ kebab .Name }}"
	
	client := req.C()
	_, err := client.R().
		SetResult(&data).
		Get(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func {{ $ModelName }}GetByID(id string) (*{{ .Name }}, error ) {
	data := new({{.Name}})
	url := host + "/{{ kebab .Name }}/{id}"
	
	client := req.C()
	_, err := client.R().
		SetPathParam("id", id).
		SetResult(&data).
		Get(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func {{ $ModelName }}Update(id string, input map[string]any) (*{{ .Name }}, error ) {
	data := new({{.Name}})
	url := host + "/{{ kebab .Name }}/{id}"
	
	client := req.C()
	_, err := client.R().
		SetPathParam("id", id).
		SetBody(input).
		SetResult(&data).
		Put(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func {{ $ModelName }}Delete(id string) error {

	url := host + "/{{ kebab .Name }}/{id}"
	client := req.C()
	_, err := client.R().
		SetPathParam("id", id).
		Delete(url)
	if err != nil {
		return err
	}

	return nil
}
