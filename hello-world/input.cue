users: {
	// the cast
	tony: role:     "owner"
	doug: role:     "admin"
	tortoise: role: "user"
	achilles: role: "user"

	// extra / metadata
	tony: extra:     "worm"
	doug: extra:     "geb"
	tortoise: extra: "reconfiguring record"
	achilles: extra: "the (w)heelie mic drop"

	// enrich, extend, validate
	[n=string]: {
		// nest key as name
		name: n
		// construct email
		email: "\(n)@hof.io"
		// restrict role values
		role: "owner" | "admin" | "user"
	}
}
