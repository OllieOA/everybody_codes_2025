package quest05

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

type Sword struct {
	Id int
	Fishbone []map[string]int
	Quality int
	PerLevelValues []int
}

func parseInput(content string) (int, []int) {
	idAndSequence := strings.Split(content, ":")
	identifier, _ := strconv.Atoi(idAndSequence[0])
	sequenceStr := idAndSequence[1]

	sequence, _ := common.ParseDelimtedList(sequenceStr, strconv.Atoi, ",")

	return identifier, sequence
}

func makeFishbone(sequence []int) []map[string]int {
	spineTemplate := map[string]int{"c": -1, "l": -1, "r": -1}
	fb := []map[string]int{}

	for _, val := range sequence {
		// initialise
		if len(fb) == 0 {
			newSeg := common.CloneMap(spineTemplate)
			newSeg["c"] = val
			fb = append(fb, newSeg)
			continue
		}

		valAdded := false
		for _, seg := range fb {
			if (val < seg["c"]) && (seg["l"] == -1) {
				seg["l"] = val
				valAdded = true
				break
			}
			
			if (val > seg["c"]) && (seg["r"] == -1) {
				seg["r"] = val
				valAdded = true
				break
			}
		}
		if !valAdded {
			newSeg := common.CloneMap(spineTemplate)
			newSeg["c"] = val
			fb = append(fb, newSeg)
		}
		// fmt.Println("After loop adding", val, "fb is\n", visualiseFishbone(fb))
	}
	return fb
}

func visualiseFishbone(fishbone []map[string]int) string {
	fishboneStr := "\n"
	for _, spine := range fishbone {
		spineStr := ""
		
		if spine["l"] == -1 {
			spineStr += "  "
		} else {
			spineStr += strconv.Itoa(spine["l"]) + "-"
		}
		
		spineStr += strconv.Itoa(spine["c"])

		if spine["r"] == -1 {
			spineStr += "  "
		} else {
			spineStr += "-" + strconv.Itoa(spine["r"])
		}

		fishboneStr += spineStr + "\n"
	}
	return fishboneStr
}

func getQuality(fishbone []map[string]int) int {
	qualityStr := ""
	for _, spine := range fishbone {
		qualityStr += strconv.Itoa(spine["c"])
	}

	quality, _ := strconv.Atoi(qualityStr)

	return quality
}

func getPerLevelValues(fishbone []map[string]int) []int {
	perLevelValues := []int{}
	for _, spine := range fishbone {
		spineValStr := ""
		if spine["l"] != -1 {
			spineValStr += strconv.Itoa(spine["l"])
		}

		spineValStr += strconv.Itoa(spine["c"])

		if spine["r"] != -1 {
			spineValStr += strconv.Itoa(spine["r"])
		}

		spineVal, _ := strconv.Atoi(spineValStr)
		perLevelValues = append(perLevelValues, spineVal)
	}
	return perLevelValues
}

func makeSwordSpecs(content string) []Sword {
	swordSpecsStr := strings.Split(content, "\n")
	swordSpecs := []Sword{}

	for _, swordSpec := range swordSpecsStr {
		identifier, sequence := parseInput(swordSpec)
		
		fishbone := makeFishbone(sequence)
		quality := getQuality(fishbone)
		perLevelValues := getPerLevelValues(fishbone)

		newSword := Sword{
			Id: identifier,
			Fishbone: fishbone,
			Quality: quality,
			PerLevelValues: perLevelValues,
		}

		swordSpecs = append(swordSpecs, newSword)
	}
	return swordSpecs
}

// Logical answers to the parts

func part1(context common.Context) int {
	part := 1
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	identifier, sequence := parseInput(content)
	_ = identifier

	fishbone := makeFishbone(sequence)
	quality := getQuality(fishbone)
	
	return quality
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	swordSpecs := makeSwordSpecs(content)

	minQuality := 999999999999999999
	maxQuality := 0

	for _, sword := range swordSpecs {
		minQuality = min(minQuality, sword.Quality)
		maxQuality = max(maxQuality, sword.Quality)
	}
	return maxQuality - minQuality
}

func SortSwords(s []Sword) {
	sort.Slice(s, func(i, j int) bool {
		a, b := s[i], s[j]
		if a.Quality != b.Quality {
			return a.Quality > b.Quality
		}
		
		for i := range min(len(a.PerLevelValues), len(b.PerLevelValues)) {
			if a.PerLevelValues[i] != b.PerLevelValues[i] {
				return a.PerLevelValues[i] > b.PerLevelValues[i]
			}
		}
		return a.Id > b.Id 
	})
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	swordSpecs := makeSwordSpecs(content)
	SortSwords(swordSpecs)

	checksum := 0
	for i, swordSpec := range swordSpecs {
		checksum += (i + 1) * swordSpec.Id
	}
	return checksum
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
