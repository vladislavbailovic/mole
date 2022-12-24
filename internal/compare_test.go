package internal

import "testing"

func Test_CompareFilelists(t *testing.T) {
	suite := map[string]struct {
		old  map[string]hash
		new  map[string]hash
		want comparison
	}{
		"simple, no diff": {
			old: map[string]hash{
				"file 1": "mock hash 1",
			},
			new: map[string]hash{
				"file 1": "mock hash 1",
			},
			want: comparison{},
		},
		"one new file": {
			old: map[string]hash{
				"file 1": "mock hash 1",
			},
			new: map[string]hash{
				"file 1": "mock hash 1",
				"file 2": "mock hash 2",
			},
			want: comparison{
				add: map[string]hash{
					"file 2": "mock hash 2",
				},
			},
		},
		"one removed file": {
			old: map[string]hash{
				"file 1": "mock hash 1",
				"file 2": "mock hash 2",
			},
			new: map[string]hash{
				"file 1": "mock hash 1",
			},
			want: comparison{
				rmv: map[string]hash{
					"file 2": "mock hash 2",
				},
			},
		},
		"one changed file": {
			old: map[string]hash{
				"file 1": "mock hash 1",
				"file 2": "mock hash 2",
			},
			new: map[string]hash{
				"file 1": "mock hash 1",
				"file 2": "mock hash 2a",
			},
			want: comparison{
				chg: map[string]hash{
					"file 2": "mock hash 2a",
				},
			},
		},
	}

	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			got := CompareFilelists(test.new, test.old)
			if len(got.add) != len(test.want.add) ||
				len(got.chg) != len(test.want.chg) ||
				len(got.rmv) != len(test.want.rmv) {
				t.Logf("wnt: %+v", test.want)
				t.Logf("got: %+v", got)
				t.Fatal("len mismatch")
			}
		})
	}
}
