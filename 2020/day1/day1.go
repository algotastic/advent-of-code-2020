package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"

	"github.com/algotastic/adventofcode/utils"
)

func main() {
	message := utils.Run("Glory")
	fmt.Println(message)

	file, err := os.Open("input_day1")

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	var matchFound bool
	for i := 0; i < len(nums) && !matchFound; i++ {
		num1 := nums[i]
		fmt.Println(i, num1)
		var num2 int
		for j := 0; i != j && j < len(nums); j++ {
			num2 = nums[j]
			matchFound = 2020 - num2 == num1
			if matchFound {
				break
			}
		}
		if matchFound {
			fmt.Printf("%d + %d = 2020\n", num1, num2)
			fmt.Printf("%d * %d = %d\n", num1, num2, num1 * num2)
			break
		}
	}
}
