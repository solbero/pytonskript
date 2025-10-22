package exec

import (
	"io"

	"github.com/solbero/pytonskript/evaluator"
	"github.com/solbero/pytonskript/lexer"
	"github.com/solbero/pytonskript/object"
	"github.com/solbero/pytonskript/parser"
)

func Start(in io.Reader, out io.Writer) {
	env := object.NewEnvironment()
	bytes, err := readContents(in)
	if err != nil {
		panic(err)
	}

	s := string(bytes)
	l := lexer.New(s)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluator.Eval(program, env)
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func readContents(f io.Reader) ([]byte, error) {
	data := make([]byte, 0, 1024)
	for {
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}

		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
	}
}
