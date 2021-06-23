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

import "strings"

type Repository struct {
	url string
}

var MavenCentral = Repository{"https://repo1.maven.org/maven2"}

func (r *Repository) Artifact(groupId, artifactId, version string) Artifact {
	gp := strings.ReplaceAll(groupId, ".", "/")
	ar := strings.Join([]string{r.url, gp, artifactId, version}, "/")

	fn := artifactId + "-" + version + ".jar"

	art := Artifact{}
	art.root = ar
	art.fileName = fn
	return art
}
