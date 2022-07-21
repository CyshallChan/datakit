// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import (
	"fmt"
	"reflect"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/parser"
)

func ParseDurationChecking(ng *parser.EngineData, node parser.Node) error {
	funcExpr := fexpr(node)
	if len(funcExpr.Param) != 1 {
		return fmt.Errorf("func %s expects 1 arg", funcExpr.Name)
	}
	switch funcExpr.Param[0].(type) {
	case *parser.Identifier, *parser.AttrExpr:
	default:
		err := fmt.Errorf("param expects Identifier, got `%+#v', type `%s'",
			funcExpr.Param[0], reflect.TypeOf(funcExpr.Param[0]).String())
		return err
	}
	return nil
}

func ParseDuration(ng *parser.EngineData, node parser.Node) interface{} {
	funcExpr := fexpr(node)
	if len(funcExpr.Param) != 1 {
		l.Debugf("parse_duration(): invalid param")

		return fmt.Errorf("func %s expects 1 arg", funcExpr.Name)
	}

	var key parser.Node
	switch v := funcExpr.Param[0].(type) {
	case *parser.Identifier, *parser.AttrExpr:
		key = v
	default:
		err := fmt.Errorf("param expects Identifier, got `%+#v', type `%s'",
			funcExpr.Param[0], reflect.TypeOf(funcExpr.Param[0]).String())

		l.Debugf("parse_duration(): %s", err)

		return err
	}

	cont, err := ng.GetContent(key)
	if err != nil {
		l.Debug(err)
		return nil
	}

	duStr, ok := cont.(string)
	if !ok {
		return fmt.Errorf("parse_duration() expect string arg")
	}

	l.Debugf("parse duration %s", duStr)
	du, err := time.ParseDuration(duStr)
	if err != nil {
		l.Debug(err)
		return nil
	}

	if err := ng.SetContent(key, int64(du)); err != nil {
		l.Debug(err)
		return nil
	}
	return nil
}
