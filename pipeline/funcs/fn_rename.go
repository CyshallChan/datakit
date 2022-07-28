// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import (
	"fmt"
	"reflect"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/parser"
)

func RenameChecking(ng *parser.EngineData, node parser.Node) error {
	funcExpr := fexpr(node)
	if len(funcExpr.Param) != 2 {
		return fmt.Errorf("func %s expected 2 args", funcExpr.Name)
	}

	switch funcExpr.Param[0].(type) {
	case *parser.AttrExpr, *parser.StringLiteral, *parser.Identifier:
	default:
		return fmt.Errorf("expect string or AttrExpr, got `%s'",
			reflect.TypeOf(funcExpr.Param[0]).String())
	}

	switch funcExpr.Param[1].(type) {
	case *parser.AttrExpr, *parser.Identifier:
	default:
		return fmt.Errorf("param key expect Identifier or AttrExpr, got `%s'",
			reflect.TypeOf(funcExpr.Param[1]).String())
	}
	return nil
}

func Rename(ng *parser.EngineData, node parser.Node) interface{} {
	funcExpr := fexpr(node)
	if len(funcExpr.Param) != 2 {
		return fmt.Errorf("func %s expected 2 args", funcExpr.Name)
	}

	var from, to parser.Node

	switch v := funcExpr.Param[0].(type) {
	case *parser.AttrExpr, *parser.StringLiteral, *parser.Identifier:
		to = v
	default:
		return fmt.Errorf("expect string or AttrExpr, got `%s'",
			reflect.TypeOf(funcExpr.Param[0]).String())
	}

	switch v := funcExpr.Param[1].(type) {
	case *parser.AttrExpr, *parser.Identifier:
		from = v
	default:
		return fmt.Errorf("param key expect Identifier or AttrExpr, got `%s'",
			reflect.TypeOf(funcExpr.Param[1]).String())
	}

	v, err := ng.GetContent(from)
	if err != nil {
		l.Debug(err)
		return nil
	}

	if err := ng.SetContent(to, v); err != nil {
		l.Debug(err)
		return nil
	}
	_ = ng.DeleteContent(from.String())

	return nil
}
