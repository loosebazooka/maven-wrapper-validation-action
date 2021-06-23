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

package jar

import (
	"bufio"
	"os"
	"testing"
)

func TestParseManifest_Good(t *testing.T) {
	f, err := os.Open("testdata/good.properties")
	if err != nil {
		t.Fatal(err)
	}
	r := bufio.NewReader(f)

	m, err := parseManifest(r)
	if err != nil {
		t.Fatal(err)
	}

	for k, want := range map[string]string{"property1": "goose", "property2": "goose moose", "property3": "goose:moose"} {
		if got := m.properties[k]; got != want {
			t.Errorf("m[%q] = %q, want %q", k, got, want)
		}
	}
}

func TestParseManifest_Bad(t *testing.T) {
	f, err := os.Open("testdata/bad.properties")
	if err != nil {
		t.Fatal(err)
	}
	r := bufio.NewReader(f)

	_, err = parseManifest(r)
	if err == nil {
		t.Fatalf("want err, got nil")
	}
}
