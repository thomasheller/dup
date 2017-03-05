package main

// DupeFinder finds map keys sharing the same value.
type DupeFinder struct{}

// Find finds map keys which map to the same value. It returns a
// multimap: keys are the duplicate values, values are sets containing
// the original keys sharing the same value.
func (df *DupeFinder) Find(m map[string]string) map[string][]string {
	mm := reverseMultiMap(m)

	dupes := make(map[string][]string)

	for k, set := range mm {
		if len(set) > 1 {
			dupes[k] = set
		}
	}

	return dupes
}

// reverseMultiMap transforms map key->value into reverse multimap
// value->key(s)
func reverseMultiMap(m map[string]string) map[string][]string {
	mm := make(map[string][]string)

	for k, v := range m {
		_, ok := mm[v]

		if ok {
			mm[v] = append(mm[v], k)
		} else {
			mm[v] = []string{k}
		}
	}

	return mm
}
