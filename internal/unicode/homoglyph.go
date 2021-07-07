/**
 * File originally from this repository but converted to golang.
 * https://github.com/codebox/homoglyph
 */

package unicode

import (
	"strings"
	"unicode/utf8"
)

var targets = []string{"maven-wrapper.jar", "maven_wrapper.jar", "mavenwrapper.jar"}
var sizes = map[int]bool{}

func init() {
	for _, w := range targets {
		sizes[utf8.RuneCountInString(w)] = true
	}
}

func LooksLikeMavenWrapperJar(val string) bool {
	if _, ok := sizes[utf8.RuneCountInString(val)]; !ok {
		return false
	}

	for _, t := range targets {
		if looksLike(val, t) {
			return true
		}
	}
	return false
}

func looksLike(val string, tgt string) bool {
	val = strings.TrimSpace(val)
	if utf8.RuneCountInString(val) != utf8.RuneCountInString(tgt) {
		return false
	}
	tr := []rune(tgt)
	vr := []rune(val)
	for i, v := range vr {
		if !strings.ContainsRune(homoglyphs[tr[i]], v) {
			return false
		}
	}
	return true
}
