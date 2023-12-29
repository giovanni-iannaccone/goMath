package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func removeIndex(array []interface{}, index int) []interface{} {
	var new_array []interface{}

	for i := 0; i < len(array); i++ {
		if i != index {
			new_array = append(new_array, array[i])
		}
	}

	return new_array
}

func solve(nums_signs []interface{}) []interface{} {
	for i := 0; i < len(nums_signs); {
		if i+1 < len(nums_signs) {
			if nums_signs[i] == "^" {
				nums_signs[i-1] = math.Pow(nums_signs[i-1].(float64), nums_signs[i+1].(float64))
				nums_signs = removeIndex(nums_signs, i)
				nums_signs = removeIndex(nums_signs, i)
			} else {
				i++
			}
		} else {
			i++
		}
	}

	for i := 0; i < len(nums_signs); {
		if i+1 < len(nums_signs) {
			if nums_signs[i] == "*" {
				nums_signs[i-1] = nums_signs[i-1].(float64) * nums_signs[i+1].(float64)
				nums_signs = removeIndex(nums_signs, i)
				nums_signs = removeIndex(nums_signs, i)
			} else if nums_signs[i] == "/" {
				nums_signs[i-1] = nums_signs[i-1].(float64) / nums_signs[i+1].(float64)
				nums_signs = removeIndex(nums_signs, i)
				nums_signs = removeIndex(nums_signs, i)
			} else {
				i++
			}
		} else {
			i++
		}
	}

	for i := 0; i < len(nums_signs); {
		if i+1 < len(nums_signs) {
			if nums_signs[i] == "+" {
				nums_signs[i-1] = nums_signs[i-1].(float64) + nums_signs[i+1].(float64)
				nums_signs = removeIndex(nums_signs, i)
				nums_signs = removeIndex(nums_signs, i)
			} else if nums_signs[i] == "-" {
				nums_signs[i-1] = nums_signs[i-1].(float64) - nums_signs[i+1].(float64)
				nums_signs = removeIndex(nums_signs, i)
				nums_signs = removeIndex(nums_signs, i)
			} else {
				i++
			}
		} else {
			i++
		}
	}

	return nums_signs
}

func main() {
	var nums_signs []interface{}
	//var vars map[string]float64

	var closed_par_idx int = -1
	var opened_par_idx int = -1

	fmt.Printf("=> ")
	reader := bufio.NewReader(os.Stdin)
	calc, _ := reader.ReadString('\n')
	calc = strings.TrimSpace(calc)
	calcs := strings.Split(calc, " ")

	for i := 0; i <= len(calcs)-1; i++ {
		float_num, err := strconv.ParseFloat(calcs[i], 64)

		if err == nil {
			nums_signs = append(nums_signs, float_num)
		} else {
			nums_signs = append(nums_signs, calcs[i])
		}
	}
	
	for i := 0; len(nums_signs) != 1; i++ {

		closed_par_idx = -1
		opened_par_idx = -1
		for i := 0; i < len(nums_signs); i++ {
			if nums_signs[i] == "(" {
				opened_par_idx = i
			} else if nums_signs[i] == ")" {
				closed_par_idx = i
				break
			}
		}

		if closed_par_idx < opened_par_idx {
			fmt.Printf("Invalid expression")
			break
		} else if closed_par_idx > 0 && opened_par_idx >= 0 {
			nums_signs = removeIndex(nums_signs, opened_par_idx)
			nums_signs = removeIndex(nums_signs, closed_par_idx - 1)
			nums_signs[opened_par_idx] = solve(nums_signs[opened_par_idx : closed_par_idx - 1])[0]

			for i := opened_par_idx + 1; i < closed_par_idx - 1; i++ {
				nums_signs = removeIndex(nums_signs, opened_par_idx + 1)
			}

		} else {
			nums_signs = solve(nums_signs)
		}
    }

	fmt.Println(nums_signs[0])
}
