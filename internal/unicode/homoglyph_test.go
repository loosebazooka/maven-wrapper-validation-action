package unicode

import (
	"fmt"
	"testing"
)

func TestLooksLikeMavenWrapperJar(t *testing.T) {
	passingCases := []string{
		"maven-wrapper.jar",
		"mavenwrapper.jar",
		"maven_wrapper.jar",
		"maven﹏wrapper.jar", // non standard unicode _ (U+FE4F)
		"maven-wrappеr.jar", // non english unicode e (U+0435) in maven-wrapp"e"r
	}
	failingCases := []string{
		"new-haven-wrapper.jar",
		"maaven-wrapper.jar", // stuff like this needs to be manually handled by user
	}

	for i, v := range passingCases {
		t.Run(fmt.Sprintf("passingCases[%v]:%v", i, v), func(t *testing.T) {
			if LooksLikeMavenWrapperJar(v) != true {
				t.Fatal()
			}
		})
	}

	for i, v := range failingCases {
		t.Run(fmt.Sprintf("failingCases[%v]:%v", i, v), func(t *testing.T) {
			if LooksLikeMavenWrapperJar(v) != false {
				t.Fatal()
			}
		})
	}
}
