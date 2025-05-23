package main

import (
    "bufio"
    "crypto/rand"
    "fmt"
    "math/big"
    "os"
)

func main() {
    // Read all lines from stdin
    var lines []string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading input:", err)
        os.Exit(1)
    }
    
    // Check if we have any lines
    if len(lines) == 0 {
        fmt.Fprintln(os.Stderr, "No input provided")
        os.Exit(1)
    }
    
    // Pick one random line using crypto/rand
    max := big.NewInt(int64(len(lines)))
    n, err := rand.Int(rand.Reader, max)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error generating random number:", err)
        os.Exit(1)
    }
    
    // Output the chosen line
    fmt.Println(lines[n.Int64()])
}