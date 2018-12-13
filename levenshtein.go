package maxleven

// Context provides a reusable int slice
type Context struct {
	intSlice []int
}

func (c *Context) getIntSlice(l int) []int {
	if cap(c.intSlice) < l {
		c.intSlice = make([]int, l)
	}
	return c.intSlice[:l]
}

// Distance is a wrapper for calling the distance function with the context struct
func Distance(s1, s2 []rune, maxDist int) int {
	c := Context{}
	return c.Distance(s1, s2, maxDist)
}

// Distance between two strings is defined as the minimum
// number of edits needed to transform one string into the other, with the
// allowable edit operations being insertion, deletion, or substitution of
// a single character
// http://en.wikipedia.org/wiki/Levenshtein_distance
//
// This implemention is optimized to use O(min(m,n)) space.
// It is based on the optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
// This version is modified to return early if maxDist is exceeded, the dist returned will be -1
func (c *Context) Distance(s1, s2 []rune, maxDist int) int {

	lenS1 := len(s1)
	lenS2 := len(s2)

	if lenS2 == 0 {
		if lenS1 <= maxDist {
			return lenS1
		}
		return -1
	}

	column := c.getIntSlice(lenS1 + 1)
	// Column[0] will be initialised at the start of the first loop before it
	// is read, unless lenS2 is zero, which we deal with above
	for i := 1; i <= lenS1; i++ {
		column[i] = i
	}
	for x := 0; x < lenS2; x++ {
		s2Rune := s2[x]
		column[0] = x + 1
		lastdiag := x
		var currMin int
		for y := 0; y < lenS1; y++ {
			olddiag := column[y+1]
			cost := 0
			if s1[y] != s2Rune {
				cost = 1
			}
			column[y+1] = min(
				column[y+1]+1,
				column[y]+1,
				lastdiag+cost,
			)
			if y == 0 || column[y+1] < currMin {
				currMin = column[y+1]
			}
			lastdiag = olddiag
		}
		if currMin > maxDist {
			return -1
		}
	}

	if column[lenS1] > maxDist {
		return -1
	}

	return column[lenS1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
