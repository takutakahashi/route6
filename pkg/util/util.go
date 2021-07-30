package util

// MatchSelector returns true if all b's k, v contain a.
func MatchSelector(a, b map[string]string) bool {
	ret := []int{}
	for k, v := range b {
		if a[k] == v {
			ret = append(ret, 1)
		}
	}
	return len(ret) == len(b)
}
