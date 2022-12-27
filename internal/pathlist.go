package internal

var DefaultGlobDepth int = 2

type Pathlist []string

func NewPathlist(from []string, globDepth int) *Pathlist {
	count := len(from) +
		// Assume one double start per path anyway
		len(from)*(globDepth+1)
	list := make(Pathlist, 0, count)

	for _, p := range from {
		list = append(list, createPathList(p, globDepth)...)
	}

	lst := Pathlist(list)
	return &lst
}
