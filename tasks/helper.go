package tasks

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const rootFolderPrefix = "../../data/"

func RunTests(folderName string, numberOfTests int, testFunc func([]string, []string) error) error {
	folderName = rootFolderPrefix + folderName

	for i := 0; i < numberOfTests; i++ {
		f := newTestFileDefinition(folderName, i)
		if !f.isInOutExists() {
			return errors.New("no in files for test folder - " + folderName)
		}

		in, err := f.readInPerLine()
		if err != nil {
			return err
		}

		out, err := f.readOutPerLine()
		if err != nil {
			return err
		}

		if err := testFunc(in, out); err != nil {
			return fmt.Errorf("test failed on number %d, original error: %w", i, err)
		}
	}

	return nil
}

type testFileDefinition struct {
	in  string
	out string
}

func newTestFileDefinition(folderName string, i int) testFileDefinition {
	return testFileDefinition{
		in:  fmt.Sprintf("%s/test.%d.in", folderName, i),
		out: fmt.Sprintf("%s/test.%d.out", folderName, i),
	}
}

func (t testFileDefinition) isInOutExists() bool {
	_, err := os.Stat(t.in)
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(t.out)
	return !os.IsNotExist(err)
}

func (t testFileDefinition) readInPerLine() ([]string, error) {
	return t.readFile(t.in)
}

func (t testFileDefinition) readOutPerLine() ([]string, error) {
	return t.readFile(t.out)
}

func (t testFileDefinition) readFile(dest string) ([]string, error) {
	file, err := os.Open(dest)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}
