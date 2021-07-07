//
// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build e2e

package e2e

import (
	"os/exec"
	"strings"
	"testing"
)

// It's quite possible these integration tests fail due to maven
// central blocking requests through rate limiting mechanisms

func runCli(t *testing.T, arg ...string) (string, error) {
	t.Helper()
	cmd := "../../validate"
	c := exec.Command(cmd, arg...)
	out, err := c.CombinedOutput()
	return string(out), err
}

func TestGoodArtifact(t *testing.T) {
	out, err := runCli(t, "testdata/good/maven-wrapper.jar")
	if err != nil {
		t.Log(out)
		t.Fatal("Unexpected failure on good artifact", err)
	}
}

func TestGoodSearchRoot(t *testing.T) {
	out, err := runCli(t, "testdata/good")
	if err != nil {
		t.Log(out)
		t.Fatal("Unexpected failure on good artifact", err)
	}
}

func TestBadArtifact(t *testing.T) {
	type TestCase struct {
		path   string
		errMsg string
	}

	testcases := []TestCase{{
		path:   "testdata/bad-sha/maven-wrapper.jar",
		errMsg: "corrupt wrapper detected",
	}, {
		path:   "testdata/unknown-version/maven-wrapper.jar",
		errMsg: "could not retrieve remote artifact -- status code:404",
	}, {
		path:   "testdata/homoglyph/maven-wrappеr.jar",  // non english unicode e (U+0435) in maven-wrapp"e"r
		errMsg: "Specified file is not named maven-wrapper.jar",
	}}

	for _, tc := range testcases {
		t.Run(tc.path, func(t *testing.T) {
			out, err := runCli(t, tc.path)
			if err == nil {
				t.Log(out)
				t.Fatal("Unexpected error free execution")
			} else if !strings.Contains(out, tc.errMsg) {
				t.Log(out)
				t.Fatal("Unexpected error")
			}
		})
	}
}

func TestBadSearchRoots(t *testing.T) {
	type TestCase struct {
		path   string
		errMsg string
	}

	testcases := []TestCase{{
		path:   "testdata", // contains all test wrappers, good and bad
		errMsg: "Failed ",
	}, {
		path:   "testdata/none",
		errMsg: "Failed to find any maven-wrapper.jar files",
	}, {
		path:   "testdata/bad-sha",
		errMsg: "corrupt wrapper detected",
	}, {
		path:   "testdata/unknown-version",
		errMsg: "could not retrieve remote artifact -- status code:404",
	}, {
		path:   "testdata/homoglyph",
		errMsg: "Potentially fraudulent maven wrapper at:",
	}}
	for _, tc := range testcases {
		t.Run(tc.path, func(t *testing.T) {
			out, err := runCli(t, tc.path)
			if err == nil {
				t.Log(out)
				t.Fatal("Unexpected error free execution")
			} else if !strings.Contains(out, tc.errMsg) {
				t.Log(out)
				t.Fatal("Unexpected error")
			}
		})
	}
}
