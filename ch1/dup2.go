package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]map[string]int)
    fileNames := make(map[string]bool)
    files := os.Args[1:]

    if len(files) == 0 {
        countLines(os.Stdin, counts, fileNames)
    } else {
        for _, arg := range files {
            //fmt.Println(arg)
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts, fileNames)
            f.Close()
        }
    }

    for fnMap, maps := range counts {
        for _, occurrences  := range maps {
            if occurrences > 1 {
                fileNames[fnMap] = true
            }
        }

    }
    for fn, isTrue := range fileNames {
        if isTrue == true {
           fmt.Printf("%s\n", fn)
        }
    }
}

func countLines(f *os.File, counts map[string]map[string]int, fileNames map[string]bool) {
    input := bufio.NewScanner(f)
    counts[f.Name()] = make(map[string]int)
    for input.Scan() {
        counts[f.Name()][input.Text()]++
    }
}