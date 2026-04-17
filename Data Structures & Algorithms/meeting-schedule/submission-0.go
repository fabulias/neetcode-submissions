/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	for ix := 1; ix < len(intervals); ix++ {
		if intervals[ix-1].end > intervals[ix].start {
			return false
		}
	}
	return true
}
