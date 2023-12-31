package main

import (
	"testing"
)

func TestAlignRanges(t *testing.T) {
	r := rangeMapping{destStart: 0, sourceStart: 10, length: 15}
	for i := 1; i < 45; i++ {
		c := entityRange{start: 16, length: i}
		ars, unalignedRanges, _ := alignRanges(r, c)
		// if done {
		// 	t.Fatalf("Should not have been done")
		// }
		if len(ars) == 0 && len(unalignedRanges) == 0 {
			continue
		}
		if totalLength(ars)+totalLength(unalignedRanges) != c.length {
			t.Fatalf("Inconsistent lengths")
		}
	}
}
