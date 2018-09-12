package main

import (
	"sort"
	"testing"
)

func Test_v_min(t *testing.T) {
	var entrys = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	N := len(entrys)
	t.Logf("a=%d:%v\n", N, entrys)

	cond := 4

	iFound := sort.Search(N, func(idx int) bool {
		return entrys[idx] >= cond
	})
	if iFound < 0 {
		t.Logf("not found %d, return %d\n", cond, iFound)
	} else if N == iFound {
		t.Logf("not found %d, return %d\n", cond, iFound)
	} else {
		t.Logf("found %d at a[%d]=%d\n", cond, iFound, entrys[iFound])
	}
}

func Test_v_max(t *testing.T) {
	var entrys = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	N := len(entrys)
	t.Logf("a=%d:%v\n", N, entrys)

	cond := 4

	iFound := sort.Search(N, func(idx int) bool {
		idx = N - idx - 1

		return entrys[idx] <= cond
	})
	if iFound < 0 {
		t.Logf("not found %d, return %d\n", cond, iFound)
	} else if N == iFound {
		t.Logf("not found %d, return %d\n", cond, iFound)
	} else {
		t.Logf("found %d at a[%d]=%d\n", cond, iFound, entrys[iFound])
	}
}

type Range struct {
	begin int
	end   int
}

func (me *Range) Compare(v Range) int {
	if me.end < v.begin {
		// |--------- me ---------|
		//                            |----- v -----|
		return -1
	} else if me.begin > v.end {
		//                  |--------- me ---------|
		// |----- v -----|
		return 1
	} else {
		//            |--------- me ---------|
		// |----- v -----|
		//                 |----- v -----|
		//                                 |----- v -----|
		return 0
	}
}

var ranges = []Range{
	0: Range{0, 1},
	1: Range{1, 2},
	2: Range{2, 3},
	3: Range{3, 4},
	4: Range{4, 5},
	5: Range{5, 6},
	6: Range{6, 7},
	7: Range{7, 8},
	8: Range{8, 9},
}

var rangCond = Range{3, 6}

func Test_range_min(t *testing.T) {
	N := len(ranges)

	cond := rangCond

	iFound := sort.Search(N, func(idx int) bool {
		// 找出最小的entry idx，确保 entry <= cond
		cmp := ranges[idx].Compare(cond)

		return cmp >= 0
	})
	if iFound < 0 {
		t.Logf("not found %v, return %d\n", cond, iFound)
	} else if N == iFound {
		t.Logf("not found %v, return %d\n", cond, iFound)
	} else {
		t.Logf("found %v at ranges[%d]=%v\n", cond, iFound, ranges[iFound])
	}
}

func Test_range_max(t *testing.T) {
	N := len(ranges)

	cond := rangCond

	iFound := sort.Search(N, func(idx int) bool {
		// 找出最小的entry idx，确保 entry < cond
		cmp := ranges[idx].Compare(cond)

		return cmp > 0
	})
	if iFound < 0 {
		t.Logf("not found %v, return %d\n", cond, iFound)
	} else if N == iFound {
		t.Logf("not found %v, return %d\n", cond, iFound)
	} else {
		t.Logf("found %v at ranges[%d]=%v\n", cond, iFound, ranges[iFound])
	}
}
