## tc

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/tc)
[![Build Status](https://travis-ci.org/ferhatelmas/tc?branch=master)](https://travis-ci.org/ferhatelmas/tc)

> Check if Turkish Republic Identification Number you have is valid for first, last name and birth year.

### Install

```
go get github.com/ferhatelmas/tc
```

### Usage

```go
import "github.com/ferhatelmas/tc"

tc.IsValid("17857715056")
//=> true

tc.IsValid("17857715050")
//=> false

tc.IsValidFor("17857715056", "ferhat", "elmas", 1988)
//=> true, nil

tc.IsValidFor("17857715055", "ferhat", "elmas", 1988)
//=> false, nil
```

`IsValid` checks if number is a possible correct identification number.

`IsValidFor` checks if number belongs to given the person identified by given first name, last name and birth year.

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
