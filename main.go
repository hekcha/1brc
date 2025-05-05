package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
	"math"
)

func min(a, b float32) float32 {
	if a == 0 || b < a {
		return b
	}
	return a
}

func max(a, b float32) float32 {
	if b > a {
		return b
	}
	return a
}



func main() {
	startTime := time.Now()

	file, err := os.Open("measurements.txt");
	if err != nil {
		fmt.Println("Not able to read the file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum = make(map[string]float32)
	var minimum = make(map[string]float32)
	var maximum = make(map[string]float32)
	var count = make(map[string]int)

	for scanner.Scan(){
		readLine := scanner.Text()

		splitString := strings.Split(readLine, ";")

		temp, err := strconv.ParseFloat(splitString[1], 32)

		if err != nil {
			fmt.Println("Error in string parsing")
			return
		}

		temp32 := float32(temp)

		sum[splitString[0]] += temp32
		minimum[splitString[0]] = min(minimum[splitString[0]], temp32)
		maximum[splitString[0]] = max(maximum[splitString[0]], temp32)
		count[splitString[0]] += 1
	}

	for key, _ := range sum {
		avg := sum[key] / float32(count[key])
		roundedAvg := float32(math.Ceil(float64(avg)*100) / 100) // rounding up to 2 decimal places
		fmt.Printf("%s: %.2f / %.2f / %.2f;\n", key, minimum[key], roundedAvg, maximum[key])	}


	fmt.Println("Total time taken" , time.Since(startTime))




}
