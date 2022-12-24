package internal

import (
	"fmt"
	"os"
)

type hash string

func NewHash(path string) (hash, error) {
	nfo, err := os.Stat(path)
	if err != nil {
		return hash(""), err
	}
	raw := fmt.Sprintf("%d-%v",
		nfo.Size(), nfo.ModTime())
	return hash(raw), nil
}

type Filelist map[string]hash

func NewFilelist(files []string) Filelist {
	lst := make(Filelist, len(files))
	for _, f := range files {
		if hash, err := NewHash(f); err != nil {
			// TODO: report error?
			continue
		} else {
			lst[f] = hash
		}
	}
	return lst
}
