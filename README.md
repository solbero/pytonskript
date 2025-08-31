# Writing an interpreter in Go

This repository contains my implementation of the Monkey interpreter as described in the book [Writing an Interpreter in Go](https://interpreterbook.com/) and the addition [The Lost Chapter: A Macro System For Monkey](https://interpreterbook.com/lost/).

Monkey is not an official language and has no official specification. It is a toy language that is used to teach the basics of writing an interpreter. You can read more about the Monkey language and see other implementations [Monkey: The programming language that lives in books](https://monkeylang.org/).

## Installation

```go
go install github.com/solbero/pyton@latest
```

## Usage

```bash
$ go run main.go
Hello mrnugget! This is the Monkey programming language!
Feel free to type in commands
>> let answer = 6 * 7;
>> answer;
42
```

## License

MIT License
