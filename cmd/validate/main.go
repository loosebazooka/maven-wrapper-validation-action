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

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/loosebazooka/mvn-validate/internal/file"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %q <search root | path to maven-wrapper.jar>", os.Args[0])
	}

	path := os.Args[1]

	info, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if info.IsDir() { // search
		wrappers, err := file.FindMavenWrappers(path)
		if err != nil {
			log.Fatal("Failed during search for maven-wrapper.jar: ", err)
		}
		if len(wrappers) == 0 {
			log.Fatal("Failed to find any maven-wrapper.jar files")
		} // do search
		for _, w := range wrappers {
			err := validate(ctx, w)
			if err != nil {
				log.Fatal("Failed to validate Jar: ", err)
			}
		}
	} else { // single file
		if filepath.Base(path) != "maven-wrapper.jar" {
			log.Fatal("Specified file is not named maven-wrapper.jar: ", path)
		}
		err := validate(ctx, path)
		if err != nil {
			log.Fatal("Failed to validate Jar: ", err)
		}
	}

	// success
	os.Exit(0)
}
