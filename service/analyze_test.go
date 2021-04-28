package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyze(t *testing.T) {
	testCases := []struct{
		name string
		givenInputs []Input
		expectedTree *AnalyzedTree
	} {
		{
			name:         "no input",
			givenInputs:  []Input{},
			expectedTree: nil,
		},
		{
			name: "multiple nodes, with circular dependencies",
			givenInputs: []Input{
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
			},
			expectedTree: func() *AnalyzedTree{
				m1 := Node{ VersionedModule: VersionedModule{"m1", "v0.1"}}
				m2 := Node{ VersionedModule: VersionedModule{"m2", "v0.2"}}
				m3 := Node{ VersionedModule: VersionedModule{"m3", "v0.3"}}
				m3_1 := Node{ VersionedModule: VersionedModule{"m3", "v0.3.1"}}
				m4 := Node{ VersionedModule: VersionedModule{"m4", "v0.4"}}
				m1.Upstreams = nil
				m1.Downstreams = []*Node{&m2, &m3}
				m2.Upstreams = []*Node{&m1, &m3}
				m2.Downstreams = []*Node{&m3, &m4}
				m3.Upstreams = []*Node{&m1, &m2}
				m3.Downstreams = []*Node{&m2}
				m3_1.Upstreams = []*Node{&m4}
				m3_1.Downstreams = nil
				m4.Upstreams = []*Node{&m2}
				m4.Downstreams = []*Node{&m3_1}
				return &AnalyzedTree{
					Root: &m1,
					SearchMap: map[string]map[string]*Node{
						"m1": {
							"v0.1": &m1,
						},
						"m2": {
							"v0.2": &m2,
						},
						"m3": {
							"v0.3": &m3,
							"v0.3.1": &m3_1,
						},
						"m4": {
							"v0.4": &m4,
						},
					},
				}
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tree := Analyze(tc.givenInputs)
			assert.Equal(t, tc.expectedTree, tree)
		})
	}
}