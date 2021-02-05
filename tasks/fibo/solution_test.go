package fibo

import (
	"errors"
	"github.com/pr00f/courses/tasks"
	"log"
	"strconv"
	"testing"
)

func TestAll(t *testing.T) {
	f := func(in []string, out []string) error {
		n, err := strconv.Atoi(in[0])
		if err != nil {
			return err
		}

		equal, err := strconv.Atoi(out[0])
		if err != nil {
			return err
		}

		t1 := SolutionRecursion(n)
		t2 := SolutionIteration(n)
		t3 := SolutionGoldenRatio(n)
		t4 := SolutionMatrix(n)

		if t1 != equal {
			return errors.New("SolutionRecursion result incorrect")
		}

		if t2 != equal {
			return errors.New("SolutionRecursion result incorrect")
		}

		if t3 != equal {
			return errors.New("SolutionGoldenRatio result incorrect")
		}

		if t4 != equal {
			return errors.New("SolutionMatrix result incorrect")
		}

		log.Printf("SolutionRecursion result : %v", t1)
		log.Printf("SolutionIteration result : %v", t2)
		log.Printf("SolutionGoldenRatio result : %v", t3)
		log.Printf("SolutionMatrix result : %v", t4)
		log.Println("---------------------------------------")

		return nil
	}

	if err := tasks.RunTests("fibo", 8, f); err != nil {
		t.Error(err)
	} else {
		t.Log("OK")
	}
}
