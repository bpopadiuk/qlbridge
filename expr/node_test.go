package expr_test

import (
	"encoding/json"
	"testing"

	u "github.com/araddon/gou"
	"github.com/bmizerany/assert"
	"github.com/gogo/protobuf/proto"

	"github.com/araddon/qlbridge/expr"
)

var pbTests = []string{
	`eq(event,"stuff") OR ge(party, 1)`,
	`"Portland" IN ("ohio")`,
	`"xyz" BETWEEN todate("1/1/2015") AND 50`,
	`name == "bob"`,
	`name = 'bob'`,
	`AND ( EXISTS x, EXISTS y )`,
	`AND ( EXISTS x, INCLUDE ref_name )`,
}

func TestNodePb(t *testing.T) {
	t.Parallel()
	for _, exprText := range pbTests {
		exp, err := expr.ParseExpression(exprText)
		assert.Equalf(t, err, nil, "Should not error parse expr but got ", err, "for ", exprText)
		pb := exp.NodePb()
		assert.Tf(t, pb != nil, "was nil PB: %#v", exp)
		pbBytes, err := proto.Marshal(pb)
		assert.Tf(t, err == nil, "Should not error on proto.Marshal but got [%v] for %s pb:%#v", err, exprText, pb)
		n2, err := expr.NodeFromPb(pbBytes)
		assert.T(t, err == nil, "Should not error from pb but got ", err, "for ", exprText)
		assert.Tf(t, exp.Equal(n2), "Equal?  %v  %v", exp, n2)
		u.Infof("pre/post: \n\t%s\n\t%s", exp, n2)
	}
}

func TestExprRoundTrip(t *testing.T) {
	t.Parallel()
	for _, et := range exprTests {
		exp, err := expr.ParseExpression(et.qlText)
		if et.ok {
			assert.Equalf(t, err, nil, "Should not error parse expr but got %v for %s", err, et.qlText)
			by, err := json.MarshalIndent(exp.Expr(), "", "  ")
			assert.Equal(t, err, nil)
			u.Debugf("%s", string(by))
			en := &expr.Expr{}
			err = json.Unmarshal(by, en)
			assert.Equal(t, err, nil)
			_, err = expr.NodeFromExpr(en)
			assert.Equal(t, err, nil)
			// TODO: Fixme
			//assert.Tf(t, nn.Equal(exp), "%s  doesn't match %s", et.qlText, nn.String())
		} else {
			assert.NotEqual(t, nil, err)
		}

	}
}

func TestNodeJson(t *testing.T) {
	t.Parallel()
	for _, exprText := range pbTests {
		exp, err := expr.ParseExpression(exprText)
		assert.Equalf(t, err, nil, "Should not error parse expr but got ", err, "for ", exprText)
		by, err := json.MarshalIndent(exp.Expr(), "", "  ")
		assert.Equal(t, err, nil)
		u.Debugf("%s", string(by))
	}
}

var _ = u.EMPTY
