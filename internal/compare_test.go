package internal

import "testing"

func Test_ComparisonAll(t *testing.T) {
	suite := map[string]struct {
		test Comparison
		want []string
	}{
		"no changes": {
			test: Comparison{},
			want: []string{},
		},
		"add": {
			test: Comparison{
				add: Filelist{"f1": "h1"},
			},
			want: []string{"f1"},
		},
		"chg": {
			test: Comparison{
				chg: Filelist{"f1": "h1"},
			},
			want: []string{"f1"},
		},
		"rmv": {
			test: Comparison{
				rmv: Filelist{"f1": "h1"},
			},
			want: []string{"f1"},
		},
		"all": {
			test: Comparison{
				add: Filelist{"f1": "h1"},
				chg: Filelist{"f2": "h2"},
				rmv: Filelist{"f3": "h3"},
			},
			want: []string{"f1", "f2", "f3"},
		},
	}
	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			got := test.test.All()
			if len(test.want) != len(got) {
				t.Log(test.want)
				t.Log(got)
				t.Fatalf(
					"len: want %d, got %d",
					len(test.want), len(got))
			}
		})
	}
}

func Test_ComparisonAny(t *testing.T) {
	suite := map[string]struct {
		test Comparison
		want bool
	}{
		"empty comparison": {
			test: Comparison{},
			want: false,
		},
		"add comparison": {
			test: Comparison{
				add: Filelist{"new": "file"},
			},
			want: true,
		},
		"chg comparison": {
			test: Comparison{
				chg: Filelist{"new": "file"},
			},
			want: true,
		},
		"rmv comparison": {
			test: Comparison{
				rmv: Filelist{"new": "file"},
			},
			want: true,
		},
	}
	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			got := test.test.Any()
			if test.want != got {
				t.Errorf(
					"want %v, got %v",
					test.want, got)
			}
		})
	}
}

func Test_CompareFilelists(t *testing.T) {
	suite := map[string]struct {
		old  Filelist
		new  Filelist
		want Comparison
	}{
		"simple, no diff": {
			old: Filelist{
				"file 1": "mock hash 1",
			},
			new: Filelist{
				"file 1": "mock hash 1",
			},
			want: Comparison{},
		},
		"one new file": {
			old: Filelist{
				"file 1": "mock hash 1",
			},
			new: Filelist{
				"file 1": "mock hash 1",
				"file 2": "mock hash 2",
			},
			want: Comparison{
				add: Filelist{
					"file 2": "mock hash 2",
				},
			},
		},
		"one removed file": {
			old: Filelist{
				"file 1": "mock hash 1",
				"file 2": "mock hash 2",
			},
			new: Filelist{
				"file 1": "mock hash 1",
			},
			want: Comparison{
				rmv: Filelist{
					"file 2": "mock hash 2",
				},
			},
		},
		"one changed file": {
			old: Filelist{
				"file 1": "mock hash 1",
				"file 2": "mock hash 2",
			},
			new: Filelist{
				"file 1": "mock hash 1",
				"file 2": "mock hash 2a",
			},
			want: Comparison{
				chg: Filelist{
					"file 2": "mock hash 2a",
				},
			},
		},
	}

	for name, test := range suite {
		t.Run(name, func(t *testing.T) {
			got := CompareFilelists(&test.new, &test.old)
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
