package internal

type Comparison struct {
	add Filelist
	rmv Filelist
	chg Filelist
}

func CompareFilelists(n *Filelist, o *Filelist) Comparison {
	var add, chg, rmv Filelist
	if len(*n) > len(*o) {
		add = make(Filelist, len(*n)-len(*o))
	} else {
		add = make(Filelist, 0)
	}
	if len(*n) < len(*o) {
		rmv = make(Filelist, len(*n)-len(*o))
	} else {
		rmv = make(Filelist, 0)
	}
	chg = make(Filelist, 0)

	for path, hash := range *n {
		// fmt.Printf("\t- path %q ", path)
		if old, ok := (*o)[path]; ok {
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

	for path, hash := range *o {
		if _, ok := (*n)[path]; ok {
			continue
		}
		// fmt.Printf("\t- path %q is REMOVED\n", path)
		rmv[path] = hash
	}

	return Comparison{add: add, chg: chg, rmv: rmv}
}

func (x Comparison) Any() bool {
	return len(x.add) > 0 ||
		len(x.chg) > 0 ||
		len(x.rmv) > 0
}
