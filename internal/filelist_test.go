package internal

import "testing"

func Test_NewHash_HappyPath(t *testing.T) {
	hash, err := NewHash(GetTestFilePath("nested/lvl1/f1.txt"))
	if err != nil {
		t.Errorf("expected hash, got error: %v", err)
	}
	if len(hash) == 0 {
		t.Errorf("invalid hash: %q", hash)
	}
}

func Test_NewHash_Error(t *testing.T) {
	hash, err := NewHash("wat")
	if err == nil {
		t.Errorf("expected error, got hash: %q", hash)
	}
	if len(hash) != 0 {
		t.Errorf("invalid hash: %q", hash)
	}
}

func Test_NewFilelist(t *testing.T) {
	p := NewPathlist(
		[]string{GetTestFilePath("/**/*.txt")},
		DefaultGlobDepth)
	src := ListFiles(p)
	lst := NewFilelist(src)
	if len(src) != len(lst) {
		t.Log(lst)
		t.Errorf("len mismatch: want %+v, got %+v", src, lst)
	}
}
