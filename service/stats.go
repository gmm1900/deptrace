package service

import (
	"sort"
	"strings"
)

type TopPathStat struct {
	DependencyPath DependencyPath
	Count int
}

func ExtractTopPaths(dependencyPaths []DependencyPath, topLevels int) []TopPathStat {
	stats := map[string]*TopPathStat{}
	for _, p := range dependencyPaths {
		truncateAt := topLevels
		if len(p) < topLevels {
			truncateAt = len(p)
		}
		truncatedPath := p[0:truncateAt]
		key := keyOf(truncatedPath)
		if _, found := stats[key]; !found {
			stats[key] = &TopPathStat{
				DependencyPath: truncatedPath,
				Count: 0,
			}
		}
		stats[key].Count++
	}

	var result []TopPathStat
	for _, v := range stats {
		result = append(result, *v)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})
	
	return result
}

func keyOf(dependencyPath DependencyPath) string {
	return strings.Join(dependencyPath, "->")
}