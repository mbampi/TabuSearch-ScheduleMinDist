package tools

// Tools and Utils package
// Managing tabu list, arrays, file reading

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadFile reads the file from path and returns processingTimes array and numTasks
func ReadFile(filepath string) ([]int, int) {
	// Open file
	log.Println("Reading ", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// Get num tasks
	scanner.Scan()
	numTasks, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err.Error())
	}

	// Read file line by line and populate tasks list
	processingTimes := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		time, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			panic("erro")
		}
		processingTimes = append(processingTimes, int(time))
	}

	return processingTimes, numTasks
}

// Abs returns the absolute value
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Contains returns true if item in array
func Contains(array [][]int, item []int) bool {
	for _, v := range array {
		if Equal(v, item) {
			return true
		}
	}
	return false
}

// Equal returns true if both array are Equal, false otherwise
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Push returns the array with the new item in the first position
func Push(array [][]int, item []int) [][]int {
	return append([][]int{item}, array...)
}

// Pop returns the array without the last item
func Pop(array [][]int) [][]int {
	return array[:len(array)-1]
}

// Max returns the Max value in the array
func Max(array []int) int {
	var max int = -1
	for _, i := range array {
		if i > max {
			max = i
		}
	}
	return max
}

// Min returns the Min value
func Min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// AddToTabuList adds the item in the list, respecting the tabu size (removes the last element if necessary).
func AddToTabuList(tabuList [][]int, tabuSize int, item []int) [][]int {
	if len(tabuList) == tabuSize {
		tabuList = Pop(tabuList)
	}
	tabuList = Push(tabuList, item)
	return tabuList
}
