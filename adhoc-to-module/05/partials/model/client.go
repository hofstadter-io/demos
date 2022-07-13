{{ if .Config.Client.go.enabled }}
{{ $ModelName := camelT .Name }}
func {{ $ModelName }}Create() {

}
func {{ $ModelName }}List() {

}
func {{ $ModelName }}GetByID(id string) {

}
func {{ $ModelName }}Update(id string) {

}
func {{ $ModelName }}Delete(id string) {

}
{{ end }}

