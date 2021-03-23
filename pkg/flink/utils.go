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
	"fmt"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
)

func literalMapToFlinkJobArgs(literals map[string]*core.Literal) ([]string, error) {
	args := []string{}
	for key, literal := range literals {
		ls, err := literalToFlinkJobArg(key, literal)
		if err != nil {
			return nil, err
		}

		args = append(args, ls...)
	}
	return args, nil
}

func literalToFlinkJobArg(key string, literal *core.Literal) ([]string, error) {
	switch l := literal.GetValue().(type) {
	case *core.Literal_Scalar:
		arg, err := scalarToFlinkJobArg(key, l.Scalar)
		if err != nil {
			return nil, err
		}

		// empty arg on boolean false value
		if len(arg) == 0 {
			return []string{}, nil
		}

		return []string{arg}, nil
	case *core.Literal_Collection:
		literals := l.Collection.GetLiterals()
		args := []string{}

		for _, l := range literals {
			strArgs, err := literalToFlinkJobArg(key, l)
			if err != nil {
				return nil, err
			}

			args = append(args, strArgs...)
		}

		return args, nil
	default:
		return nil, fmt.Errorf("not supported type: %s", l)
	}
}

func scalarToFlinkJobArg(arg string, scalar *core.Scalar) (string, error) {
	switch s := scalar.GetValue().(type) {
	case *core.Scalar_Primitive:
		return primitiveToFlinkJobArg(arg, s.Primitive)
	default:
		return "", fmt.Errorf("not supported type: %s", s)
	}
}

func primitiveToFlinkJobArg(arg string, primitive *core.Primitive) (string, error) {
	switch p := primitive.GetValue().(type) {
	case *core.Primitive_Integer:
		return fmt.Sprintf("--%s=%d", arg, p.Integer), nil
	case *core.Primitive_FloatValue:
		return fmt.Sprintf("--%s=%f", arg, p.FloatValue), nil
	case *core.Primitive_Boolean:
		if p.Boolean {
			return fmt.Sprintf("--%s", arg), nil
		}

		return "", nil
	case *core.Primitive_StringValue:
		return fmt.Sprintf("--%s=%s", arg, p.StringValue), nil
	case *core.Primitive_Datetime:
		return fmt.Sprintf("--%s=%s", arg, p.Datetime.String()), nil
	case *core.Primitive_Duration:
		return fmt.Sprintf("--%s=%s", arg, p.Duration.String()), nil
	default:
		return "", fmt.Errorf("not supported type: %s", p)
	}
}
