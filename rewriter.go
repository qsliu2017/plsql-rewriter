package plsql

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/qsliu2017/plsql-rewriter/parser"
)

func rewrite(is *antlr.InputStream) string {
	tokens := antlr.NewCommonTokenStream(
		parser.NewPlSqlLexer(is),
		antlr.TokenDefaultChannel,
	)
	r := &rewriter{
		antlr.NewTokenStreamRewriter(tokens),
		&parser.BasePlSqlParserListener{},
	}
	p := parser.NewPlSqlParser(tokens)
	p.AddParseListener(r)
	p.Sql_script()
	return r.GetTextDefault()
}

func RewriteString(origin string) string {
	return rewrite(antlr.NewInputStream(origin))
}

type rewriter struct {
	*antlr.TokenStreamRewriter
	*parser.BasePlSqlParserListener
}

func (r *rewriter) ExitColumn_definition(ctx *parser.Column_definitionContext) {
	dt := ctx.Datatype()
	if dt == nil {
		return
	}
	varchar2 := dt.Native_datatype_element().VARCHAR2()
	if varchar2 == nil {
		return
	}
	r.ReplaceTokenDefaultPos(varchar2.GetSymbol(), "varchar")
}
