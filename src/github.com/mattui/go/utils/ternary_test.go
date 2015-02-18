package utils

import "testing"

func TestTernary(t *testing.T) {
	if Ternary(true, "true", "false") != "true" {
		t.Error("Ternary(true, \"true\", \"false\") != \"true\"")
	}

	if Ternary(false, "true", "false") != "false" {
		t.Error("Ternary(false, \"true\", \"false\") != \"false\"")
	}
}
