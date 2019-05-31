package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func main() {
	i1 := Interval{Start:2, End:3}
	i2 := Interval{Start:4, End:5}
	i3 := Interval{Start:4, End:6}
	i4 := Interval{Start:4, End:4}
	i5 := Interval{Start:4, End:10}
	i6 := Interval{Start:5, End:9}

	ret := merge([]Interval{i1, i2, i3, i4, i5, i6})
	fmt.Println("ret=", ret)
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	src := Intervals(intervals)
	sort.Sort(src) // 快排使之顺序，根据start排序

	dst := []Interval{src[0]}

	for i := 0; i < src.Len(); i++ {
		if dst[len(dst)-1].End < src[i].Start {
			dst = append(dst, src[i]) // 上一个的end与排序的start比较大小
		} else {
			dst[len(dst)-1].End = max(dst[len(dst)-1].End, src[i].End)
		}
	}

	return dst

}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Intervals []Interval

func (is Intervals) Len() int {
	return len(is)
}

func (is Intervals) Less(i, j int) bool {
	return is[i].Start < is[j].Start
}

func (is Intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}
