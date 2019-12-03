go mod init example.com/hello
go list -m all

From the output of go list -m all, we can see we're using an untagged version of golang.org/x/text. Let's upgrade to the latest tagged version and test that everything still works:

go get golang.org/x/text
output:-
go: finding golang.org/x/text v0.3.0

Let's list the available tagged versions of that module:-
go list -m -versions rsc.io/sampler

go doc rsc.io/quote/v3
`
```
go mod init creates a new module, initializing the go.mod file that describes it.
go build, go test, and other package-building commands add new dependencies to go.mod as needed.
go list -m all prints the current module’s dependencies.
go get changes the required version of a dependency (or adds a new dependency).
go mod tidy removes unused dependencies.
```
