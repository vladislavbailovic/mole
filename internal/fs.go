package internal

import (
	"path/filepath"
	"strings"
)

var doubleStarMaxGlobDepth int = 5

func ListFiles(lst *Pathlist) []string {
	tmp := map[string]struct{}{}
	for _, path := range *lst {
		for _, f := range listFiles(path) {
			tmp[f] = struct{}{}
		}
	}

	result := make([]string, 0, len(tmp))
	for f := range tmp {
		result = append(result, f)
	}

	return result
}

func listFiles(expr string) []string {
	var result []string
	if list, err := filepath.Glob(expr); err == nil {
		for _, f := range list {
			result = append(result, f)
		}
	}
	return result
}

func extractRoot(expr string) (string, string) {
	res := strings.SplitN(expr, "**", 2)
	return strings.TrimSuffix(res[0], "/") + "/",
		strings.TrimPrefix(res[1], "/")
}

func createPathList(expr string, globDepth int) []string {
	if !strings.Contains(expr, "**") {
		return []string{expr}
	}

	root, rest := extractRoot(expr)
	roots := make([]string, 0, globDepth+1)
	for i := 0; i <= globDepth; i++ {
		roots = append(
			roots,
			root+strings.Repeat("*/", i)+rest)
	}

	return roots
}
