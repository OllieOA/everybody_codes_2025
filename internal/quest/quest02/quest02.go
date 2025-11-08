package quest02

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/OllieOA/everybody_codes_2025/internal/common"
)

// Boilerplate

var questNumber int

func init() {
	// Automatically determine quest number from package name
	questNumber = common.GetQuestNumber()
	common.Register(questNumber, fmt.Sprintf("Quest %02d placeholder", questNumber), Run)
}

// Shared helper functions
func getA(content string) ([2]int, error) {
	re := regexp.MustCompile(`A=\[(-?\d+),(-?\d+)\]`)
	matches := re.FindStringSubmatch(content)
	if len(matches) < 3 {
		return [2]int{-1, -1}, fmt.Errorf("failed to parse A from content")
	}
	A1, err1 := strconv.Atoi(matches[1])
	A2, err2 := strconv.Atoi(matches[2])
	if err1 != nil || err2 != nil {
		return [2]int{-1, -1}, fmt.Errorf("failed to convert A values to integers")
	}
	return [2]int{A1, A2}, nil
}

func addComplexNum(C1 [2]int, C2 [2]int) [2]int {
	return [2]int{C1[0] + C2[0], C1[1] + C2[1]}
}

func multComplexNum(C1 [2]int, C2 [2]int) [2]int {
	return [2]int{C1[0]*C2[0] - C1[1]*C2[1], C1[0]*C2[1] + C1[1]*C2[0]}
}

func divComplexNum(C1 [2]int, C2 [2]int) [2]int {
	return [2]int{int(C1[0]/C2[0]), int(C1[1]/C2[1])}
}

func sequenceStep(N [2]int, M [2]int, size int) [2]int {
	N = multComplexNum(N, N)
	N = divComplexNum(N, [2]int{size, size})
	N = addComplexNum(N, M)
	return N
}

func runAlgorithmForPoint(P [2]int, iterations int, threshold float64) bool {
	currNum := [2]int{0, 0}
	for i := 0; i < iterations; i++ {
		currNum = sequenceStep(currNum, P, 100000)
		if math.Abs(float64(currNum[0])) > threshold || math.Abs(float64(currNum[1])) > threshold {
			return false
		}
	}
	return true
}

func countEngravedPoints(points map[[2]int]bool) int {
	count := 0
	for _, isTrue := range points {
		if isTrue {
			count++
		}
	}
	return count
}

func writeMapToImage(points map[[2]int]bool, filename string) error {
	// image is known to be square so we can just take square root
	imgSize := int(math.Sqrt(float64(len(points))))

	minX := 10000000
	minY := 10000000

	maxX := -10000000
	maxY := -10000000


	for point := range points {
		if point[0] < minX {
			minX = point[0]
		}
		if point[1] < minY {
			minY = point[1]
		}
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[1] > maxY {
			maxY = point[1]
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	// Iterate over points and write the pixels
	for point, isEngraved := range points {
		// We need to normalize the coordinates to start from (0,0)
		normX := point[0] - minX
		normY := point[1] - minY

		// Scale to image size
		imgX := int(float64(normX) / float64(maxX-minX) * float64(imgSize-1))
		imgY := int(float64(normY) / float64(maxY-minY) * float64(imgSize-1))

		if isEngraved {
			img.Set(imgX, imgY, image.White)
		} else {
			img.Set(imgX, imgY, image.Black)
		}
	}

	// Save the image to file
	outFile, err := os.Create("visualisations/" + filename)
	if err != nil {
		return fmt.Errorf("failed to create image file: %w", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		return fmt.Errorf("failed to encode image to PNG: %w", err)
	}
	return nil
}
	


// Logical answers to the parts

func part1(context common.Context) string {
	part := 1
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return ""
	}

	// Logic below
	
	// Initialise some 'complex numbers'
	currNum := [2]int{0, 0}
	A, err := getA(content)
	if err != nil {
		fmt.Printf("Error parsing A: %s\n", err)
		return ""
	}

	for range 3 {
		currNum = sequenceStep(currNum, A, 10)
	}

	return fmt.Sprintf("[%d,%d]", currNum[0], currNum[1])
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	gridSize := 100
	iterations := 100
	pointThreshold := 1000000.0

	A, err := getA(content)
	if err != nil {
		fmt.Printf("Error parsing A: %s\n", err)
		return 0
	}

	finalCoord := addComplexNum([2]int{1000,1000}, A)
	coords := make(map[[2]int]bool)
	stepSize := int((finalCoord[0] - A[0]) / gridSize)

	for x := A[0]; x <= finalCoord[0]; x += stepSize {
		for y := A[1]; y <= finalCoord[1]; y += stepSize {
			point := [2]int{x, y}
			coords[point] = runAlgorithmForPoint(point, iterations, pointThreshold)
		}
	}

	count := countEngravedPoints(coords)

	// Optionally write the map to an image for visualization
	err = writeMapToImage(coords, "quest02_part2_map.png")
	if err != nil {
		fmt.Printf("Error writing map to image: %s\n", err)
	}
	return count
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	gridSize := 1000
	iterations := 100
	pointThreshold := 1000000.0

	A, err := getA(content)
	if err != nil {
		fmt.Printf("Error parsing A: %s\n", err)
		return 0
	}

	finalCoord := addComplexNum([2]int{1000,1000}, A)
	coords := make(map[[2]int]bool)
	stepSize := int((finalCoord[0] - A[0]) / gridSize)
	
	// Note - I used copilot here to generate the concurrent map population code
	// but I will add the explanations for my own understanding.
	var mu sync.Mutex  // This is our data gatekeeper on the result structure
	var wg sync.WaitGroup  // This is our "queue" kind of

	for x := A[0]; x <= finalCoord[0]; x += stepSize {
		for y := A[1]; y <= finalCoord[1]; y += stepSize {
			// We are in a typical loop for populating the map
			wg.Add(1) // This effectively just adds a counter - the specific goroutine doesn't matter
			go func(point [2]int) {  // Here, "go" is a keyword to say "run this in the background" basically
				defer wg.Done() // This is like saying "remember to mark this task as done when finished" - if it was at the end, it wouldn't be executed if the algorithm failed
				result := runAlgorithmForPoint(point, iterations, pointThreshold)
				mu.Lock() // Lock the mutex to ensure exclusive access to the map. This is shared across all of the goroutines
				coords[point] = result
				mu.Unlock()  // Unlock the mutex to allow other goroutines to access the map - we're done with it
			}([2]int{x, y}) // This is the "point" argument passed to the goroutine
		}
	}
	
	wg.Wait() // This blocks until all goroutines have called wg.Done(). Basically it just blocks the main thread

	count := countEngravedPoints(coords)

	// Optionally write the map to an image for visualization
	err = writeMapToImage(coords, "quest02_part3_map.png")
	if err != nil {
		fmt.Printf("Error writing map to image: %s\n", err)
	}
	return count
}

// Finally, main runner function

func Run(ctx common.Context) error {
	base_start := time.Now()

	part1_answer := part1(ctx)
	part1_time := time.Since(base_start)
	fmt.Printf("Part 1 answer: %s (computed in %s)\n", part1_answer, part1_time)

	part2_answer := part2(ctx)
	part2_time := time.Since(base_start) - part1_time
	fmt.Printf("Part 2 answer: %d (computed in %s)\n", part2_answer, part2_time)

	part3_answer := part3(ctx)
	part3_time := time.Since(base_start) - part1_time - part2_time
	fmt.Printf("Part 3 answer: %d (computed in %s)\n", part3_answer, part3_time)

	total_time := time.Since(base_start)
	fmt.Printf("Total time: %s\n", total_time)

	return nil
}
