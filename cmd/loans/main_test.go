package main_test

import (
	"os/exec"
	"strings"
	"testing"
)

type testCase struct {
	description  string
	amount       string
	term         string
	errorMessage string
}

func TestInvalidInputReturnsError(t *testing.T) {
	cases := []testCase{
		{"missing loan", "", "12", "invalid value \"\" for flag -amount"},
		{"invalid loan value", "abc", "12", "invalid value \"abc\" for flag -amount"},
		{"amount below minimum", "500.00", "12", "below lower boundary"},
		{"negative amount", "-1000.00", "12", "Amount must be above 0"},
		{"missing term", "1000.00", "", "invalid value \"\" for flag -term"},
		{"unsupported term", "1000.00", "36", "term 36 not supported"},
	}

	for _, test := range cases {
		test := test

		t.Run(test.description, func(t *testing.T) {
			cmd := exec.Command(
				"go", "run", "main.go",
				"-amount="+test.amount,
				"-term="+test.term,
			)

			output, err := cmd.CombinedOutput()

			exitErr, ok := err.(*exec.ExitError)
			if !ok {
				t.Fatalf("expected exit error, got: %v", err)
			}

			if exitErr.ExitCode() != 1 {
				t.Errorf("expected exit code 1, got %d", exitErr.ExitCode())
			}

			if !strings.Contains(string(output), test.errorMessage) {
				t.Errorf("expected error %q in output %q", test.errorMessage, string(output))
			}
		})
	}
}
