/*
 * Copyright (C) 2019 Yunify, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this work except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ruleql

import (
	"context"
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/imneov/PostgreCRD/internal/parse/ruleql/parser"
	"strconv"
	"strings"
)

// QingQLListener construct expr
type QingQLListener struct {
	ctx context.Context
	*parser.BaseQingQLListener
	stack  []Expr
	expr   Expr
	errors []string
}

func (l *QingQLListener) appendErrorf(format string, a ...interface{}) {
	l.errors = append(l.errors, fmt.Sprintf(format, a...))
}

func (l *QingQLListener) push(i Expr) {
	l.stack = append(l.stack, i)
}

func (l *QingQLListener) pop() Expr {
	if len(l.stack) < 1 {
		//panic("stack is empty unable to pop")
		return nil
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

// Expr expr
func (l *QingQLListener) Expr() Expr {
	return l.stack[len(l.stack)-1]
}

func (l *QingQLListener) ExitRoot(c *parser.RootContext) {
	//fmt.Println("ExitRoot")
	r := &SelectStatementExpr{
		filter: &FilterExpr{},
	}
	if c.Group() != nil {
		expr := l.pop()
		switch expr := expr.(type) {
		case *DimensionsExpr:
			r.dimensions = expr
		default:
			l.appendErrorf("[+]parse group error[%s]", typeOf(expr))
		}
	}
	if c.Where() != nil {
		expr := l.pop()
		switch expr := expr.(type) {
		case *FilterExpr:
			r.filter = expr
		default:
			l.appendErrorf("[+]parse where error[%s]", typeOf(expr))
		}
	}
	if c.Topic() != nil {
		expr := l.pop()
		switch expr := expr.(type) {
		case TopicExpr:
			r.topic = expr
		default:
			l.appendErrorf("[+]parse topic error[%s]", typeOf(expr))
		}
	}
	if c.Fields() != nil {
		expr := l.pop()
		switch expr := expr.(type) {
		case FieldsExpr:
			r.fields = expr
		default:
			l.appendErrorf("[+]parse fields error[%s]", typeOf(expr))
		}
	}
	l.push(r)
}

// ExitSelect_list construct fields from select statement
func (l *QingQLListener) ExitFields(c *parser.FieldsContext) {
	//fmt.Println("ExitFields")
	var (
		size = len(c.AllField_elem())
		data FieldsExpr
	)
	data = make([]*FieldExpr, 0, size)
	for size > 0 {
		elem := l.pop()
		switch elem := elem.(type) {
		case *FieldExpr:
			data = append([]*FieldExpr{elem}, data...)
		}
		size--
	}
	l.push(data)
}

func (l *QingQLListener) ExitTopic(c *parser.TopicContext) {
	//fmt.Println("ExitTopic",c.GetText())
	n := len(c.AllTopic_item())
	d := len(c.AllDIV())
	list := make(TopicExpr, 0, n+1)
	if n == d {
		list = append(list, "")
	}
	for i := 0; i < n; i++ {
		elem := c.Topic_item(i)
		list = append(list, elem.GetText())
	}
	l.push(list)
}

func (l *QingQLListener) ExitWhere(c *parser.WhereContext) {
	//fmt.Println("ExitWhere",c.GetText())
	expr := l.pop()
	l.push(&FilterExpr{
		exp: expr,
	})
}

func (l *QingQLListener) ExitXpathField(c *parser.XpathFieldContext) {
	//fmt.Println("ExitXpathField")
	var alias string
	xpath := c.GetText()
	if xpath != "*" {
		arr := strings.Split(xpath, ".")
		alias = arr[len(arr)-1]
	}
	l.push(&FieldExpr{
		exp: &JSONPathExpr{
			xpath,
		},
		alias: alias,
	})
}

func (l *QingQLListener) ExitExprField(c *parser.ExprFieldContext) {
	//fmt.Println("ExitExprField", c.Expr(), c.Xpath_name())
	var alias string
	expr := l.pop()
	path := c.Xpath_name()
	if path != nil {
		alias = path.GetText()
		l.push(&FieldExpr{
			exp:   expr,
			alias: alias,
		})
	} else {
		//fmt.Println("+", expr, alias)
		l.push(&FieldExpr{
			exp:   expr,
			alias: "",
		})
	}
}

func (l *QingQLListener) ExitBinary(c *parser.BinaryContext) {
	right, left := l.pop(), l.pop()
	//fmt.Println("ExitBinary", c.GetText(), left, c.GetOp().GetText(), right)
	l.push(&BinaryExpr{
		Op:  c.GetOp().GetTokenType(),
		LHS: left,
		RHS: right,
	})
}

func (l *QingQLListener) ExitString(c *parser.StringContext) {
	//fmt.Println("ExitString", c.GetText())
	str := c.GetText()
	if len(str) >= 2 &&
		str[0] == '\'' &&
		str[len(str)-1] == '\'' {
		l.push(StringNode(str[1 : len(str)-1]))
	}
}

func (l *QingQLListener) ExitInteger(c *parser.IntegerContext) {
	//fmt.Println("ExitInteger", c.GetText())
	i, err := strconv.ParseInt(c.GetText(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	l.push(IntNode(i))
}

func (l *QingQLListener) ExitFloat(c *parser.FloatContext) {
	//fmt.Println("ExitFloat", c.GetText())
	i, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(err.Error())
	}
	l.push(FloatNode(i))
}

func (l *QingQLListener) ExitBoolean(c *parser.BooleanContext) {
	//fmt.Println("ExitBoolean", c.GetText())
	i, err := strconv.ParseBool(strings.ToLower(c.GetText()))
	if err != nil {
		panic(err.Error())
	}
	l.push(BoolNode(i))
}

func (l *QingQLListener) ExitXPath(c *parser.XPathContext) {
	//fmt.Println("ExitXpath", c.GetText())
	str := c.GetText()
	if str[0] != '"' {
		l.push(&JSONPathExpr{
			str,
		})
	} else {
		if str[len(str)-1] == '"' {
			l.push(&JSONPathExpr{
				str[1 : len(str)-1],
			})
		}
	}
	//error
}

func (l *QingQLListener) ExitCall_expr(c *parser.Call_exprContext) {
	//fmt.Println("ExitCall_expr", c.GetText(), c.AllExpr())
	n := len(c.AllExpr())
	temp := make([]Expr, 0, n)
	for i := 0; i < n; i++ {
		temp = append(temp, l.pop())
	}
	list := make([]Expr, 0, n)
	for i := n - 1; i >= 0; i-- {
		list = append(list, temp[i])
	}
	l.push(&CallExpr{
		raw:  c.GetText(),
		key:  c.GetKey().GetText(),
		args: list,
	})
}

func (l *QingQLListener) ExitSwitch_stmt(c *parser.Switch_stmtContext) {
	//fmt.Println("[-]ExitSwitch_stmt", c.GetText(), len(c.AllExpr()))
	n := len(c.AllExpr())
	expr := &SwitchExpr{
		list: make([]*CaseExpr, 0, n/2-1),
	}
	if n < 2 {
		return
	}
	if n%2 == 0 {
		expr.last = l.pop()
		n--
	} else {

	}
	for n > 1 {
		then, when := l.pop(), l.pop()
		expr.list = append(expr.list, &CaseExpr{
			then: then,
			when: when,
		})
		n -= 2
	}
	expr.exp = l.pop()

	l.push(expr)
}

func (l *QingQLListener) ExitFilter_condition(c *parser.Filter_conditionContext) {
	//fmt.Println("ExitFliter_condition", c.GetText(), c.AllAND())
	level := len(c.AllAND())
	left := l.pop()
	for level > 0 {
		right := l.pop()
		left = &BinaryExpr{
			Op:  parser.QingQLLexerAND,
			LHS: left,
			RHS: right,
		}
		level--
	}
	l.push(left)
}

func (l *QingQLListener) ExitFilter_condition_or(c *parser.Filter_condition_orContext) {
	//fmt.Println("ExitFilter_condition_or", c.GetText(), c.AllOR())
	level := len(c.AllOR())
	left := l.pop()
	for level > 0 {
		right := l.pop()
		left = &BinaryExpr{
			Op:  parser.QingQLLexerOR,
			LHS: left,
			RHS: right,
		}
		level--
	}
	l.push(left)
}

func (l *QingQLListener) ExitFilter_condition_not(c *parser.Filter_condition_notContext) {
	//fmt.Println("ExitFilter_condition_not", c.GetText())
	if c.NOT() != nil {
		right := l.pop()
		l.push(&BinaryExpr{
			Op:  parser.QingQLLexerNOT,
			LHS: nil,
			RHS: right,
		})
	}
}

func (l *QingQLListener) ExitGroup(c *parser.GroupContext) {
	//fmt.Println("ExitGroup", c.GetText())
}

func (l *QingQLListener) ExitGroup_by_elements(c *parser.Group_by_elementsContext) {
	//fmt.Println("ExitGroup_by_elements", c.GetText())
	//fmt.Println("ExitGroup", c.AllGroup_by_element())
	dimensions := &DimensionsExpr{}
	windowCnt := 0
	elements := c.AllGroup_by_element()
	for i := len(elements) - 1; i >= 0; i-- {
		expr := elements[i]
		switch expr.(type) {
		case *parser.GroupFieldContext:
			if expr, ok := l.pop().(*JSONPathExpr); ok {
				dimensions.exprs = append(dimensions.exprs, expr)
			}

		case *parser.GroupWindowsContext:
			if windowCnt > 0 {
				l.appendErrorf("parse group error, more windows[%s]", c.GetText())
			}
			windowCnt++
			if expr, ok := l.pop().(*WindowExpr); ok {
				dimensions.window = expr
			}
		default:
			l.appendErrorf("parse group error[%s]", c.GetText())
		}
	}
	l.push(dimensions)
}

func (l *QingQLListener) ExitGroup_by_element(c *parser.Group_by_elementContext) {
	//fmt.Println("ExitGroup_by_element", c.GetText())
}

func (l *QingQLListener) ExitGroupField(c *parser.GroupFieldContext) {
	//fmt.Println("ExitGroupField", c.GetText())
	str := c.GetText()
	var expr Expr
	if str[0] != '"' {
		expr = &JSONPathExpr{
			str,
		}
	} else {
		if str[len(str)-1] == '"' {
			expr = &JSONPathExpr{
				str[1 : len(str)-1],
			}
		}
	}
	l.push(expr)
}

func (l *QingQLListener) ExitGroup_window_unit(c *parser.Group_window_unitContext) {
	//fmt.Println("ExitGroup_window_unit", c.GetText())
}

func (l *QingQLListener) ExitGroup_window_length(c *parser.Group_window_lengthContext) {
	//fmt.Println("ExitGroup_window_length", c.GetText())
}

func (l *QingQLListener) ExitGroup_window_interval(c *parser.Group_window_intervalContext) {
	//fmt.Println("ExitGroup_window_interval", c.GetText())
}

func (l *QingQLListener) ExitGroup_tumbling_window(c *parser.Group_tumbling_windowContext) {
	//fmt.Println("ExitWindows", c.GetText())
	interval, err := strconv.ParseInt(c.Group_window_interval().GetText(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	window := &WindowExpr{
		WindowType: TUMBLING_WINDOW,
		Length:     WindowLength(1),
		Interval:   WindowInterval(interval),
	}
	l.push(window)

}

func (l *QingQLListener) ExitDotnotation(c *parser.DotnotationContext) {
	//fmt.Println("ExitDotnotation", c.GetText())
}

func (l *QingQLListener) ExitIdentifierWithQualifier(c *parser.IdentifierWithQualifierContext) {
	//fmt.Println("IdentifierWithQualifier", c.GetText())
}

func (l *QingQLListener) ExitIdentifierWithTOPICITEM(c *parser.IdentifierWithTOPICITEMContext) {
	//fmt.Println("IdentifierWithTOPICITEM", c.GetText())
}

func (l *QingQLListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//fmt.Println("SyntaxError", recognizer, offendingSymbol, line, column, msg, e)
	l.errors = append(l.errors, fmt.Sprintf("[%d:%d]%s", line, column, msg))

	return
}

func (l *QingQLListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportAmbiguity", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *QingQLListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportAttemptingFullContext", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *QingQLListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportContextSensitivity", recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func (l *QingQLListener) error() error {
	if len(l.errors) == 0 {
		return nil
	}
	return errors.New(strings.Join(l.errors, "\n"))
}
