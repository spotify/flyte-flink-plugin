// Copyright 2021 Spotify AB.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flink

import (
	"reflect"
	"sort"
	"testing"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/utils"
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
		"--string=foobar",
		"--boolean_true",
		"--collection=1",
		"--collection=2",
	}
	assert.Assert(t, len(args) == len(expected))

	sort.Strings(args)
	sort.Strings(expected)
	assert.Assert(t, reflect.DeepEqual(args, expected))
}

func TestFlinkClusterNameRegex(t *testing.T) {
	t.Run("Valid cluster name", func(t *testing.T) {
		validNames := []string{
			"valid-name",
			"valid-name.have.periods",
			"2valid-name.starts.with.a.number",
			"valid-name.ends.with.a.number4",
			"valid-name.with.multiple-dashes",
		}

		for _, valid := range validNames {
			err := validate(valid, regexpFlinkClusterName)
			assert.NilError(t, err)
		}
	})

	t.Run("Invalid cluster name", func(t *testing.T) {
		invalidNames := []string{
			".invalid-name",
			"invalid-name.",
			"_invalid-name",
			"invalid-name_",
			"invalid_name",
			"invalid name",
		}

		for _, invalid := range invalidNames {
			err := validate(invalid, regexpFlinkClusterName)
			assert.ErrorContains(t, err, "Validation error: ")
		}
	})
}
