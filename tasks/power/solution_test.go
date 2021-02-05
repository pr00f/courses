package power

import (
	"github.com/pr00f/courses/tasks"
	"log"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	f := func(in []string, out []string) error {
		a, err := strconv.ParseFloat(in[0], 64)
		if err != nil {
			return err
		}

		n, err := strconv.Atoi(in[1])
		if err != nil {
			return err
		}

		log.Printf("SolutionSimple result : %v", SolutionSimple(a, n))
		log.Printf("SolutionPowerOfTwoWithMultiplication result : %v", SolutionPowerOfTwoWithMultiplication(a, n))
		log.Printf("SolutionBinaryExpansionOfExponent result : %v", SolutionBinaryExpansionOfExponent(a, n))
		log.Println("---------------------------------------")

		return nil
	}

	if err := tasks.RunTests("power", 10, f); err != nil {
		t.Error(err)
	} else {
		t.Log("OK")
	}
}
