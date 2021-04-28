package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractTopPaths(t *testing.T) {
	givenDependencyPaths := []DependencyPath{
		{"m1@v1", "m2@v2", "m3@v3", "m4@v4"},
		{"m1@v1", "m2@v2", "m3@v3", "m5@v5"},
		{"m1@v1", "m2@v2", "m5@v5", "m6@v6"},
		{"m1@v1", "m2@v2"},
		{"m1@v1", "m3@v3"},
	}

	result := ExtractTopPaths(givenDependencyPaths, 3)
	expected := []TopPathStat {
		{
			DependencyPath: []string{"m1@v1", "m2@v2", "m3@v3"},
			Count: 2,
		}, {
			DependencyPath: []string{"m1@v1", "m2@v2", "m5@v5"},
			Count: 1,
		}, {
			DependencyPath: []string{"m1@v1", "m2@v2"},
			Count: 1,
		}, {
			DependencyPath: []string{"m1@v1", "m3@v3"},
			Count: 1,
		},
	}
	assert.Equal(t, expected, result)
}