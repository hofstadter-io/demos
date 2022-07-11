# hello-world

```sh
hof gen input.cue -T template.txt

hof gen input.cue -T template.txt=out.txt --watch
```

#### input.cue

```cue
users: {
	// the cast
	tony: role: "owner"
	doug: role: "admin"
	tortoise: role: "user"
	achilles: role: "user"

	// extra / metadata
	tony: extra: "worm"
	doug: extra: "geb"
	tortoise: extra: "reconfiguring record"
	achilles: extra: "the (w)heelie mic drop"

	// enrich, extend, validate
	[n=string]: { 
		// nest key as name
		name: n,
		// construct email
		email: "\(n)@hof.io"
		// restrict role values
		role: "owner" | "admin" | "user"
	}
}
```

#### template.txt

```
// starring...

{{ range .users -}}
{{ .name }} [{{ .role }}] ({{ .email }}) ... {{ .extra }}
{{ end }}

// fin...
```

