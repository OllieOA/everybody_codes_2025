package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/OllieOA/everybody_codes_2025/internal/common"

	// each quest package must be imported to register its handler with init()
	// Uncommented as I complete them
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest00"
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest01"
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest02"
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest03"
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest04"
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest05"
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest06"
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest07"
	_ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest08"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest09"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest10"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest11"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest12"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest13"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest14"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest15"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest16"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest17"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest18"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest19"
	// _ "github.com/OllieOA/everybody_codes_2025/internal/quest/quest20"
)

func main() {
	var sampleFlag bool
	flag.BoolVar(&sampleFlag, "s", false, "sample flag")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: ec [-s] <day 1-20>")
		os.Exit(2)
	}

	questNumber, err := strconv.Atoi(args[0])
	if err != nil || questNumber < 0 || questNumber > 20 {
		fmt.Fprintln(os.Stderr, "error: quest must be an integer between 0 and 20 (where 0 is a placeholder quest for testing)")
		os.Exit(2)
	}

	handler := common.GetHandler(questNumber)
	if handler == nil {
		fmt.Fprintf(os.Stderr, "error: quest %d is not yet implemented (or you have not imported it in main.go)\n", questNumber)
		os.Exit(1)
	}

	ctx := common.Context{
		Quest:  questNumber,
		Sample: sampleFlag,
	}

	if err := handler(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error running quest %d: %v\n", questNumber, err)
		os.Exit(1)
	}

}