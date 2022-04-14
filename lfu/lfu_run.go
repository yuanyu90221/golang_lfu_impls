package lfu

import "fmt"

func RunLFU(actions []string, value [][]int) []string {
	length := len(actions)
	result := make([]string, length)
	result[0] = "null"
	lfu := Constructor(value[0][0])
	for idx := 1; idx < length; idx++ {
		action := actions[idx]
		switch action {
		case "put":
			lfu.Put(value[idx][0], value[idx][1])
			result[idx] = "null"
		case "get":
			result[idx] = fmt.Sprintf("%d", lfu.Get(value[idx][0]))
		}
	}
	return result
}
