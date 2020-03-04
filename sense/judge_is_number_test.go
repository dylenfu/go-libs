package sense

import (
	"regexp"
	"testing"
)

func isNumber(str string) bool {
	re := regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`)
	isNum := re.Match([]byte(str))
	return isNum
}

// go test -v github.com/dylenfu/go-libs/sense -run TestIsNumber
func TestIsNumber(t *testing.T) {
	t.Log(isNumber("2374892349273857192834023"))
	t.Log(isNumber("23223748923e"))
	t.Log(isNumber("23223748923.28342"))
	t.Log(isNumber("23223748923.283d42"))
}
