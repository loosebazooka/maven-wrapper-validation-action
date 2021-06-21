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

package maven

import (
	"context"
	"testing"
)

func TestHashes(t *testing.T) {
	type TestCase struct {
		testName      string
		hashFunc      func(ctx context.Context) (string, error)
		expectedValue string
	}

	a := MavenCentral.Artifact("com.google.guava", "guava", "30.0-jre")

	testCases := []TestCase{{
		testName:      "sha1",
		hashFunc:      a.SHA1,
		expectedValue: "8ddbc8769f73309fe09b54c5951163f10b0d89fa",
	}, {
		testName:      "md5",
		hashFunc:      a.MD5,
		expectedValue: "a7379037e654274e3544f2488a29b31f",
	}, {
		testName:      "sha256",
		hashFunc:      a.CalculateSHA256,
		expectedValue: "56b292df9ec29d102820c1fd7dd581cd749d5c416c7b3aeac008dbda3b984cc2",
	}}

	for _, tc := range testCases {
		if hash, err := tc.hashFunc(context.Background()); err != nil {
			t.Fatal(err)
		} else if hash != tc.expectedValue {
			t.Errorf("%q = %q, want %q'", tc.testName, hash, tc.expectedValue)
		}
	}
}
