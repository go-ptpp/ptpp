package ptpp

// Levenshtein computes the Levenshtein distance for two words.
func Levenshtein(v, w string) int {
	vcs := []rune(v)
	wcs := []rune(w)

	d := make([]int, len(vcs)+1)
	for i := range d {
		d[i] = i
	}

	for i, wc := range wcs {
		last := d[0]
		d[0] = i
		for j, vc := range vcs {
			var min int
			if d[j+1] < d[j] {
				min = d[j+1] + 1
			} else {
				min = d[j] + 1
			}
			if vc != wc {
				if last+1 < min {
					min = last + 1
				}
			} else {
				if last < min {
					min = last
				}
			}
			last, d[j+1] = d[j+1], min
		}
	}

	return d[len(d)-1]
}
