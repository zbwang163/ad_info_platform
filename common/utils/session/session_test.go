package session

import (
	"testing"
)

func TestGetCoreUserIdFromSession(t *testing.T) {
	uid := GetCoreUserIdFromSession("")
	t.Log(uid)
}
