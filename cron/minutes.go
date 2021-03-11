package cron

import (
	"fmt"
	"strconv"
	"strings"
)

func (c *Cron) parseMinutes(s string) (err error) {
	m := make([]int,0)
	if s == "*" {
		for i := 0; i <= 59; i++ {
			m = append(m, i)
		}
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
			for i := 0; i <= 59 + 1 - divider; i += divider {
				m = append(m, i)
			}
			c.Minutes = m
			return
		}
		num, err = strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("cannot parse minutes: %s", err)
		}

		if num < 0 || num > 59 {
			return fmt.Errorf("minutes is out of range: %d", num)
		}

		m = append(m, num)
	}
	c.Minutes = m
	return
}


