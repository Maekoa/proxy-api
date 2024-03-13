package routing

func (dict Set) MatchSuffix(domain string) bool {
	if _, ok := dict[domain]; ok {
		return true
	}
	ds := []byte(domain)
	for i := 0; i < len(ds); i++ {
		if ds[i] == '.' {
			if _, ok := dict[string(ds[i+1:])]; ok {
				return true
			}
		}
	}
	return false
}
