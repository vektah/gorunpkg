gorunpkg
========

Like go run, but with go builds package resolution.

Very useful in go:generate stanzas to use the currently vendored version of a command.

For example if you were using [campoy/jsonenums](http://github.com/campoy/jsonenums) to generate some code you could:
```go
//go:generate gorunpkg github.com/campoy/jsonenums -type=Pill

package painkiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
```

and gorunpkg would find jsonenums installed in vendor, build it, then run it with the args ` -type=Pill`.
