package routing

var Null = struct{}{}

type Set map[string]struct{}

func (dict Set) Add(payload string) {
	dict[payload] = Null
}

// If payload is nil or there is no such element, delete is a no-op.
func (dict Set) Remove(payload string) {
	delete(dict, payload)
}

/*
func (dict Set) Has(payload string) { // has up to 40% performance loss
	_, ok:=dict[payload]
	return ok
}
*/
