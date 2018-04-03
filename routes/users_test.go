package routes

import (
	"testing"
)

func TestGetUsers (t *testing.T) {
	users := GetUsers()
	if users != 2 {
		t.Errorf("Get users failed. Expecte %d, got %d", 1, users)
	}
}
