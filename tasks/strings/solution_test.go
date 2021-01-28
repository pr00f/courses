package strings

import (
	"errors"
	"github.com/pr00f/courses/tasks"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	f := func(in []string, out []string) error {
		if strconv.Itoa(Solution(in[0])) != out[0] {
			return errors.New("result incorrect")
		}

		return nil
	}

	if err := tasks.RunTests("strings", 5, f); err != nil {
		t.Error(err)
	} else {
		t.Log("OK")
	}
}
