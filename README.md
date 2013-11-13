# TDD Code katas in the Go programming language

Includes following problems:

 * [Mars Rover](https://github.com/mriehl/go-tdd-katas/tree/master/go/rover)
 * [Fizzbuzz](https://github.com/mriehl/go-tdd-katas/tree/master/go/fizz)
 * [Fibonacci closure](https://github.com/mriehl/go-tdd-katas/tree/master/go/fibo)
 * [Picture generation](https://github.com/mriehl/go-tdd-katas/tree/master/go/pic)
 * [Roman numbers](https://github.com/mriehl/go-tdd-katas/tree/master/go/roman)
 * [Slice expansion](https://github.com/mriehl/go-tdd-katas/tree/master/go/slices)
 * [Square root computation](https://github.com/mriehl/go-tdd-katas/tree/master/go/sqrt)
 * [Wordcount](https://github.com/mriehl/go-tdd-katas/tree/master/go/wordcount)

## Dependencies
All katas were developed with test driven development and thus make heavy
usage of the packages
  * [github.com/stretchr/testify](https://github.com/stretchr/testify)
  * testing (builtin)

## Running
Tests are run with
```bash
go test
```
in the project directory.

For example you may run the `Mars Rover` kata tests with
```bash
go get github.com/mriehl/go-tdd-katas/go/rover
cd $GOPATH/src/github.com/mriehl/go-tdd-katas/go/rover
go test
```
