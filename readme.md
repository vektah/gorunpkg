gorunpkg
========

Like go run, but instead of running a file it runs a package which can be pinned in vendor.

### why
A few reasons:
 - you can pin the version you use to generate code using `go dep`
 - you can have different versions per project
 - you dont need to install the binary on your path at all

This is particularly nice when committing go generated code so you are in control of when the generator updates.

### usage example

For example if you were using [campoy/jsonenums](http://github.com/campoy/jsonenums) to generate some code you would
install gorunpkg on your gopath. This is the only part that needs to be gopath wide and is small enough that it should be stable.
```bash
go get github.com/vektah/gorunpkg
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
