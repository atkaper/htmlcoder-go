package main

// Written as first try at GoLang for me.
//
// This is a simple command line tool, which you can use to html encode or decode some data.
// You can either use this as a shell (pipe) filter, or you can process a file, or you can
// pass in the data to be encoded/decoded as command line arguments.
//
// Example use:
//
// # show help:
// htmlcoder -h
//
// # decode command line data:
// htmlcoder -d -c -- "&#34;hello&#34;" test
//
// # decode a file:
// htmlcoder -d -f /tmp/somefile.html
//
// # encode output from another command:
// echo '"hello" test' | htmlcoder
//
// Build using:
//
// go build htmlcoder.go
//
// Thijs Kaper, November 6, 2021.

import (
	"bufio"
	"flag"
	"fmt"
	"html"
	"io"
	"os"
	"strings"
)

// Parse Command Line.
// returns: boolean "decode", string "filename"
// decode: true = decode, false = encode.
// filename: "" = use cmdline args, "-" = stdin, other string values = filename.
func parseCommandLine() (bool, string) {
	// Define our command line options, and parse command line.
	decode := flag.Bool("d", false, "decode html entities (default = encode)")
	filename := flag.String("f", "-", "use filename instead of stdin, \"-\" = stdin")
	cmdline := flag.Bool("c", false, "use remainig commandline arguments as input (instead of file or stdin)")
	flag.Parse()

	// Here we check if the used option *combination* does make sense...
	var source string
	switch {
	case (*filename == "-") && !*cmdline && (len(flag.Args()) == 0):
		// "-" indicates to use stdin.
		source = "-"
	case (*filename == "-") && *cmdline && (len(flag.Args()) > 0):
		// clear source to indicate we want command line input.
		source = ""
	case (*filename != "-") && !*cmdline && (len(flag.Args()) == 0):
		// go for a file-name.
		source = *filename
	default:
		// not a supported combination.
		fmt.Println("Illegal flag/argument combination")
		// show cmd-line options (help).
		fmt.Printf("\nUsage of %v:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	return *decode, source
}

// Take the fileName setting, and convert it to the proper "reader".
// This takes either stdin, a file, or the remaining command line arguments.
func openInput(fileName *string) io.Reader {
	var reader io.Reader
	switch {
	case *fileName == "":
		// Take the remaining command line arguments, concatenate them, and convert to a reader.
		reader = strings.NewReader(strings.Join(flag.Args(), " "))
	case *fileName == "-":
		// Just use stdin.
		reader = os.Stdin
	default:
		// Try to open the mentioned file.
		file, err := os.Open(*fileName)
		if err != nil {
			panic(err)
		}
		reader = file
	}
	return reader
}

// Take input reader stream, convert it, and show on stdout.
func convertData(decode *bool, reader *io.Reader) {
	scanner := bufio.NewScanner(*reader)

	for scanner.Scan() {
		var line string
		if *decode {
			line = html.UnescapeString(scanner.Text())
		} else {
			line = html.EscapeString(scanner.Text())
		}
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// Program start.
func main() {
	decode, fileName := parseCommandLine()
	reader := openInput(&fileName)
	convertData(&decode, &reader)
	// Note: not bothering to "close" the input (file) stream here, as program exists anyway.
}
