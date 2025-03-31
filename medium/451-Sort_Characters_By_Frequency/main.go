package main

import (
	"bytes"
	"slices"
)

type ele struct {
	k    rune
	freq int32
}

func frequencySort(s string) string {
	m := make(map[rune]ele)
	for _, r := range s {
		e := m[r]
		e.k = r
		e.freq += 1
		m[r] = e
	}
	sli := make([]ele, 0, len(m))
	for _, e := range m {
		sli = append(sli, e)
	}
	slices.SortFunc(sli, func(a, b ele) int {
		if a.freq > b.freq {
			return -1
		}
		if a.freq < b.freq {
			return 1
		}
		return 0
	})
	var buf bytes.Buffer
	for _, e := range sli {
		for i := 0; i < int(e.freq); i++ {
			buf.WriteRune(e.k)
		}
	}
	return buf.String()
}
