package tickets

import (
	"errors"
	"github.com/pr00f/courses/tasks"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {

	f := func(in []string, out []string) error {

		input, _ := strconv.Atoi(in[0])
		if strconv.Itoa(Solution(input)) != out[0] {
			return errors.New("result incorrect")
		}

		return nil
	}

	if err := tasks.RunTests("tickets", 10, f); err != nil {
		t.Error(err)
	} else {
		t.Log("OK")
	}
}
