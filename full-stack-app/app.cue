package app

import (
	"github.com/hofstadter-io/demos/adhoc-to-module/04:demo"
)

// This is using your generator
App: demo.#DemoGenerator & {
	@gen(app)

	// inputs to the generator
	Name: "app"
	Module: "github.com/verdverm/app"

	// add our datamodel
	"Types": Types,	

	// other settings
	Outdir: "./out/"
}
