## tc

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/tc)
[![Build Status](https://travis-ci.org/ferhatelmas/tc?branch=master)](https://travis-ci.org/ferhatelmas/tc)

> Check if TC No you have is valid for first, last name and birth year.

### Install

```
go get github.com/ferhatelmas/tc
```

### Usage

```go
import "github.com/ferhatelmas/tc"

tc.IsValid("17857715056", "ferhat", "elmas", 1988)
//=> true, nil

tc.IsValid("17857715055", "ferhat", "elmas", 1988)
//=> false, nil
```

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
