/*
 * Application:     Character Freqency Counter
 * File:            freq_counter.go
 * Description:     Source file for "freq_counter" module
 * Language:        go
 * Dev Env:         Arch Linux 64-bit
 *
 * Author:          Kyle Thomas
 * Date Started:    January 22, 2019
 */

package main

import (
	"fmt"       // printing
	"io/ioutil" // reading file
	"os"        // exiting for errors
	"sort"      // sorting results
)

// freq struct stores the unicode character & number of times its been seen
type freq struct {
	char rune
	freq int
}

// countFreq looks through a string and adds each charcter into a map with the
// 		value being the # of times the character has been seen. It returns the map
func countFreq(s string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, c := range s {
		freqMap[c] += 1
	}
	return freqMap
}

// printFreq calls sortFreq on the map of chars and then prints it to the screen
//		in a pretty format
func printFreq(m map[rune]int) {
	var (
		totalC, uniqueC int
	)
	f := sortFreq(m)
	fmt.Println("Char\tFreq")
	for _, s := range f {
		totalC += s.freq
		uniqueC++
		fmt.Printf("%q\t%d\n", s.char, s.freq)
	}
	fmt.Println("Unique chars:\t", uniqueC)
	fmt.Println("Total chars:\t", totalC)
}

// readFile takes in a filename and returns the contents as a string
// if the file cannot be read the program will exit
func readFile(fname string) string {
	contents, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("Fatal: cannot open file")
		os.Exit(1)
	}
	return string(contents)
}

// sortFreq converts the map of chars into a slice of struct freq and sorts
// 		in descending order based on the frequency of the character
func sortFreq(m map[rune]int) []freq {
	var (
		f    []freq
		tmpF *freq
	)
	for key, val := range m {
		tmpF = new(freq)
		tmpF.char = key
		tmpF.freq = val
		f = append(f, *tmpF)
	}
	sort.Slice(f, func(i, j int) bool { return f[i].freq > f[j].freq })
	return f
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fatal: incorrect arguments.")
		fmt.Println("usage: $ freq-counter file-to-analyze")
		os.Exit(1)
	}
	printFreq(countFreq(readFile(os.Args[1])))
	os.Exit(0)
}
