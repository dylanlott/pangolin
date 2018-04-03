package main

import (
	"testing"
	"github.com/dylanlott/pangolin/routes"
)

func TestGet(t *testing.T) {
	get := routes.GetUsers()
	if get != 1 {
		t.Errorf("Get did not return correctly, expeected %d, got %d", 1, get)
	}
}
