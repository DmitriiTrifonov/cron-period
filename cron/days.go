package cron

import (
	"fmt"
	"strconv"
	"strings"
)

const maxDays = 31

func (c *Cron) parseDays(s string) (err error) {
	d := make([]int,0)
	if s == "*" {
		for i := 1; i <= 31; i++ {
			d = append(d, i)
		}
		c.Days = d
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
			for i := 1; i <= maxDays+ 1 - divider; i += divider {
				d = append(d, i)
			}
			c.Days = d
			return
		}
		num, err = strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("cannot parse weekday: %s", err)
		}

		if num < 1 || num > 31 {
			return fmt.Errorf("weekday is out of range: %d", num)
		}

		d = append(d, num)
	}
	c.Days = d
	return
}
