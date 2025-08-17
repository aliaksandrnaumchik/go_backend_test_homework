// list.go
package list

func Min(list []int) int {
	var min int

	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}
