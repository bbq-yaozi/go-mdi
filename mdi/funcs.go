package mdi

var namedIconMap = map[string]*Icon{
	"ab-testing": &AbTesting,
}

func Name(name string) *Icon {
	if v, ok := namedIconMap[name]; ok {
		return v
	}
	return nil
}

func Icons() []*Icon {
	var result []*Icon = make([]*Icon, 0, len(namedIconMap))
	for _, v := range namedIconMap {
		result = append(result, v)
	}
	return result
}
