package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func removeFloat64Index(array []float64, index int) []float64 {
	var new_array []float64

	for i := 0; i < len(array); i++ {
		if i != index {
			new_array = append(new_array, array[i])
		}
	}

	return new_array
}

func removeStringIndex(array []string, index int) []string {
	var new_array []string

	for i := 0; i < len(array); i++ {
		if i != index {
			new_array = append(new_array, array[i])
		}
	}

	return new_array
}

func solve(nums []float64, signs []string) []float64 {
	for i := 0; i < len(signs); i++ {
		if strings.Compare(signs[i], "^") == 0 {
			nums[i] = math.Pow(nums[i], nums[i + 1])
			nums = removeFloat64Index(nums, i + 1)
			signs = removeStringIndex(signs, i)
		}
	}

	for i := 0; i < len(signs); i++ {
		if strings.Compare(signs[i], "*") == 0 {
			nums[i] = nums[i] * nums[i + 1]
			nums = removeFloat64Index(nums, i + 1)
			signs = removeStringIndex(signs, i)

		} else if strings.Compare(signs[i], "/") == 0 {
			nums[i] = nums[i] / nums[i + 1]
			nums = removeFloat64Index(nums, i + 1)
			signs = removeStringIndex(signs, i)
		}
	}

	for i := 0; i < len(signs); i++ {
		if strings.Compare(signs[i], "+") == 0 {
			nums[i] = nums[i] + nums[i + 1]
			nums = removeFloat64Index(nums, i + 1)
			signs = removeStringIndex(signs, i)

		} else if strings.Compare(signs[i], "-") == 0 {
			nums[i] = nums[i] - nums[i + 1]
			nums = removeFloat64Index(nums, i + 1)
			signs = removeStringIndex(signs, i)
		}
	}

	return nums
}

func main() {

	var nums []float64
	var signs []string
	//var vars map[string]float64

	var closed_par_idx int = -1
	var opened_par_idx int = -1

	fmt.Printf("=> ")
	reader := bufio.NewReader(os.Stdin)
	calc, _ := reader.ReadString('\n')
	calc = strings.TrimSpace(calc)
	calcs := strings.Split(calc, " ")

	for i := 0; i <= len(calcs) - 1; i++ {
		float_num, err := strconv.ParseFloat(calcs[i], 64)

		if err != nil {
			sign := calcs[i]
			signs = append(signs, sign)			
		} else {
			nums = append(nums, float_num)
		}
	}

	for i := 0; i < len(signs) - 1; i++ {
		if strings.Compare(signs[i], "(") == 0 {
			opened_par_idx = i
		} else if strings.Compare(signs[i], ")") == 0 {
			closed_par_idx = i
		}

		if closed_par_idx > 0 {
			break
		}
	}

	if closed_par_idx != opened_par_idx || closed_par_idx < opened_par_idx {
		fmt.Printf("Invalid expression")
	} else if closed_par_idx > 0 && opened_par_idx >= 0 {
		nums = solve(nums[opened_par_idx : closed_par_idx], signs[opened_par_idx : closed_par_idx])
	} else {
		nums = solve(nums, signs)
	}

	fmt.Printf("%f", nums[0])
}
