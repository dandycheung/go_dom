package dom

// Map returns a new slice by applying a "mapping" function to each element
// of the original slice
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, s := range vs {
		vsm[i] = f(s)
	}
	return vsm
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
