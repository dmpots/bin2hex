package main

import "fmt"
import "flag"
import "os"
import "io/ioutil"

type opts struct {
	arrayName      string
	inputFile      string
	outputFileName *string
}

func parseOptions() opts {
	args := opts{}
	//arrayName := flag.String("array-name", "bin", "array name for the c array")
	flag.StringVar(&args.arrayName, "array-name", "bin", "array name for the c array")
	outputPath := flag.String("output", "", "output path")

	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "usage: bin2hex [flags] input-file")
		flag.PrintDefaults()
		os.Exit(1)
	}
	args.inputFile = flag.Args()[0]

	if *outputPath != "" {
		args.outputFileName = outputPath
	}

	return args
}

func main() {
	options := parseOptions()

	bytes, err := ioutil.ReadFile(options.inputFile)
	if err != nil {
		panic(err)
	}

	output := os.Stdout
	if options.outputFileName != nil {
		o, err := os.Create(*options.outputFileName)
		if err != nil {
			panic(err)
		}
		output = o
	}

	fmt.Fprintf(output, "char %s[] = {\n", options.arrayName)
	for _, b := range bytes {
		fmt.Fprintf(output, "  0x%X,\n", b)
	}
	fmt.Fprintf(output, "}\n")
}
