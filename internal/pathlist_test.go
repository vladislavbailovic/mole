package internal

import "testing"

func Test_NewPathlist(t *testing.T) {
	suite := map[string]struct {
		test  []string
		depth int
		want  int
	}{
		"simple": {
			test:  []string{"*.txt"},
			depth: DefaultGlobDepth,
			want:  1,
		},
		"simple two paths": {
			test:  []string{"*.txt", "*.md"},
			depth: DefaultGlobDepth,
			want:  2,
		},
		"expr 0": {
			test:  []string{"./**/*.go"},
			depth: 0,
			want:  1,
		},
		"expr 1": {
			test:  []string{"./**/*.go"},
			depth: 1,
			want:  2,
		},
		"expr 2": {
			test:  []string{"./**/*.go"},
			depth: 2,
			want:  3,
		},
		"two exprs 0": {
			test:  []string{"./**/*.go", "./**/*_test.go"},
			depth: 0,
			want:  2,
		},
		"two exprs 1": {
			test:  []string{"./**/*.go", "./**/*_test.go"},
			depth: 1,
			want:  4,
		},
		"two exprs 2": {
			test:  []string{"./**/*.go", "./**/*_test.go"},
			depth: 2,
			want:  6,
		},
	}
	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			got := NewPathlist(test.test, test.depth)
			if len(*got) != test.want {
				t.Log("src", test.test)
				t.Log("got", *got)
				t.Errorf("wanted %d, got %d",
					test.want, len(*got))
			}
		})
	}
}
