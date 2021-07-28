package kmp

// Kmp struct
type Kmp struct {
	// store the next slice
	next []int
	// store the pattern string
	pattern string
}

// New funciton will return a kmp struct
func New(pattern string) *Kmp {
	length := len(pattern)
	// if the pattern sting is empty string
	// return the nil struct
	if length == 0 {
		return nil
	}

	next := make([]int, length)
	next[0] = -1
	i, j := 0, -1
	for i < length-1 {
		// when j becomes -1 or when pattern strings are equals
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			// optimizing the next slice
			if pattern[i] != pattern[j] {
				next[i] = j
			} else {
				next[i] = next[j]
			}
		} else {
			j = next[j]
		}
	}
	kmp := &Kmp{
		pattern: pattern,
		next:    next,
	}
	return kmp
}

// Index function will return the index in the string of the pattern
func (k *Kmp) Index(s string) int {
	len1, len2 := len(s), len(k.pattern)
	i, j := 0, 0
	for i < len1 && j < len2 {
		if j == -1 || s[i] == k.pattern[j] {
			i++
			j++
		} else {
			j = k.next[j]
		}
	}
	if j == len2 {
		return i - j
	} else {
		return -1
	}
}

func Index(s, pattern string) int {
	kmp := New(pattern)
	return kmp.Index(s)
}
