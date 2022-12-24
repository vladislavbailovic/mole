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

func (x Comparison) Added() []string {
	result := make([]string, 0, len(x.add))
	for f := range x.add {
		result = append(result, f)
	}
	return result
}

func (x Comparison) Changed() []string {
	result := make([]string, 0, len(x.chg))
	for f := range x.chg {
		result = append(result, f)
	}
	return result
}

func (x Comparison) Removed() []string {
	result := make([]string, 0, len(x.rmv))
	for f := range x.rmv {
		result = append(result, f)
	}
	return result
}

func (x Comparison) All() []string {
	result := make([]string, 0,
		len(x.add)+len(x.chg)+len(x.rmv))

	result = append(result, x.Added()...)
	result = append(result, x.Changed()...)
	result = append(result, x.Removed()...)

	return result
}
