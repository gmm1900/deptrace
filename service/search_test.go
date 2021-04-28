package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchPath(t *testing.T) {
	tree := Analyze([]Input{
		{
			Dependant:  VersionedModule{"m1", "v0.1"},
			Dependency: VersionedModule{"m2", "v0.2"},
		}, {
			Dependant:  VersionedModule{"m1", "v0.1"},
			Dependency: VersionedModule{"m3", "v0.3"},
		}, {
			Dependant:  VersionedModule{"m2", "v0.2"},
			Dependency: VersionedModule{"m3", "v0.3"},
		}, {
			Dependant:  VersionedModule{"m2", "v0.2"},
			Dependency: VersionedModule{"m4", "v0.4"},
		}, {
			Dependant:  VersionedModule{"m4", "v0.4"},
			Dependency: VersionedModule{"m3", "v0.3.1"},
		}, {
			Dependant:  VersionedModule{"m3", "v0.3"},
			Dependency: VersionedModule{"m2", "v0.2"}, // circular dependency
		},
	})

	// exact version search
	dependencyPaths := tree.SearchPath("m3@v0.3")
	assert.Equal(t, []DependencyPath{
		{ "m1@v0.1", "m3@v0.3" }, // path 1
		{ "m1@v0.1", "m2@v0.2", "m3@v0.3" }, // path 2
	}, dependencyPaths)

	// arbitrary version search
	dependencyPaths = tree.SearchPath("m3")
	assert.Equal(t, []DependencyPath{
		{ "m1@v0.1", "m3@v0.3" },
		{ "m1@v0.1", "m2@v0.2", "m3@v0.3" },
		{ "m1@v0.1", "m2@v0.2", "m4@v0.4", "m3@v0.3.1" },
		{ "m1@v0.1", "m3@v0.3", "m2@v0.2", "m4@v0.4", "m3@v0.3.1" },
	}, dependencyPaths)
}