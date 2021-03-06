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
// limitations under the License

package flink

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-version"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
)

func Validate(job *flinkIdl.FlinkJob) error {
	err := job.Validate()
	if err != nil {
		return err
	}

	jarFiles := len(job.GetJarFiles()) + len(job.GetJflyte().GetArtifacts())
	if jarFiles == 0 {
		return fmt.Errorf("no artifacts provided")
	}

	if len(job.GetFlinkVersion()) != 0 {
		if _, err = version.NewVersion(job.GetFlinkVersion()); err != nil {
			return err
		}
	}

	return nil
}

func ValidateRegEx(value string, r *regexp.Regexp) error {
	if r.MatchString(value) {
		return nil
	}
	return fmt.Errorf("validation error: %v doesn't match with the given regex expr: %v", value, r.String())
}
