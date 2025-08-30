// evaluator/evaluator_test.go

package evaluator

import (
	"testing"

	"github.com/solbero/monkey/lexer"
	"github.com/solbero/monkey/object"
	"github.com/solbero/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30}, // 2 * (15)
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},                // 3 * (9) + 10
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50}, // (5 + 20 + 5) * 2 + -10
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		checkIntegerObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"hvis (sant) { 10 }", 10},
		{"hvis (falskt) { 10 }", nil},
		{"hvis (1) { 10 }", 10}, // 1 is truthy
		{"hvis (1 < 2) { 10 }", 10},
		{"hvis (1 > 2) { 10 }", nil},
		{"hvis (1 > 2) { 10 } ellers { 20 }", 20},
		{"hvis (1 < 2) { 10 } ellers { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			checkIntegerObject(t, evaluated, int64(integer))
		} else {
			checkNullObject(t, evaluated)
		}
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sant", true},
		{"falskt", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"sant == sant", true},
		{"falskt == falskt", true},
		{"sant == falskt", false},
		{"sant != falskt", true},
		{"falskt != sant", true},
		{"(1 < 2) == sant", true},
		{"(1 < 2) == falskt", false},
		{"(1 > 2) == sant", false},
		{"(1 > 2) == falskt", true},
		{`"hello" == "hello"`, true},
		{`"hello" == "goodbye"`, false},
		{`"hello" != "hello"`, false},
		{`"hello" != "goodbye"`, true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		checkBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!sant", false},
		{"!falskt", true},
		{"!5", false}, // 5 is truthy
		{"!!sant", true},
		{"!!falskt", false},
		{"!!5", true}, // 5 is truthy
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		checkBooleanObject(t, evaluated, tt.expected)
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"returner 10;", 10},
		{"returner 10; 9;", 10}, // 9; is ignored
		{"returner 2 * 5; 9;", 10},
		{"9; returner 2 * 5; 9;", 10}, // 9; is ignored
		{"hvis (10 > 1) {hvis (10 > 1) {returner 10;} returner 1}", 10},
		{"la f = funksjon(x) {returner x; x + 10;}; f(10);", 10}, // x + 10; is ignored
		{"la f = funksjon(x) {la result = x + 10; returner result; returner 10;}; f(10);", 20}, // return 10; is ignored
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		checkIntegerObject(t, evaluated, tt.expected)
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"la a = 5; a;", 5},
		{"la a = 5 * 5; a;", 25},
		{"la a = 5; la b = a; b;", 5},
		{"la a = 5; la b = a; la c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		checkIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "funksjon(x) { x + 2; };"
	evaluated := testEval(input)

	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function, got %T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters, got %d, want 1", len(fn.Parameters))
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x', got %q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q, got %q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"la identity = funksjon(x) { x; }; identity(5);", 5},
		{"la identity = funksjon(x) { returner x; }; identity(5);", 5}, // return statement
		{"la double = funksjon(x) { x * 2; }; double(5);", 10},
		{"la add = funksjon(x, y) { x + y; }; add(5, 5);", 10},
		{"la add = funksjon(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20}, // nested function calls
		{"funksjon(x) { x; }(5)", 5},                                        // immediately invoked function expression
	}

	for _, tt := range tests {
		checkIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
	la newAdder = funksjon(x) {
		funksjon(y) { x + y };
	};

	la addTwo = newAdder(2);
	addTwo(2);
	`
	checkIntegerObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: `"hello world"`, expected: "hello world"},
		{input: `"hello \"world\""`, expected: "hello \"world\""},
		{input: `"hello\nworld"`, expected: "hello\nworld"},
		{input: `"hello\t\t\tworld"`, expected: "hello\t\t\tworld"},
		{input: `"hello\\world"`, expected: "hello\\world"},
		{input: `"hello\bworld"`, expected: "helloworld"},
		{input: `"Hello" + " " + "World!"`, expected: "Hello World!"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		str, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String, got %T (%+v)", evaluated, evaluated)
		}

		if str.Value != tt.expected {
			t.Errorf("String has wrong value, expected %q got %q", tt.expected, str.Value)
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{input: `lengde("")`, expected: 0},
		{input: `lengde("four")`, expected: 4},
		{input: `lengde("hello world")`, expected: 11},
		{input: `lengde(1)`, expected: "argument to 'lengde' not supported, got INTEGER"},
		{input: `lengde("one", "two")`, expected: "wrong number of arguments, got 2, want 1"},
		{input: `første([1, 2, 3])`, expected: 1},
		{input: `første([])`, expected: nil},
		{input: `første(1)`, expected: "argument to 'første' must be ARRAY, got INTEGER"},
		{input: `første([1, 2], [3, 4])`, expected: "wrong number of arguments, got 2, want 1"},
		{input: `siste([1, 2, 3])`, expected: 3},
		{input: `siste([])`, expected: nil},
		{input: `siste(1)`, expected: "argument to 'siste' must be ARRAY, got INTEGER"},
		{input: `siste([1, 2], [3, 4])`, expected: "wrong number of arguments, got 2, want 1"},
		{input: `resten([1, 2, 3])`, expected: []int64{2, 3}},
		{input: `resten([])`, expected: nil},
		{input: `resten(1)`, expected: "argument to 'resten' must be ARRAY, got INTEGER"},
		{input: `resten([1, 2], [3, 4])`, expected: "wrong number of arguments, got 2, want 1"},
		{input: `dytt([], 1)`, expected: []int64{1}},
		{input: `dytt(1, 1)`, expected: "argument to 'dytt' must be ARRAY, got INTEGER"},
		{input: `dytt([1, 2], 1, 2)`, expected: "wrong number of arguments, got 3, want 2"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			checkIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error, got %T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message, expected %q, got %q", expected, errObj.Message)
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array, got %T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements, got %d", len(result.Elements))
	}

	checkIntegerObject(t, result.Elements[0], 1)
	checkIntegerObject(t, result.Elements[1], 4)
	checkIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{input: "[1, 2, 3][0]", expected: 1},
		{input: "[1, 2, 3][1]", expected: 2},
		{input: "[1, 2, 3][2]", expected: 3},
		{input: "la i = 0; [1][i];", expected: 1},
		{input: "[1, 2, 3][1 + 1];", expected: 3},
		{input: "la myArray = [1, 2, 3]; myArray[2];", expected: 3},
		{input: "la myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];", expected: 6},
		{input: "la myArray = [1, 2, 3]; la i = myArray[0]; myArray[i]", expected: 2},
		{input: "[1, 2, 3][3]", expected: nil},
		{input: "[1, 2, 3][-1]", expected: nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			checkIntegerObject(t, evaluated, int64(integer))
			continue
		}

		checkNullObject(t, evaluated)
	}
}

func TestHashLiterals(t *testing.T) {
	input := `la two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		sant: 5,
		falskt: 6
	}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash, got %T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs, got %d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		checkIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{input: `{"foo": 5}["foo"]`, expected: 5},
		{input: `{"foo": 5}["bar"]`, expected: nil},
		{input: `la key = "foo"; {"foo": 5}[key]`, expected: 5},
		{input: `{5: 5}[5]`, expected: 5},
		{input: `{sant: 5}[sant]`, expected: 5},
		{input: `{falskt: 5}[falskt]`, expected: 5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			checkIntegerObject(t, evaluated, int64(integer))
			continue
		}

		checkNullObject(t, evaluated)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input       string
		expectedMsg string
	}{
		{"5 + sant;", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + sant; 5;", "type mismatch: INTEGER + BOOLEAN"},
		{"-sant", "unknown operator: -BOOLEAN"},
		{"sant + falskt;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; sant + falskt; 5", "unknown operator: BOOLEAN + BOOLEAN"},
		{"hvis (10 > 1) { sant + falskt; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{"hvis (10 > 1) { hvis (10 > 1) { returner sant + falskt; } returner 1 }", "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "identifier not found: foobar"},
		{`"Hello" - "World!"`, "unknown operator: STRING - STRING"},
		{`{"name": "Monkey"}[funksjon(x) { x }];`, "unusable as hash key: FUNCTION"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned, got %T (%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMsg {
			t.Errorf("wrong error message, expected %q, got %q", tt.expectedMsg, errObj.Message)
		}
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	env := object.NewEnvironment()
	program := p.ParseProgram()
	return Eval(program, env)
}

func checkNullObject(t *testing.T, obj object.Object) bool {
	t.Helper()
	if obj != NULL {
		t.Errorf("object is not NULL, got %T (%+v)", obj, obj)
		return false
	}
	return true
}

func checkIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	t.Helper()
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer, got %T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value, got %d want %d", result.Value, expected)
		return false
	}
	return true
}

func checkBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	t.Helper()
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean, got %T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value, got %t want %t", result.Value, expected)
		return false
	}
	return true
}
