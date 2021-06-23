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
	out, err := runCli(t, "testdata/maven-wrapper.jar")
	if err != nil {
		t.Log(out)
		t.Fatal("Unexpected failure on good artifact", err)
	}
}

func TestBadSha(t *testing.T) {
	out, err := runCli(t, "testdata/maven-wrapper-bad-sha.jar")
	if err == nil {
		t.Log(out)
		t.Fatal("Unexpected pass on bad artifact")
	} else if !strings.Contains(out, "corrupt wrapper detected") {
		t.Log(out)
		t.Fatal("Unknown failure type")
	}
}

func TestUnknownVersion(t *testing.T) {
	out, err := runCli(t, "testdata/maven-wrapper-unknown-version.jar")
	if err == nil {
		t.Log(out)
		t.Fatal("Unexpected pass on bad artifact")
	} else if !strings.Contains(out, "could not retrieve remote artifact -- status code:404") {
		t.Log(out)
		t.Fatal("Unknown failure type")
	}
}
