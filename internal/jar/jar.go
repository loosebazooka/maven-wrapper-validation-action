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
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	manifestPath = "META-INF/MANIFEST.MF"
)

func ExtractManifest(jar string) (*Manifest, error) {
	zipfile, err := zip.OpenReader(jar)
	if err != nil {
		return nil, err
	}
	defer zipfile.Close()

	for _, file := range zipfile.File {
		// check if the file matches the name for application portfolio xml
		if file.Name == manifestPath {
			reader, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer reader.Close()

			return parseManifest(reader)
		}
	}
	return nil, fmt.Errorf("no manifest found in %q at %q", jar, manifestPath)
}

func parseManifest(in io.Reader) (*Manifest, error) {
	manifest := Manifest{}
	manifestMap := make(map[string]string)

	sc := bufio.NewScanner(in)

	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			continue
		}
		i := strings.Index(line, ":")
		if i == -1 {
			return nil, fmt.Errorf("unparseable properties line '%q' in %q", line, manifestPath)
		}
		key := strings.TrimSpace(line[0:i])
		value := strings.TrimSpace(line[i+1:])
		manifestMap[key] = value
	}
	manifest.properties = manifestMap
	return &manifest, nil
}
