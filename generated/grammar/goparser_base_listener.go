// Code generated from grammar/GoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GoParser
import "github.com/antlr4-go/antlr/v4"

// BaseGoParserListener is a complete listener for a parse tree produced by GoParser.
type BaseGoParserListener struct{}

var _ GoParserListener = &BaseGoParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGoParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGoParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseGoParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseGoParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseGoParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseGoParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseGoParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseGoParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseGoParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseGoParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGoParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGoParserListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseGoParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseGoParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseGoParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseGoParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseGoParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseGoParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseGoParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseGoParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseGoParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseGoParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseGoParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseGoParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoParserListener) ExitExpression(ctx *ExpressionContext) {}
