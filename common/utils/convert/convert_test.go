package convert

import (
	"strconv"
	"testing"
)

func TestName(t *testing.T) {
	tem, err := strconv.ParseInt("", 10, 64)
	t.Log(err)
	t.Log(tem)
}
