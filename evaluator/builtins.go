// evaluator/builtins.go

package evaluator

import (
	"github.com/solbero/pyton/object"
)

var builtins = map[string]*object.Builtin{
	"lengde": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got %d, want 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to 'lengde' not supported, got %s", args[0].Type())
			}
		},
	},
	"første": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got %d, want 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'første' must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"siste": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got %d, want 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'siste' must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"resten": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got %d, want 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'resten' must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"dytt": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments, got %d, want 2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'dytt' must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"skriv": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				println(arg.Inspect())
			}

			return NULL
		},
	},
	"skjær": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments, got %d, want at least 2", len(args))
			} else if len(args) > 3 {
				return newError("wrong number of arguments, got %d, want less than 3", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to 'skjær' must be ARRAY, got %s", args[0].Type())
			}

			if args[1].Type() != object.INTEGER_OBJ {
				return newError("second argument to 'skjær' must be INTEGER, got %s", args[1].Type())
			}

			if len(args) == 3 && args[2].Type() != object.INTEGER_OBJ {
				return newError("third argument to 'skjær' must be INTEGER, got %s", args[2].Type())
			}

			arr := args[0].(*object.Array)
			start := args[1].(*object.Integer).Value
			stop := int64(len(arr.Elements))
			if len(args) == 3 {
				stop = args[2].(*object.Integer).Value
			}

			if start < 0 || stop > int64(len(arr.Elements)) || start > stop {
				return newError("invalid slice indices: start=%d, stop=%d", start, stop)
			}

			newElements := arr.Elements[start:stop]
			return &object.Array{Elements: newElements}
		},
	},
}
