package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func BLUE_TEXT()  { fmt.Printf("\033[1m\033[34m") }
func RED_TEXT()   { fmt.Printf("\033[1m\033[31m") }
func RESET_TEXT() { fmt.Printf("\033[0m") }

func do_calc(nums_signs []interface{}) []interface{} {
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
			} else if nums_signs[i] == "%" {
				nums_signs[i-1] = int(nums_signs[i-1].(float64)) % int(nums_signs[i+1].(float64))
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

func removeIndex(array []interface{}, index int) []interface{} {
	var new_array []interface{}

	for i := 0; i < len(array); i++ {
		if i != index {
			new_array = append(new_array, array[i])
		}
	}

	return new_array
}

func solve(nums_signs []interface{}) float64 {
	for i := 0; len(nums_signs) != 1; i++ {

		closed_par_idx := -1
		opened_par_idx := -1
		for i := 0; i < len(nums_signs); i++ {
			if nums_signs[i] == "(" {
				opened_par_idx = i
			} else if nums_signs[i] == ")" {
				closed_par_idx = i
				break
			}
		}

		if closed_par_idx < opened_par_idx || (closed_par_idx >= 0 && opened_par_idx < 0) {
			RED_TEXT()
			fmt.Printf("Invalid expression => ")
			if closed_par_idx < 0 {
				fmt.Printf("'(' not closed")
			} else {
				fmt.Printf("')' not opened")
			}
			RESET_TEXT()
			return 0

		} else if closed_par_idx > 0 && opened_par_idx >= 0 {
			nums_signs = removeIndex(nums_signs, opened_par_idx)
			nums_signs = removeIndex(nums_signs, closed_par_idx-1)
			nums_signs[opened_par_idx] = do_calc(nums_signs[opened_par_idx : closed_par_idx-1])[0]

			for i := opened_par_idx + 1; i < closed_par_idx-1; i++ {
				nums_signs = removeIndex(nums_signs, opened_par_idx+1)
			}

		} else {
			nums_signs = do_calc(nums_signs)

		}
	}
	fmt.Println(nums_signs[0])
	return nums_signs[0].(float64)
}

func main() {
	vars := make(map[string]float64)

	fmt.Printf("Type -help to list informations, -exit to close the execution")
	for {
		var nums_signs []interface{}
		BLUE_TEXT()
		fmt.Printf("\n=> ")
		RESET_TEXT()
		reader := bufio.NewReader(os.Stdin)
		calc, _ := reader.ReadString('\n')
		calc = strings.TrimSpace(calc)
		calcs := strings.Split(calc, " ")

		for i := 0; i <= len(calcs)-1; i++ {
			float_num, err := strconv.ParseFloat(calcs[i], 64)
			value, ok := vars[calcs[i]]

			if err == nil {
				nums_signs = append(nums_signs, float_num)
			} else if ok {
				nums_signs = append(nums_signs, value)
			} else {
				nums_signs = append(nums_signs, calcs[i])
			}
		}

		if len(nums_signs) == 1 {
			if nums_signs[0] == "-exit" {
				RED_TEXT()
				fmt.Println("Exiting...")
				RESET_TEXT()
				os.Exit(0)

			} else if nums_signs[0] == "-help" {
				fmt.Println("Type any mathematical expression you want using the +, -, *, /, % and ^")
				fmt.Println("Use the order you want by adding parenthesis")
				fmt.Println("Save values for future use with this syntax: variable_name = value")
				fmt.Println("Ex: 5 + ( 4 * 3 ) + 2 ^ ( 2 - 3 ) ")
			} else {
				fmt.Println(nums_signs[0])
			}

		} else {
			if nums_signs[1] == "=" {
				if len(nums_signs) > 1 {
					vars[nums_signs[0].(string)] = solve(nums_signs[2:len(nums_signs)])
				} else {
					vars[nums_signs[0].(string)] = 0
				}

			} else {
				solve(nums_signs)
			}
		}
	}
}
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func BLUE_TEXT()  { fmt.Printf("\033[1m\033[34m") }
func RED_TEXT()   { fmt.Printf("\033[1m\033[31m") }
func RESET_TEXT() { fmt.Printf("\033[0m") }

func do_calc(nums_signs []interface{}) []interface{} {
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
			} else if nums_signs[i] == "%" {
				nums_signs[i-1] = int(nums_signs[i-1].(float64)) % int(nums_signs[i+1].(float64))
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

func removeIndex(array []interface{}, index int) []interface{} {
	var new_array []interface{}

	for i := 0; i < len(array); i++ {
		if i != index {
			new_array = append(new_array, array[i])
		}
	}

	return new_array
}

func solve(nums_signs []interface{}) float64 {
	for i := 0; len(nums_signs) != 1; i++ {

		closed_par_idx := -1
		opened_par_idx := -1
		for i := 0; i < len(nums_signs); i++ {
			if nums_signs[i] == "(" {
				opened_par_idx = i
			} else if nums_signs[i] == ")" {
				closed_par_idx = i
				break
			}
		}

		if closed_par_idx < opened_par_idx || (closed_par_idx >= 0 && opened_par_idx < 0) {
			RED_TEXT()
			fmt.Printf("Invalid expression => ")
			if closed_par_idx < 0 {
				fmt.Printf("'(' not closed")
			} else {
				fmt.Printf("')' not opened")
			}
			RESET_TEXT()
			return 0

		} else if closed_par_idx > 0 && opened_par_idx >= 0 {
			nums_signs = removeIndex(nums_signs, opened_par_idx)
			nums_signs = removeIndex(nums_signs, closed_par_idx-1)
			nums_signs[opened_par_idx] = do_calc(nums_signs[opened_par_idx : closed_par_idx-1])[0]

			for i := opened_par_idx + 1; i < closed_par_idx-1; i++ {
				nums_signs = removeIndex(nums_signs, opened_par_idx+1)
			}

		} else {
			nums_signs = do_calc(nums_signs)

		}
	}
	fmt.Println(nums_signs[0])
	return nums_signs[0].(float64)
}

func main() {
	vars := make(map[string]float64)

	fmt.Printf("Type -help to list informations, -exit to close the execution")
	for {
		var nums_signs []interface{}
		BLUE_TEXT()
		fmt.Printf("\n=> ")
		RESET_TEXT()
		reader := bufio.NewReader(os.Stdin)
		calc, _ := reader.ReadString('\n')
		calc = strings.TrimSpace(calc)
		calcs := strings.Split(calc, " ")

		for i := 0; i <= len(calcs)-1; i++ {
			float_num, err := strconv.ParseFloat(calcs[i], 64)
			value, ok := vars[calcs[i]]

			if err == nil {
				nums_signs = append(nums_signs, float_num)
			} else if ok {
				nums_signs = append(nums_signs, value)
			} else {
				nums_signs = append(nums_signs, calcs[i])
			}
		}

		if len(nums_signs) == 1 {
			if nums_signs[0] == "-exit" {
				RED_TEXT()
				fmt.Println("Exiting...")
				RESET_TEXT()
				os.Exit(0)

			} else if nums_signs[0] == "-help" {
				fmt.Println("Type any mathematical expression you want using the +, -, *, /, % and ^")
				fmt.Println("Use the order you want by adding parenthesis")
				fmt.Println("Save values for future use with this syntax: variable_name = value")
				fmt.Println("Ex: 5 + ( 4 * 3 ) + 2 ^ ( 2 - 3 ) ")
			} else {
				fmt.Println(nums_signs[0])
			}

		} else {
			if nums_signs[1] == "=" {
				if len(nums_signs) > 1 {
					vars[nums_signs[0].(string)] = solve(nums_signs[2:len(nums_signs)])
				} else {
					vars[nums_signs[0].(string)] = 0
				}

			} else {
				solve(nums_signs)
			}
		}
	}
}
