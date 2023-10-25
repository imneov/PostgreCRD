// Generated from QingQL.g4 by ANTLR 4.7.

package parser // QingQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QingQLListener is a complete listener for a parse tree produced by QingQLParser.
type QingQLListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterExprField is called when entering the ExprField production.
	EnterExprField(c *ExprFieldContext)

	// EnterXpathField is called when entering the XpathField production.
	EnterXpathField(c *XpathFieldContext)

	// EnterTopic is called when entering the topic production.
	EnterTopic(c *TopicContext)

	// EnterTopic_item is called when entering the topic_item production.
	EnterTopic_item(c *Topic_itemContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterFilter_condition is called when entering the filter_condition production.
	EnterFilter_condition(c *Filter_conditionContext)

	// EnterFilter_condition_or is called when entering the filter_condition_or production.
	EnterFilter_condition_or(c *Filter_condition_orContext)

	// EnterFilter_condition_not is called when entering the filter_condition_not production.
	EnterFilter_condition_not(c *Filter_condition_notContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterGroup_by_elements is called when entering the group_by_elements production.
	EnterGroup_by_elements(c *Group_by_elementsContext)

	// EnterGroupWindows is called when entering the GroupWindows production.
	EnterGroupWindows(c *GroupWindowsContext)

	// EnterGroupField is called when entering the GroupField production.
	EnterGroupField(c *GroupFieldContext)

	// EnterGroup_window_unit is called when entering the group_window_unit production.
	EnterGroup_window_unit(c *Group_window_unitContext)

	// EnterGroup_window_length is called when entering the group_window_length production.
	EnterGroup_window_length(c *Group_window_lengthContext)

	// EnterGroup_window_interval is called when entering the group_window_interval production.
	EnterGroup_window_interval(c *Group_window_intervalContext)

	// EnterGroup_windows is called when entering the group_windows production.
	EnterGroup_windows(c *Group_windowsContext)

	// EnterGroup_tumbling_window is called when entering the group_tumbling_window production.
	EnterGroup_tumbling_window(c *Group_tumbling_windowContext)

	// EnterGroup_hopping_window is called when entering the group_hopping_window production.
	EnterGroup_hopping_window(c *Group_hopping_windowContext)

	// EnterGroup_sliding_window is called when entering the group_sliding_window production.
	EnterGroup_sliding_window(c *Group_sliding_windowContext)

	// EnterGroup_session_window is called when entering the group_session_window production.
	EnterGroup_session_window(c *Group_session_windowContext)

	// EnterFunction is called when entering the Function production.
	EnterFunction(c *FunctionContext)

	// EnterBraces is called when entering the Braces production.
	EnterBraces(c *BracesContext)

	// EnterSwitch is called when entering the Switch production.
	EnterSwitch(c *SwitchContext)

	// EnterBinary is called when entering the Binary production.
	EnterBinary(c *BinaryContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterInteger is called when entering the Integer production.
	EnterInteger(c *IntegerContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterXPath is called when entering the XPath production.
	EnterXPath(c *XPathContext)

	// EnterSwitch_stmt is called when entering the switch_stmt production.
	EnterSwitch_stmt(c *Switch_stmtContext)

	// EnterCall_expr is called when entering the call_expr production.
	EnterCall_expr(c *Call_exprContext)

	// EnterAsterisk is called when entering the asterisk production.
	EnterAsterisk(c *AsteriskContext)

	// EnterXpath_name is called when entering the xpath_name production.
	EnterXpath_name(c *Xpath_nameContext)

	// EnterDotnotation is called when entering the dotnotation production.
	EnterDotnotation(c *DotnotationContext)

	// EnterIdentifierWithTOPICITEM is called when entering the identifierWithTOPICITEM production.
	EnterIdentifierWithTOPICITEM(c *IdentifierWithTOPICITEMContext)

	// EnterIdentifierWithQualifier is called when entering the identifierWithQualifier production.
	EnterIdentifierWithQualifier(c *IdentifierWithQualifierContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitExprField is called when exiting the ExprField production.
	ExitExprField(c *ExprFieldContext)

	// ExitXpathField is called when exiting the XpathField production.
	ExitXpathField(c *XpathFieldContext)

	// ExitTopic is called when exiting the topic production.
	ExitTopic(c *TopicContext)

	// ExitTopic_item is called when exiting the topic_item production.
	ExitTopic_item(c *Topic_itemContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitFilter_condition is called when exiting the filter_condition production.
	ExitFilter_condition(c *Filter_conditionContext)

	// ExitFilter_condition_or is called when exiting the filter_condition_or production.
	ExitFilter_condition_or(c *Filter_condition_orContext)

	// ExitFilter_condition_not is called when exiting the filter_condition_not production.
	ExitFilter_condition_not(c *Filter_condition_notContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitGroup_by_elements is called when exiting the group_by_elements production.
	ExitGroup_by_elements(c *Group_by_elementsContext)

	// ExitGroupWindows is called when exiting the GroupWindows production.
	ExitGroupWindows(c *GroupWindowsContext)

	// ExitGroupField is called when exiting the GroupField production.
	ExitGroupField(c *GroupFieldContext)

	// ExitGroup_window_unit is called when exiting the group_window_unit production.
	ExitGroup_window_unit(c *Group_window_unitContext)

	// ExitGroup_window_length is called when exiting the group_window_length production.
	ExitGroup_window_length(c *Group_window_lengthContext)

	// ExitGroup_window_interval is called when exiting the group_window_interval production.
	ExitGroup_window_interval(c *Group_window_intervalContext)

	// ExitGroup_windows is called when exiting the group_windows production.
	ExitGroup_windows(c *Group_windowsContext)

	// ExitGroup_tumbling_window is called when exiting the group_tumbling_window production.
	ExitGroup_tumbling_window(c *Group_tumbling_windowContext)

	// ExitGroup_hopping_window is called when exiting the group_hopping_window production.
	ExitGroup_hopping_window(c *Group_hopping_windowContext)

	// ExitGroup_sliding_window is called when exiting the group_sliding_window production.
	ExitGroup_sliding_window(c *Group_sliding_windowContext)

	// ExitGroup_session_window is called when exiting the group_session_window production.
	ExitGroup_session_window(c *Group_session_windowContext)

	// ExitFunction is called when exiting the Function production.
	ExitFunction(c *FunctionContext)

	// ExitBraces is called when exiting the Braces production.
	ExitBraces(c *BracesContext)

	// ExitSwitch is called when exiting the Switch production.
	ExitSwitch(c *SwitchContext)

	// ExitBinary is called when exiting the Binary production.
	ExitBinary(c *BinaryContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitInteger is called when exiting the Integer production.
	ExitInteger(c *IntegerContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitXPath is called when exiting the XPath production.
	ExitXPath(c *XPathContext)

	// ExitSwitch_stmt is called when exiting the switch_stmt production.
	ExitSwitch_stmt(c *Switch_stmtContext)

	// ExitCall_expr is called when exiting the call_expr production.
	ExitCall_expr(c *Call_exprContext)

	// ExitAsterisk is called when exiting the asterisk production.
	ExitAsterisk(c *AsteriskContext)

	// ExitXpath_name is called when exiting the xpath_name production.
	ExitXpath_name(c *Xpath_nameContext)

	// ExitDotnotation is called when exiting the dotnotation production.
	ExitDotnotation(c *DotnotationContext)

	// ExitIdentifierWithTOPICITEM is called when exiting the identifierWithTOPICITEM production.
	ExitIdentifierWithTOPICITEM(c *IdentifierWithTOPICITEMContext)

	// ExitIdentifierWithQualifier is called when exiting the identifierWithQualifier production.
	ExitIdentifierWithQualifier(c *IdentifierWithQualifierContext)
}
