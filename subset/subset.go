package subset

import "fmt"

func Print(s []int32) {
	subset := []int32{}
	res := [][]int32{}

	var backtrac func(subset []int32, index int)
	backtrac = func(subset []int32, index int) {
		res = append(res, subset)

		for i := index; i < len(s); i++ {
			tmp := make([]int32, len(subset))
			copy(tmp, subset)
			tmp = append(tmp, s[i])

			backtrac(tmp, i+1)

			tmp = tmp[:len(tmp)-1]
		}
		return
	}
	backtrac(subset, 0)

	fmt.Print(res)
}
