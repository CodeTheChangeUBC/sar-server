package db

import "testing"

func TestValidate(t *testing.T) {
	p := Point{"34.5231", "-43.23421"}
	if !p.Validate() {
		t.Fatalf("Failed to validate Point: %v", p)
	}
}
