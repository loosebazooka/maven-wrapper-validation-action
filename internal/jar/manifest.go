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

import "fmt"

const (
	implementationTitle   = "Implementation-Title"
	implementationVersion = "Implementation-Version"
)

type Manifest struct {
	properties map[string]string
}

func (m *Manifest) ImplementationVersion() (string, error) {
	return m.property(implementationVersion)
}

func (m *Manifest) ImplementationTitle() (string, error) {
	return m.property(implementationTitle)
}

func (m *Manifest) property(key string) (string, error) {
	val, ok := m.properties[key]
	if !ok {
		return "", fmt.Errorf("%q not found in manifest", key)
	}
	return val, nil
}
