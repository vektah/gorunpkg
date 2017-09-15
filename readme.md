gorunpkg
========

Like go run, but with go builds package resolution.

### why
So the cli dependency is under version control per-project, rather than being a random gopath wide executable.

### usage example

For example if you were using [campoy/jsonenums](http://github.com/campoy/jsonenums) to generate some code you would
install gorunpkg on your gopath. This is the only part that needs to be gopath wide and is small enough that it should be stable.
```bash
go get github.com/Vektah/gorunpkg
```

add the dep to projects Gopkg.toml:
```toml
required = ["github.com/campoy/jsonenums"]
```

fetch and pin the dep:
```bash
dep ensure
```

add some code using it:
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

then go generate:
```bash
go generate ./...
```
