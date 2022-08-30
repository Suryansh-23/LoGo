package main

import (
	"regexp"
	"strconv"
)

func getCommandsToLoop(s string) (string, int) {
	var re = regexp.MustCompile(`(repeat)(\x20+)([0-9]+)(\x20)+(\x5B)(?P<cmds>[[:ascii:]]*)(\x5D)`)
	matches := re.FindStringSubmatch(s)
	loop, _ := strconv.Atoi(matches[3])
	return matches[6], loop
}

func parseCmds(c []string) []string {
	cmds := []string{}
	for i, v := range c {
		if i%2 == 0 {
			cmds = append(cmds, v+" "+c[i+1])
		}
	}
	return cmds
}
