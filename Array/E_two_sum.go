package array

import (
	"fmt"
	"time"
)

func TwoSum() {
	nums := []int{2, 15, 11, 7}
	t := 9

	kn := make(map[int]int)

	func() {
		println("Best way")
		defer println("Best way")

		start := time.Now()

		for i, num := range nums {
			fmt.Println("Num:", num)
			fmt.Println("Map: ", kn)
			if index, ok := kn[t-num]; ok {
				fmt.Println("Out:", []int{index, i})

				fmt.Println("Elapsed Time:", time.Since(start))
				return
			}

			kn[num] = i
			fmt.Println("Know num:", kn[num])
		}
		println()
	}()

	func() {
		println("Worst way")
		defer println("Worst way")

		start := time.Now()

		for i, num := range nums {
			for j := i + 1; j < len(nums); j++ {
				fmt.Printf("I: %d; J: %d\n", nums[i], nums[j])
				if num+nums[j] == t {
					fmt.Println("Elapsed Time:", time.Since(start))
					return
				}
			}
		}
	}()
}
