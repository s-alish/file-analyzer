package main

import ( 
	"fmt"
	"os"
	"bufio"
	"strings"
)


func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run analyzer.go <filename>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()
 
	var (
		lineCount int
		wordCount int
		charCount int
		uniqueCount = make(map[string]bool)
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		charCount += len(line) + 1 

		words := strings.Fields(line)
		wordCount += len(words)

		for _, word := range words {
			word = strings.ToLower(strings.Trim(word, ".,!?;:"))
			if word != "" {
				uniqueCount[word] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	fmt.Println("┌─────────────────┬────────┐")
	fmt.Printf("│ %-15s │ %6s │\n", "Metric", "Count")
	fmt.Println("├─────────────────┼────────┤")
	fmt.Printf("│ %-15s │ %6d │\n", "Lines", lineCount)
	fmt.Printf("│ %-15s │ %6d │\n", "Words", wordCount)
	fmt.Printf("│ %-15s │ %6d │\n", "Characters", charCount)
	fmt.Printf("│ %-15s │ %6d │\n", "Unique Words", len(uniqueCount))
	fmt.Println("└─────────────────┴────────┘")
}