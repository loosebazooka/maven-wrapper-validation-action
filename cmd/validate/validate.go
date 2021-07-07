package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/loosebazooka/mvn-validate/internal/jar"
	"github.com/loosebazooka/mvn-validate/internal/maven"
)

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
