package flink

import (
	"reflect"
	"sort"
	"testing"

	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/utils"
	"gotest.tools/assert"
)

func TestLiteralMapToArgs(t *testing.T) {
	integer, _ := utils.MakeLiteral(1)
	float, _ := utils.MakeLiteral(1.01)
	str, _ := utils.MakeLiteral("foobar")
	boolTrue, _ := utils.MakeLiteral(true)
	boolFalse, _ := utils.MakeLiteral(false)
	coll, _ := utils.MakeLiteral([]interface{}{1, 2})

	literals := map[string]*core.Literal{
		"string":        str,
		"integer":       integer,
		"float":         float,
		"boolean_true":  boolTrue,
		"boolean_false": boolFalse,
		"collection":    coll,
	}

	args, err := literalMapToFlinkJobArgs(literals)
	assert.NilError(t, err)

	expected := []string{
		"--integer=1",
		"--float=1.010000",
		"--string=\"foobar\"",
		"--boolean_true",
		"--collection=1",
		"--collection=2",
	}
	assert.Assert(t, len(args) == len(expected))

	sort.Strings(args)
	sort.Strings(expected)
	assert.Assert(t, reflect.DeepEqual(args, expected))
}
