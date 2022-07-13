package schema

// extra config for our generator
#Config: {

	// metadata
	About: string
	Help:  string

	// generate client libs
	Client: go: enabled: bool | *true

	// generate a CLI
	Cli: enabled: bool | *true
}
