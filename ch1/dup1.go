// dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }

    // Note: ignoring potential errors from input.Err()
    if len(counts) > 0 {
        fmt.Printf("\n%s\t%s\n", "Count", "Value")
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}