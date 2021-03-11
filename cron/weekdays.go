package cron

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	sun = iota
	mon
	tue
	wen
	thu
	fri
	sat
)

var week = map[string]int{
	"SUN": sun,
	"MON": mon,
	"TUE": tue,
	"WEN": wen,
	"THU": thu,
	"FRI": fri,
	"SAT": sat,
}

var fullWeek = []int{
	sun,
	mon,
	tue,
	wen,
	thu,
	fri,
	sat,
}

func (c *Cron) parseWeekdays(s string) (err error) {
	w := make([]int, 0)
	divider := 1
	if s == "*" {
		w = fullWeek
		c.Weekdays = w
		return
	}

	if strings.Contains(s, "/") {
		divider, err = strconv.Atoi(strings.Split(s, "/")[1])
		if err != nil {
			return
		}
	}
	num := 0
	for _, str := range strings.Split(strings.Split(s, "/")[0], ",") {
		if str == "*" {
			for i := 0; i <= len(week)-divider; i += divider {
				w = append(w, i)
			}
			c.Weekdays = w
			return
		}
		if wd, ok := week[str]; ok {
			w = append(w, wd)
			continue
		}
		num, err = strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("cannot parse weekday: %s", err)
		}

		if num < 0 || num > 6 {
			return fmt.Errorf("weekday is out of range: %d", num)
		}

		w = append(w, num)
	}
	return
}
