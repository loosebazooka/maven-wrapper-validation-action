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

import "testing"

func TestImplementationVersion(t *testing.T) {
	type TestCase struct {
		testName      string
		properties    map[string]string
		expectedValue string
		shouldErr     bool
	}
	testCases := []TestCase{{
		testName:      "found",
		properties:    map[string]string{"Implementation-Version": "10"},
		shouldErr:     false,
		expectedValue: "10",
	}, {
		testName:   "not found",
		properties: map[string]string{"Other": "10"},
		shouldErr:  true,
	}}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			manifest := &Manifest{
				properties: tc.properties,
			}
			val, err := manifest.ImplementationVersion()
			if (err != nil) != tc.shouldErr {
				t.Fatalf("error: %v, should-error: %v", err, tc.shouldErr)
			} else if err == nil && val != tc.expectedValue {
				t.Fatalf("value was: %v, expected: %v", val, tc.expectedValue)
			}
		})
	}

}

func TestImplementationTitle(t *testing.T) {
	type TestCase struct {
		testName      string
		properties    map[string]string
		expectedValue string
		shouldErr     bool
	}
	testCases := []TestCase{{
		testName:      "found",
		properties:    map[string]string{"Implementation-Title": "goose"},
		shouldErr:     false,
		expectedValue: "goose",
	}, {
		testName:   "not found",
		properties: map[string]string{"Other": "10"},
		shouldErr:  true,
	}}

	for _, tc := range testCases {
		manifest := &Manifest{
			properties: tc.properties,
		}
		val, err := manifest.ImplementationTitle()
		if (err != nil) != tc.shouldErr {
			t.Fatalf("(%v) Unexpected state -- error: %v, should-error: %v", tc.testName, err != nil, tc.shouldErr)
		} else if err == nil && val != tc.expectedValue {
			t.Fatalf("(%v) Unexpected value was: %v, expected: %v", tc.testName, val, tc.expectedValue)
		}
	}
}
