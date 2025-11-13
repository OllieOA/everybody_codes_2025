package quest08

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func getOppositeNailMap(numNails int) map[int]int {
	oppositeNails := map[int]int{}
	for nail := numNails; nail > numNails / 2; nail-- {
		oppositeNails[nail] = nail - (numNails/2)
	}
	return oppositeNails
}

// Logical answers to the parts

func part1(context common.Context) int {
	part := 1
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	numNails := 32
	if context.Sample {
		numNails = 8
	}
	// Logic below

	nailMap := getOppositeNailMap(numNails)
	centerPasses := 0
	threadOrder, _ := common.ParseDelimitedList(content, strconv.Atoi, ",")

	for i := 1; i < len(threadOrder); i++ {
		prevNail := threadOrder[i-1]
		currNail := threadOrder[i]
		if nailMap[max(prevNail, currNail)] == min(prevNail, currNail) {
			centerPasses += 1
		}
	}

	return centerPasses
}

var arcPaths = map[string]map[string][]int{}
func getArcs(start int, end int, numNails int) map[string][]int {
	// I thought so hard about how to handle the wrapping that I was paralyzed
	// and just decided to do it in the dumbest way possible to move on.
	
	// There is almost certainly a simple way to make these arcs with modulo
	// but I just don't care enough to figure it out right now
	stringPair := fmt.Sprintf("%d,%d", start, end)

	if out, ok := arcPaths[stringPair]; ok {
		return out
	}

	negStart := start-1
	if negStart == 0 {
		negStart = numNails
	}

	negSteps := []int{negStart}
	if negStart == end { // Special case when we are a neighbour
		negSteps = []int{}
	} else {
		for {
			nextStep := negSteps[len(negSteps)-1] - 1
			if nextStep == 0 {
				nextStep = numNails
			}
			if nextStep == end {
				break
			}
			negSteps = append(negSteps, nextStep)
		}
	}

	posStart := start+1
	if posStart > numNails {
		posStart = 1
	}
	posSteps := []int{posStart}
	if posStart == end { // Special case when we are a neighbour
		posSteps = []int{}
	} else {
		for {
			nextStep := posSteps[len(posSteps)-1] + 1
			if nextStep > numNails {
				nextStep = 1
			}
			if nextStep == end {
				break
			}
			posSteps = append(posSteps, nextStep)
		}
	}

	var shortArc []int
	var longArc []int

	if len(posSteps) > len(negSteps) {
		shortArc = negSteps
		longArc = posSteps
	} else {
		shortArc = posSteps
		longArc = negSteps
	}
	outMap := map[string][]int{"short": shortArc, "long": longArc}
	arcPaths[stringPair] = outMap

	return outMap
}

func makeBaseConnectionMap(numNails int) map[int]int {
	baseConnectionMap := make(map[int]int, numNails)
	for i := 1; i <= numNails; i++ {
		baseConnectionMap[i] = 0
	}
	return baseConnectionMap
}

func getKnotCount(content string, numNails int) (int, map[int]map[int]int) {
	// This is a map of how many connections to other nails a given nail has
	nailThreadConnections := make(map[int]map[int]int, numNails)
	for i := 1; i <= numNails; i++ { // We basically create a map to all other nails for each nail
		nailThreadConnections[i] = makeBaseConnectionMap(numNails)
	}

	threadOrder, _ := common.ParseDelimitedList(content, strconv.Atoi, ",")
	knotCount := 0

	for i := 1; i < len(threadOrder); i++ {
		prevNail := threadOrder[i-1]
		currNail := threadOrder[i]
		// Find the arcs between prevNail and currNail noting it wraps 
		// back around
		arcs := getArcs(prevNail, currNail, numNails)
		shortArc := arcs["short"]
		longArc := arcs["long"]

		// Now we will step through each nail in the short arc and see how many
		// connections it has to members of the long arc. This count is
		// maintained in nailThreadConnections
		for _, shortArcNail := range shortArc {
			for _, longArcNail := range longArc {
				knotCount += nailThreadConnections[shortArcNail][longArcNail]
			}
		}

		// Finally, we log the new nail connections
		nailThreadConnections[prevNail][currNail] += 1
		nailThreadConnections[currNail][prevNail] += 1
	}
	return knotCount, nailThreadConnections
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	numNails := 256
	if context.Sample {
		numNails = 8
	}
	// Logic below
	knotCount, _ := getKnotCount(content, numNails)
	
	return knotCount
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	numNails := 256
	if context.Sample {
		numNails = 8
	}
	// Logic below
	_, nailThreadConnections := getKnotCount(content, numNails)

	saveOutMap := make(map[string]map[string]int, len(nailThreadConnections))

	for k, v := range nailThreadConnections {
		saveOutMap[strconv.Itoa(k)] = common.MakeIntMapJSONSerialisable(v)
	}

	jsonData, _ := json.Marshal(saveOutMap)
	os.WriteFile("visualisations/nail_connections.json", jsonData, 0644)

	maxThreadsCut := 0
	for n := 1; n < numNails; n++ {
		for o := 1; o < numNails; o++ {
			if n == o {
				continue
			}
			arcs := getArcs(n, o, numNails)
			shortArc := arcs["short"]
			longArc := arcs["long"]
			crossings := 0

			// Add the connection between n and o if it exists - this counts if connected
			crossings += nailThreadConnections[n][o]

			for _, shortArcNail := range shortArc {
				for _, longArcNail := range longArc {
					crossings += nailThreadConnections[shortArcNail][longArcNail]
				}
			}
			maxThreadsCut = max(maxThreadsCut, crossings)
		}
	}
	return maxThreadsCut
}

// Finally, main runner function

func Run(ctx common.Context) error {
	base_start := time.Now()

	part1_answer := part1(ctx)
	part1_time := time.Since(base_start)
	fmt.Printf("Part 1 answer: %d (computed in %s)\n", part1_answer, part1_time)

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