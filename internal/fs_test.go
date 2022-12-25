package internal

import (
	"fmt"
	"testing"
)

func Test_ListFiles(t *testing.T) {
	suite := []struct {
		path  string
		level int
		want  []string
	}{
		{
			path:  "**/*",
			level: 0,
			want:  []string{"molerc.json", "nested"},
		},
		{
			path:  "**/*",
			level: 1,
			want: []string{
				"molerc.json",
				"nested",
				"nested/lvl1",
			},
		},
		{
			path:  "**/*",
			level: 2,
			want: []string{
				"molerc.json",
				"nested",
				"nested/lvl1",
				"nested/lvl1/f1.txt",
				"nested/lvl1/lvl2",
			},
		},
		{
			path:  "**/*",
			level: 3,
			want: []string{
				"molerc.json",
				"nested",
				"nested/lvl1",
				"nested/lvl1/f1.txt",
				"nested/lvl1/lvl2",
				"nested/lvl1/lvl2/f1.txt",
				"nested/lvl1/lvl2/lvl3",
			},
		},
		{
			path:  "**/*.txt",
			level: 3,
			want: []string{
				"nested/lvl1/f1.txt",
				"nested/lvl1/lvl2/f1.txt",
			},
		},
	}
	root := GetTestFilePath("/")
	for _, test := range suite {
		name := fmt.Sprintf("level %d", test.level)
		t.Run(name, func(t *testing.T) {
			got := ListFiles(root+"/"+test.path, test.level)
			if len(got) != len(test.want) {
				t.Log("want", test.want)
				t.Log("got", got)
				t.Fatalf(
					"path length mismatch, want %d got %d",
					len(test.want), len(got))
			}
			for _, want := range test.want {
				matches := false
				for _, item := range got {
					if root+"/"+want == item {
						matches = true
						break
					}
				}
				if !matches {
					t.Log("got", got)
					t.Fatalf(
						"part mismatch: want %q",
						root+"/"+want)
				}
			}
		})
	}
}

func Test_createPathList(t *testing.T) {
	suite := []struct {
		path  string
		level int
		want  []string
	}{
		{
			path:  "test/**/*.go",
			level: 0,
			want:  []string{"test/*.go"},
		}, {
			path:  "test/**/other/*.txt",
			level: 1,
			want: []string{
				"test/other/*.txt",
				"test/*/other/*.txt",
			},
		},
		{
			path:  "test/**/*.go",
			level: 2,
			want: []string{
				"test/*.go",
				"test/*/*.go",
				"test/*/*/*.go",
			},
		}, {
			path:  "test/**/other/*.txt",
			level: 3,
			want: []string{
				"test/other/*.txt",
				"test/*/other/*.txt",
				"test/*/*/other/*.txt",
				"test/*/*/*/other/*.txt",
			},
		},
	}
	for _, test := range suite {
		t.Run(test.path, func(t *testing.T) {
			got := createPathList(test.path, test.level)
			if len(got) != len(test.want) {
				t.Log("want", test.want)
				t.Log("got", got)
				t.Fatalf(
					"path length mismatch, want %d got %d",
					len(test.want), len(got))
			}
			for idx, want := range test.want {
				if got[idx] != want {
					t.Log("want", test.want)
					t.Log("got", got)
					t.Fatalf(
						"part mismatch: want %q got %q",
						want, got[idx])
				}
			}
		})
	}
}

func Test_extractRoot(t *testing.T) {
	suite := map[string][]string{
		"test/**/*.go": []string{
			"test/", "*.go",
		},
		"test/**/other/*.txt": []string{
			"test/", "other/*.txt",
		},
	}
	for test, want := range suite {
		t.Run(test, func(t *testing.T) {
			root, rest := extractRoot(test)
			if root != want[0] {
				t.Errorf("want %q, got %q",
					want[0], root)
			}
			if rest != want[1] {
				t.Errorf("want %q, got %q",
					want[1], rest)
			}
		})
	}
}
