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

import "testing"

func TestRepositoryArtifactGeneration(t *testing.T) {
	a := MavenCentral.Artifact("com.google.cloud.tools", "jib-core", "0.19.0")

	if got, want := a.root, "https://repo1.maven.org/maven2/com/google/cloud/tools/jib-core/0.19.0"; got != want {
		t.Errorf("root = %q, want %q", got, want)
	}

	if got, want := a.fileName, "jib-core-0.19.0.jar"; got != want {
		t.Errorf("filename = %q, want %q", got, want)
	}
}
