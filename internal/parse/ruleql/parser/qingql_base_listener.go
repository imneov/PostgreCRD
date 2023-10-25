// Generated from QingQL.g4 by ANTLR 4.7.

package parser // QingQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQingQLListener is a complete listener for a parse tree produced by QingQLParser.
type BaseQingQLListener struct{}

var _ QingQLListener = &BaseQingQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQingQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQingQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQingQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQingQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseQingQLListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseQingQLListener) ExitRoot(ctx *RootContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseQingQLListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseQingQLListener) ExitFields(ctx *FieldsContext) {}

// EnterExprField is called when production ExprField is entered.
func (s *BaseQingQLListener) EnterExprField(ctx *ExprFieldContext) {}

// ExitExprField is called when production ExprField is exited.
func (s *BaseQingQLListener) ExitExprField(ctx *ExprFieldContext) {}

// EnterXpathField is called when production XpathField is entered.
func (s *BaseQingQLListener) EnterXpathField(ctx *XpathFieldContext) {}

// ExitXpathField is called when production XpathField is exited.
func (s *BaseQingQLListener) ExitXpathField(ctx *XpathFieldContext) {}

// EnterTopic is called when production topic is entered.
func (s *BaseQingQLListener) EnterTopic(ctx *TopicContext) {}

// ExitTopic is called when production topic is exited.
func (s *BaseQingQLListener) ExitTopic(ctx *TopicContext) {}

// EnterTopic_item is called when production topic_item is entered.
func (s *BaseQingQLListener) EnterTopic_item(ctx *Topic_itemContext) {}

// ExitTopic_item is called when production topic_item is exited.
func (s *BaseQingQLListener) ExitTopic_item(ctx *Topic_itemContext) {}

// EnterWhere is called when production where is entered.
func (s *BaseQingQLListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BaseQingQLListener) ExitWhere(ctx *WhereContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseQingQLListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseQingQLListener) ExitFilter(ctx *FilterContext) {}

// EnterFilter_condition is called when production filter_condition is entered.
func (s *BaseQingQLListener) EnterFilter_condition(ctx *Filter_conditionContext) {}

// ExitFilter_condition is called when production filter_condition is exited.
func (s *BaseQingQLListener) ExitFilter_condition(ctx *Filter_conditionContext) {}

// EnterFilter_condition_or is called when production filter_condition_or is entered.
func (s *BaseQingQLListener) EnterFilter_condition_or(ctx *Filter_condition_orContext) {}

// ExitFilter_condition_or is called when production filter_condition_or is exited.
func (s *BaseQingQLListener) ExitFilter_condition_or(ctx *Filter_condition_orContext) {}

// EnterFilter_condition_not is called when production filter_condition_not is entered.
func (s *BaseQingQLListener) EnterFilter_condition_not(ctx *Filter_condition_notContext) {}

// ExitFilter_condition_not is called when production filter_condition_not is exited.
func (s *BaseQingQLListener) ExitFilter_condition_not(ctx *Filter_condition_notContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseQingQLListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseQingQLListener) ExitGroup(ctx *GroupContext) {}

// EnterGroup_by_elements is called when production group_by_elements is entered.
func (s *BaseQingQLListener) EnterGroup_by_elements(ctx *Group_by_elementsContext) {}

// ExitGroup_by_elements is called when production group_by_elements is exited.
func (s *BaseQingQLListener) ExitGroup_by_elements(ctx *Group_by_elementsContext) {}

// EnterGroupWindows is called when production GroupWindows is entered.
func (s *BaseQingQLListener) EnterGroupWindows(ctx *GroupWindowsContext) {}

// ExitGroupWindows is called when production GroupWindows is exited.
func (s *BaseQingQLListener) ExitGroupWindows(ctx *GroupWindowsContext) {}

// EnterGroupField is called when production GroupField is entered.
func (s *BaseQingQLListener) EnterGroupField(ctx *GroupFieldContext) {}

// ExitGroupField is called when production GroupField is exited.
func (s *BaseQingQLListener) ExitGroupField(ctx *GroupFieldContext) {}

// EnterGroup_window_unit is called when production group_window_unit is entered.
func (s *BaseQingQLListener) EnterGroup_window_unit(ctx *Group_window_unitContext) {}

// ExitGroup_window_unit is called when production group_window_unit is exited.
func (s *BaseQingQLListener) ExitGroup_window_unit(ctx *Group_window_unitContext) {}

// EnterGroup_window_length is called when production group_window_length is entered.
func (s *BaseQingQLListener) EnterGroup_window_length(ctx *Group_window_lengthContext) {}

// ExitGroup_window_length is called when production group_window_length is exited.
func (s *BaseQingQLListener) ExitGroup_window_length(ctx *Group_window_lengthContext) {}

// EnterGroup_window_interval is called when production group_window_interval is entered.
func (s *BaseQingQLListener) EnterGroup_window_interval(ctx *Group_window_intervalContext) {}

// ExitGroup_window_interval is called when production group_window_interval is exited.
func (s *BaseQingQLListener) ExitGroup_window_interval(ctx *Group_window_intervalContext) {}

// EnterGroup_windows is called when production group_windows is entered.
func (s *BaseQingQLListener) EnterGroup_windows(ctx *Group_windowsContext) {}

// ExitGroup_windows is called when production group_windows is exited.
func (s *BaseQingQLListener) ExitGroup_windows(ctx *Group_windowsContext) {}

// EnterGroup_tumbling_window is called when production group_tumbling_window is entered.
func (s *BaseQingQLListener) EnterGroup_tumbling_window(ctx *Group_tumbling_windowContext) {}

// ExitGroup_tumbling_window is called when production group_tumbling_window is exited.
func (s *BaseQingQLListener) ExitGroup_tumbling_window(ctx *Group_tumbling_windowContext) {}

// EnterGroup_hopping_window is called when production group_hopping_window is entered.
func (s *BaseQingQLListener) EnterGroup_hopping_window(ctx *Group_hopping_windowContext) {}

// ExitGroup_hopping_window is called when production group_hopping_window is exited.
func (s *BaseQingQLListener) ExitGroup_hopping_window(ctx *Group_hopping_windowContext) {}

// EnterGroup_sliding_window is called when production group_sliding_window is entered.
func (s *BaseQingQLListener) EnterGroup_sliding_window(ctx *Group_sliding_windowContext) {}

// ExitGroup_sliding_window is called when production group_sliding_window is exited.
func (s *BaseQingQLListener) ExitGroup_sliding_window(ctx *Group_sliding_windowContext) {}

// EnterGroup_session_window is called when production group_session_window is entered.
func (s *BaseQingQLListener) EnterGroup_session_window(ctx *Group_session_windowContext) {}

// ExitGroup_session_window is called when production group_session_window is exited.
func (s *BaseQingQLListener) ExitGroup_session_window(ctx *Group_session_windowContext) {}

// EnterFunction is called when production Function is entered.
func (s *BaseQingQLListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production Function is exited.
func (s *BaseQingQLListener) ExitFunction(ctx *FunctionContext) {}

// EnterBraces is called when production Braces is entered.
func (s *BaseQingQLListener) EnterBraces(ctx *BracesContext) {}

// ExitBraces is called when production Braces is exited.
func (s *BaseQingQLListener) ExitBraces(ctx *BracesContext) {}

// EnterSwitch is called when production Switch is entered.
func (s *BaseQingQLListener) EnterSwitch(ctx *SwitchContext) {}

// ExitSwitch is called when production Switch is exited.
func (s *BaseQingQLListener) ExitSwitch(ctx *SwitchContext) {}

// EnterBinary is called when production Binary is entered.
func (s *BaseQingQLListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production Binary is exited.
func (s *BaseQingQLListener) ExitBinary(ctx *BinaryContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BaseQingQLListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BaseQingQLListener) ExitBoolean(ctx *BooleanContext) {}

// EnterInteger is called when production Integer is entered.
func (s *BaseQingQLListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production Integer is exited.
func (s *BaseQingQLListener) ExitInteger(ctx *IntegerContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseQingQLListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseQingQLListener) ExitFloat(ctx *FloatContext) {}

// EnterString is called when production String is entered.
func (s *BaseQingQLListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseQingQLListener) ExitString(ctx *StringContext) {}

// EnterXPath is called when production XPath is entered.
func (s *BaseQingQLListener) EnterXPath(ctx *XPathContext) {}

// ExitXPath is called when production XPath is exited.
func (s *BaseQingQLListener) ExitXPath(ctx *XPathContext) {}

// EnterSwitch_stmt is called when production switch_stmt is entered.
func (s *BaseQingQLListener) EnterSwitch_stmt(ctx *Switch_stmtContext) {}

// ExitSwitch_stmt is called when production switch_stmt is exited.
func (s *BaseQingQLListener) ExitSwitch_stmt(ctx *Switch_stmtContext) {}

// EnterCall_expr is called when production call_expr is entered.
func (s *BaseQingQLListener) EnterCall_expr(ctx *Call_exprContext) {}

// ExitCall_expr is called when production call_expr is exited.
func (s *BaseQingQLListener) ExitCall_expr(ctx *Call_exprContext) {}

// EnterAsterisk is called when production asterisk is entered.
func (s *BaseQingQLListener) EnterAsterisk(ctx *AsteriskContext) {}

// ExitAsterisk is called when production asterisk is exited.
func (s *BaseQingQLListener) ExitAsterisk(ctx *AsteriskContext) {}

// EnterXpath_name is called when production xpath_name is entered.
func (s *BaseQingQLListener) EnterXpath_name(ctx *Xpath_nameContext) {}

// ExitXpath_name is called when production xpath_name is exited.
func (s *BaseQingQLListener) ExitXpath_name(ctx *Xpath_nameContext) {}

// EnterDotnotation is called when production dotnotation is entered.
func (s *BaseQingQLListener) EnterDotnotation(ctx *DotnotationContext) {}

// ExitDotnotation is called when production dotnotation is exited.
func (s *BaseQingQLListener) ExitDotnotation(ctx *DotnotationContext) {}

// EnterIdentifierWithTOPICITEM is called when production identifierWithTOPICITEM is entered.
func (s *BaseQingQLListener) EnterIdentifierWithTOPICITEM(ctx *IdentifierWithTOPICITEMContext) {}

// ExitIdentifierWithTOPICITEM is called when production identifierWithTOPICITEM is exited.
func (s *BaseQingQLListener) ExitIdentifierWithTOPICITEM(ctx *IdentifierWithTOPICITEMContext) {}

// EnterIdentifierWithQualifier is called when production identifierWithQualifier is entered.
func (s *BaseQingQLListener) EnterIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}

// ExitIdentifierWithQualifier is called when production identifierWithQualifier is exited.
func (s *BaseQingQLListener) ExitIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}
