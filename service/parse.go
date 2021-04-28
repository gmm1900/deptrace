package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type VersionedModule struct {
	Module string
	Version string
}

func (mv *VersionedModule) ToString() string {
	return mv.Module + "@" + mv.Version
}

func parseVersionedModule(versionedModuleStr string) (*VersionedModule, error) {
	strs := strings.Split(versionedModuleStr, "@")
	if len(strs) == 1 {
		return &VersionedModule{
			Module: strs[0],
			Version: "",
		}, nil
	}
	if len(strs) == 2 {
		return &VersionedModule{
			Module:  strs[0],
			Version: strs[1],
		}, nil
	}
	return nil, fmt.Errorf("parsing versioned module %s: expect 2 segments separated by @, but got %d", versionedModuleStr, len(strs))
}

// Input correspond to one line in the `go mod graph` output, e.g.,
// line: github.com/prometheus/common@v0.2.0 gopkg.in/alecthomas/kingpin.v2@v2.2.6
// results in Dependant = github.com/prometheus/common@v0.2.0, Dependency = gopkg.in/alecthomas/kingpin.v2@v2.2.6
type Input struct {
	Dependant  VersionedModule
	Dependency VersionedModule
}

// NewInput creates an input from one line from the `go mod graph` output
func parseInput(rawInput string) (*Input, error) {
	strs := strings.Split(rawInput, " ")
	if len(strs) != 2 {
		return nil, fmt.Errorf("parse input %s: expect 2 versioned modules but got %d", rawInput, len(strs))
	}

	dependant, err := parseVersionedModule(strs[0])
	if err != nil {
		return nil, err
	}
	dependency, err := parseVersionedModule(strs[1])
	if err != nil {
		return nil, err
	}

	return &Input{
		Dependant: *dependant,
		Dependency: *dependency,
	}, nil
}

func ParseFile(pathToFile string) ([]Input, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lineErrorMsgs []string
	var inputs []Input
	lineNum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input, err := parseInput(line)
		if err != nil {
			lineErrorMsgs = append(lineErrorMsgs, fmt.Sprintf("line %d: %v", lineNum, err))
		} else {
			inputs = append(inputs, *input)
		}
		lineNum++
	}
	if len(lineErrorMsgs) > 0 {
		return nil, fmt.Errorf("file parsing error: %v", strings.Join(lineErrorMsgs, "\n"))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inputs, nil
}