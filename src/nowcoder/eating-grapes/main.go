package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readStdin() (lines []string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != io.EOF {
			lines = append(lines, line)
		} else {
			break
		}
	}
	return
}

func maxKey(m map[string]int) (key string) {
	maxValue := 0
	for k, v := range m {
		if v >= maxValue {
			maxValue = v
			key = k
		}
	}
	return
}

func solve(slice []string) (num int) {
	fMap := make(map[string]int)
	for i, s := range slice {
		n, _ := strconv.Atoi(s)
		switch i {
		case 0:
			fMap["a"] = n
		case 1:
			fMap["b"] = n
		case 2:
			fMap["c"] = n
		}
	}

	rMap := map[string][2]string{
		"a": {"1", "3"},
		"b": {"1", "2"},
		"c": {"2", "3"},
	}

	pMap := map[string]int{"1": 0, "2": 0, "3": 0}
	for fMap["a"]+fMap["b"]+fMap["c"] > 0 {
		key := maxKey(fMap)
		fMap[key]--
		switch key {
		case key:
			if pMap[rMap[key][0]] < pMap[rMap[key][1]] {
				pMap[rMap[key][0]]++
			} else {
				pMap[rMap[key][1]]++
			}
		}
	}

	num = pMap[maxKey(pMap)]
	return
}

func main() {
	lines := readStdin()
	for i, line := range lines {
		if i != 0 {
			line = strings.Replace(line, "\n", "", -1)
			line = strings.Replace(line, "\r", "", -1)
			slice := strings.Split(line, " ")
			num := solve(slice)
			fmt.Println(num)
		}
	}
}
