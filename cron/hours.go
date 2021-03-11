package cron

import (
	"fmt"
	"strconv"
	"strings"
)

func (c *Cron) parseHours(s string) (err error) {
	h := make([]int,0)
	if s == "*" {
		for i := 0; i <= 23; i++ {
			h = append(h, i)
		}
		c.Hours = h
		return
	}

	divider := 1

	if strings.Contains(s, "/") {
		divider, err = strconv.Atoi(strings.Split(s, "/")[1])
		if err != nil {
			return
		}
	}

	num := 0
	for _, str := range strings.Split(strings.Split(s, "/")[0], ",") {
		if str == "*" {
			for i := 0; i <= 23 + 1 - divider; i += divider {
				h = append(h, i)
			}
			c.Hours = h
			return
		}
		num, err = strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("cannot parse hour: %s", err)
		}

		if num < 0 || num > 23 {
			return fmt.Errorf("hour is out of range: %d", num)
		}

		h = append(h, num)
	}
	c.Hours = h
	return
}
