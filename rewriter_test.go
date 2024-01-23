package plsql_test

import (
	"testing"

	plsql "github.com/qsliu2017/plsql-rewriter"
)

func TestRewriteString(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected string
	}{
		{`create table t1 (id int, name varchar2(20));`, `create table t1 (id int, name varchar(20));`},
	} {
		rewritten := plsql.RewriteString(tc.input)
		if rewritten != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, rewritten)
		}
	}
}
