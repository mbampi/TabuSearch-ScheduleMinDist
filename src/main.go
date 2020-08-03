package main

import (
	"SchedulingMinDist/src/tools"
	"log"
	"math/rand"
	"time"
)

func main() {
	fileNameVariation := []string{"1"}
	for _, nFile := range fileNameVariation {
		log.Println()
		log.Println(" -- Tabu Search on Scheduling with Minimum Distances -- ")

		const maxIter = 1000
		const tabuSize = 30
		const numNeighbors = 10
		filepath := "./input/trsp_50_" + nFile + ".dat"

		// Read processing times and num tasks from file
		processingTimes, numTasks := tools.ReadFile(filepath)
		log.Println("Tasks:", numTasks)
		log.Printf("%v", processingTimes)

		// Print total cumulative time
		var cumTime int = 0
		for _, pTime := range processingTimes {
			cumTime += pTime
		}
		log.Println("Cumulative Time:", cumTime)

		// Solve trivially the problem
		startTimes := greedySolver(processingTimes)
		log.Println("FirstSolution Result", makespan(processingTimes, startTimes))
		// processingTimes, startTimes := trivialSolver(processingTimes)
		// fmt.Println("\nFirstSolution Result", fitness(processingTimes, startTimes))

		// Tabu Search
		start := time.Now()
		tabuSol := tabuSearch(processingTimes, startTimes, tabuSize, numNeighbors, maxIter)
		log.Printf("Time: %s", time.Since(start))
		log.Println("Tabu Result: ", makespan(processingTimes, tabuSol))
	}
}

func tabuSearch(processingTimes, startTimes []int, tabuSize, numNeighbors, maxIter int) []int {
	bestSolIter := 0
	var bestSolution = make([]int, len(startTimes))
	var bestCandidate = make([]int, len(startTimes))
	copy(bestSolution, startTimes)
	copy(bestCandidate, startTimes)

	var tabuList [][]int
	var neighborhood [][]int
	tabuList = tools.AddToTabuList(tabuList, tabuSize, startTimes)

	for numIter := 0; (numIter - bestSolIter) < maxIter; numIter++ {
		neighborhood = getNeighbors(processingTimes, bestCandidate, numNeighbors)
		bestCandidate = neighborhood[0]
		for _, candidate := range neighborhood {
			if (!tools.Contains(tabuList, candidate)) && (makespan(processingTimes, candidate) < makespan(processingTimes, bestCandidate)) {
				bestCandidate = candidate
			}
		}
		if makespan(processingTimes, bestCandidate) < makespan(processingTimes, bestSolution) {
			bestSolution = bestCandidate
			bestSolIter = numIter
			log.Println(bestSolIter)
		}
		tabuList = tools.AddToTabuList(tabuList, tabuSize, bestCandidate)
	}
	return bestSolution
}

func greedySolver(processingTimes []int) []int {
	var minStartList []int
	startTimes := []int{0}
	for i := 1; i < len(processingTimes); i++ {
		currentProcTime := processingTimes[i]

		for j := 0; j < i; j++ {
			previousStart := startTimes[j]
			previousPTime := processingTimes[j]
			minProcTime := tools.Min(currentProcTime, previousPTime)

			minStartList = append(minStartList, previousStart+minProcTime)
		}
		startTimes = append(startTimes, tools.Max(minStartList))
	}
	return startTimes
}

// isSolution returns true if its a possible solution, and false otherwise
func isSolution(processingTimes, startTimes []int) bool {
	for i := 0; i < len(processingTimes); i++ {
		for j := 0; j < len(processingTimes); j++ {
			if i != j {
				if !(tools.Abs(startTimes[i]-startTimes[j]) >= tools.Min(processingTimes[i], processingTimes[j])) {
					return false
				}
			}
		}
	}
	return true
}

// getNeighbors generate <numNeighbors> neighbors using createNeighbor and returns an array of solutions
func getNeighbors(processingTimes, startTimes []int, numNeighbors int) [][]int {
	var neighbor []int
	var neighborhood [][]int
	for i := 0; i < numNeighbors; i++ {
		for {
			neighbor = createNeighbor(startTimes)
			if isSolution(processingTimes, neighbor) {
				break
			}
		}
		neighborhood = append(neighborhood, neighbor)
	}
	return neighborhood
}

// createNeighbor swaps one start time from startTimes array on random positions
func createNeighbor(startTimes []int) []int {
	var neighbor = make([]int, len(startTimes))
	copy(neighbor, startTimes)

	randPos := rand.Intn(len(startTimes))
	randPos2 := rand.Intn(len(startTimes))
	// It does not assure to be different, but chances are minimal
	neighbor[randPos], neighbor[randPos2] = neighbor[randPos2], neighbor[randPos]
	return neighbor
}

// makespan returns the makespan. Max(startTime + processingTime)
func makespan(processingTimes, startTimes []int) int {
	if !isSolution(processingTimes, startTimes) {
		const MaxInt = int(^uint(0) >> 1)
		return MaxInt
	}
	var maxMakespan int = -1
	var makespan int
	for i := range processingTimes {
		makespan = startTimes[i] + processingTimes[i]
		if makespan > maxMakespan {
			maxMakespan = makespan
		}
	}
	return maxMakespan
}
