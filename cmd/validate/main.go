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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/loosebazooka/mvn-validate/internal/jar"
	"github.com/loosebazooka/mvn-validate/internal/maven"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %q <path to maven-wrapper.jar>", os.Args[0])
	}

	jarFile := os.Args[1]

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := validate(ctx, jarFile)
	if err != nil {
		log.Fatal("Failed to validate Jar: ", err)
	}

	// success
	os.Exit(0)
}

func validate(ctx context.Context, jarFile string) error {
	file, err := os.Open(jarFile)
	if err != nil {
		return err
	}
	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return err
	}

	sb := h.Sum(make([]byte, 0, h.Size()))
	localSha := hex.EncodeToString(sb)

	manifest, err := jar.ExtractManifest(jarFile)
	if err != nil {
		return err
	}

	version, err := manifest.ImplementationVersion()
	if err != nil {
		return err
	}

	artifact := maven.MavenCentral.Artifact("io.takari", "maven-wrapper", version)

	remoteSha, err := artifact.CalculateSHA256(ctx)
	if err != nil {
		return err
	}

	if localSha != remoteSha {
		return &errMismatch{localSha, remoteSha}
	}
	return nil
}

type errMismatch struct {
	local  string
	remote string
}

func (e *errMismatch) Error() string {
	return fmt.Sprintf("corrupt wrapper detected\nlocal sha256: %s\nremote sha256: %s", e.local, e.remote)
}
