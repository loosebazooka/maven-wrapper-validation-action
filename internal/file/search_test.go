package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSearchSuccessfully(t *testing.T) {
	testRoot := t.TempDir()
	mw := filepath.Join(testRoot, "maven-wrapper.jar")
	err := ioutil.WriteFile(mw, []byte{}, 0644)
	if err != nil {
		t.Fatal(err)
	}

	subDir, err := ioutil.TempDir(testRoot, "sub")
	if err != nil {
		t.Fatal(err)
	}
	mw2 := filepath.Join(subDir, "maven-wrapper.jar")
	err = ioutil.WriteFile(mw2, []byte{}, 0644)

	defer os.Remove(mw)
	defer os.Remove(mw2)

	if err != nil {
		t.Fatal(err)
	}

	want := []string{mw, mw2}

	got, err := FindMavenWrappers(testRoot)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got:%v != want:%v", got, want)
	}
}

func TestFindFraudulentWrapper(t *testing.T) {
	testRoot := t.TempDir()
	mw := filepath.Join(testRoot, "maven_wrapper.jar")
	err := ioutil.WriteFile(mw, []byte{}, 0644)
	if err != nil {
		t.Fatal(err)
	}

	got, err := FindMavenWrappers(testRoot)
	if err != nil && err.Error() != fmt.Sprintf("Potentially fraudulent maven wrapper at: %q", mw) {
		t.Fatal(err)
	} else if err == nil {
		t.Fatalf("want error, got %v", got)
	}
}

func TestFindNothing(t *testing.T) {
	testRoot := t.TempDir()

	got, err := FindMavenWrappers(testRoot)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 0 {
		t.Fatalf("got:%v != want:[]", got)
	}
}
