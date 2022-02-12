package NC37

import "sort"

type Interval struct {
	Start int
	End   int
}

/**
 *
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
func merge(intervals []*Interval) []*Interval {
	var res []*Interval
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	if len(intervals) <= 1 {
		return intervals
	}
	start := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if start.End >= intervals[i].Start {
			start.End = max(start.End, intervals[i].End)
		} else {
			res = append(res, start)
			start = intervals[i]
		}
	}
	res = append(res, start)
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
