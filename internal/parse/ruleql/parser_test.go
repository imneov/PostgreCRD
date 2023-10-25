package ruleql

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseError(t *testing.T) {

	tests := []struct {
		expr   string
		hasErr bool
		err    interface{}
	}{
		// calc
//		{"select a from \n  aaaa", false, nil},
//		{"select a from \n  aaaa", false, nil},
//		{" 1 from /my/topic", true, `[1:1]missing SELECT at '1'
//[1:3]mismatched input 'from' expecting AS`},
//		{"select aaa", true, `3[1:1]missing SELECT at '1'
//[1:3]mismatched input 'from' expecting AS`},
		{"select 1 from /my/topic", true, `[1:9]mismatched input 'from' expecting AS`},
		{"select a from /my/topic where ", true, `[1:30]mismatched input '<EOF>' expecting {'(', '"', CASE, NOT, '*', TRUE, FALSE, INDENTIFIER, NUMBER, INTEGER, FLOAT, PATHITEM, STRING}
[+]parse topic error[FieldsExpr]
[+]parse fields error[<nil>]`},
		{"select a1() from /my/topic", true, `[1:12]mismatched input 'from' expecting AS`},
		{"select a(),b() from /my/topic", true, `[1:10]mismatched input ',' expecting AS
[1:15]mismatched input 'from' expecting AS`},
		{"select a() as a1,b()as a2 from /my/topic where a > 0 and ", true, `[1:57]mismatched input '<EOF>' expecting {'(', '"', CASE, NOT, '*', TRUE, FALSE, INDENTIFIER, NUMBER, INTEGER, FLOAT, PATHITEM, STRING}
[+]parse topic error[FieldsExpr]
[+]parse fields error[<nil>]`},
	}
	idx := 1
	for _, tt := range tests {
		t.Run("tt.name", func(t *testing.T) {
			_, err := Parse(tt.expr)
			fmt.Println("================================")
			fmt.Println(tt.expr)
			fmt.Println("--------------------------------")
			fmt.Println(err)
			fmt.Println("================================")
			if (err != nil) != tt.hasErr {
				t.Errorf("[%d] (%v), want (%v)", idx, err, tt.err)
			}
			if tt.hasErr && err.Error() != tt.err {
				t.Errorf("[%d] %v, want %v", idx, err, tt.err)
			}
			idx++
		})
	}
}

func TestParseNumber(t *testing.T) {

	tests := []struct {
		expr string
		want interface{}
	}{
		// calc
		{"1", IntNode(1)},
		{"0", IntNode(0)},
		{"-1", IntNode(-1)},
		{"1.0", FloatNode(1)},
		{"0.0", FloatNode(0)},
		{".0", FloatNode(0)},
		{"0.", FloatNode(0)},
		{"-1.0", FloatNode(-1)},
		{"1 + 1", IntNode(2)},
		{"1 - 1", IntNode(0)},
		{"7 * 3", IntNode(21)},
		{"7 / 3", IntNode(2)},
		{"1 - 1.0", FloatNode(0)},
		{"1.0 - 1", FloatNode(0)},
		{"3 * 7.0", FloatNode(21)},
		{"7.0 / 3.0", FloatNode(2.3333333333333335)},
	}

	for i, tt := range tests {
		expr, _ := ParseExpr(tt.expr)
		if got := eval(DefaultValue, expr); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("[%d] %v, want %v", i, got, tt.want)
		}
	}
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		expr string
		want bool
	}{
		// bool
		{"tRUE", true},
		{"fALSE", false},
		{"true", true},
		{"false", false},
		{"True", true},
		{"False", false},
		{"TRUE", true},
		{"FALSE", false},
	}
	for i, tt := range tests {
		expr, _ := ParseExpr(tt.expr)
		if got := eval(DefaultValue, expr); !reflect.DeepEqual(got, BoolNode(tt.want)) {
			t.Errorf("[%d] %v, want %v", i, got, tt.want)
		}
	}
}

func TestTopic(t *testing.T) {
	tests := []struct {
		name    string
		context Context
		args    string
		want    interface{}
	}{
		{"", NewJSONContext(JSONRaw.JSON), `select * from abc`, `abc`},
		{"", NewJSONContext(JSONRaw.JSON), `select * from abc/abc`, `abc/abc`},
		{"", NewJSONContext(JSONRaw.JSON), `select * from abc-def`, `abc-def`},
		{"", NewJSONContext(JSONRaw.JSON), `select * from abc/abc-def`, `abc/abc-def`},
		{"", NewJSONContext(JSONRaw.JSON), `select * from abc/000/ddd`, `abc/000/ddd`},
	}
	for idx, tt := range tests {
		Convey(fmt.Sprintf("Test Float [%d]%s", idx, tt.name), t, func() {
			expr, _ := Parse(tt.args)
			t, b := GetTopic(expr)
			So(b, ShouldBeTrue)
			So(t, ShouldEqual, tt.want)
		})
	}
}

func TestGroup(t *testing.T) {
	tests := []struct {
		name    string
		context Context
		args    string
		want    interface{}
	}{
		{"", NewJSONContext(JSONRaw.JSON), `select * from abc GROUP BY TUMBLINGWINDOW(ss, 10), f1`, `abc`},
	}
	for idx, tt := range tests {
		Convey(fmt.Sprintf("Test Float [%d]%s", idx, tt.name), t, func() {
			expr, _ := Parse(tt.args)
			t, b := GetTopic(expr)
			So(b, ShouldBeTrue)
			So(t, ShouldEqual, tt.want)
		})
	}
} //

func TestJson(t *testing.T) {
	tests := []struct {
		name    string
		context Context
		args    string
		want    interface{}
	}{
		{"", NewJSONContext(JSONRaw.JSON), `age`, IntNode(37)},
		{"", NewJSONContext(JSONRaw.JSON), `age + 1`, IntNode(38)},
		{"", NewJSONContext(JSONRaw.JSON), `(age + 1) * 2`, IntNode(76)},
		{"", NewJSONContext(JSONRaw.JSON), `(1 + 1) * age`, IntNode(74)},
		{"", NewJSONContext(JSONRaw.JSON), `name.first`, "Tom"},
		{"1", NewJSONContext(JSONRaw.JSON), `movie.1111`, "[{\"1111\": \"Tom\", \"last\": \"Anderson\"}]"},
		{"1", NewJSONContext(JSONRaw.JSON), `movie.1111[0].1111`, "Tom"},
		{"1", NewJSONContext(JSONRaw.JSON), `movie.1111[#].1111`, "[\"Tom\"]"},
		{"1", NewJSONContext(JSONRaw.JSON), `movie.1111[#]`, 1},
	}
	for idx, tt := range tests {
		Convey(fmt.Sprintf("Test Float [%d]%s", idx, tt.name), t, func() {
			expr, err := ParseExpr(tt.args)
			So(err, ShouldBeNil)
			So(eval(tt.context, expr), ShouldEqual, tt.want)
		})
	}
}

func TestGroupParse(t *testing.T) {
	tests := []struct {
		name string
		sql  string
		expr Expr
	}{

		{
			"",
			`select * from abc GROUP BY a1,b2,TUMBLINGWINDOW(ss, 10), f1, f2`,
			IntNode(37),
		},
	}
	for i, tt := range tests {
		expr, err := Parse(tt.sql)
		if err != nil {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.name, tt.sql, expr)
		}
	}
}

func ExamplePrint1() {
	// Setup the input
	sql := `select
			*,
			base64(*) as payload,
			temperature + '°F' as temperature.f,
			((temperature - 32) * 5 / 9.0) + '°C' as temperature.c,
			color as metadata.color,
			case "color"
			when 'red' then '红色'
			when 'green' then '绿色'
			else 'on' AS color_zh,
			ports[2] AS dev,
			metadata.name
			from /sys/user/+/msg/#
			where  
				color = 'red' and 
				temperature > 49
			GROUP BY f2, TUMBLINGWINDOW(ss, 10), f1
`

	expr, _ := Parse(sql)
	//Dump(expr)
	print(expr)
	// Output:

}
