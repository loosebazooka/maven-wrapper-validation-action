package file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/loosebazooka/mvn-validate/internal/unicode"
)

const (
	mavenWrapper = "maven-wrapper.jar"
)

func FindMavenWrappers(root string) ([]string, error) {
	var targets []string
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == mavenWrapper {
				targets = append(targets, path)
			} else if unicode.LooksLikeMavenWrapperJar(info.Name()) {
				return fmt.Errorf("Potentially fraudulent maven wrapper at: %q", path)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return targets, nil
}
