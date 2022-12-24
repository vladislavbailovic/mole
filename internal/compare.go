package internal

type hash string

type filelist map[string]hash

type comparison struct {
	add filelist
	rmv filelist
	chg filelist
}

func CompareFilelists(n filelist, o filelist) comparison {
	var add, chg, rmv filelist
	if len(n) > len(o) {
		add = make(filelist, len(n)-len(o))
	} else {
		add = make(filelist, 0)
	}
	if len(n) < len(o) {
		rmv = make(filelist, len(n)-len(o))
	} else {
		rmv = make(filelist, 0)
	}
	chg = make(filelist, 0)

	for path, hash := range n {
		// fmt.Printf("\t- path %q ", path)
		if old, ok := o[path]; ok {
			// fmt.Printf("DOES exist in old files")
			if hash != old {
				// fmt.Printf(", but the hash: want %q got %q", hash, old)
				chg[path] = hash
			}
		} else {
			// fmt.Printf("does NOT exist in old files")
			add[path] = hash
		}
		// fmt.Println()
	}

	for path, hash := range o {
		if _, ok := n[path]; ok {
			continue
		}
		// fmt.Printf("\t- path %q is REMOVED\n", path)
		rmv[path] = hash
	}

	return comparison{add: add, chg: chg, rmv: rmv}
}
