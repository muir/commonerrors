# commonerrors - tags for common causes of errors

[![GoDoc](https://godoc.org/github.com/muir/commonerrors?status.png)](https://pkg.go.dev/github.com/muir/commonerrors)
[![report card](https://goreportcard.com/badge/github.com/muir/commonerrors)](https://goreportcard.com/report/github.com/muir/commonerrors)

Install:

	go get github.com/muir/commonerrors

---

## Idea

Commonerrors is a collection of error tags. 

It offers:

- UsageError
- EnvironmentError
- ConfigurationError
- ProgrammerError
- LibraryError

## Basis

Right now, this package is based on github.com/pkg/errors.  That package's support of WithStack
is a bit broken so the base error package could change in the future.
