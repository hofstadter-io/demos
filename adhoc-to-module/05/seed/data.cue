Users: [ for _, u in _users { u }]
_users: {
	// the cast
	tony: Role:     "owner"
	doug: Role:     "admin"
	tortoise: Role: "user"
	achilles: Role: "user"

	// enrich, extend, validate
	[n=string]: {
		// nest key as name
		Name: n
		// construct email
		Email: "\(n)@hof.io"
		// restrict role values
		Role: "owner" | "admin" | "user"
	}

	tony: Posts: [{
		Title: "hello world"
		Content: "my first post"
	},{
		Title: "ipsum lorem"
		Content: "beep boop bop"
	}]

	doug: Posts: [{
		Title: "My favorite pokemon"
		Content: "it's MIU, if you don't get it, read GEB"
	}]
	achilles: Posts: [{
		Title: "they got me"
		Content: "damn archers, not sure i'll ever run again Mr. T"
	}]
}
