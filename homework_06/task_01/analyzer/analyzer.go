package analyzer

import (
	"sort"
	"strings"
	"unicode"
)

// 'WordFrequency' represents a word and its frequency
type WordFrequency struct {
	Word  string
	Count int
}

// 'AnalyzeText' analyzes text and returns a word frequency map
func AnalyzeText(text string) map[string]int {
	wordFrequency := make(map[string]int)

	// Breaking the text into words
	words := extractWords(text)

	// Count the frequency of each word
	for _, word := range words {
		if word != "" {
			// Convert to lowercase
			lowercaseWord := strings.ToLower(word)
			wordFrequency[lowercaseWord]++
		}
	}

	return wordFrequency
}

// 'extractWords' extracts words from text (letters only)
func extractWords(text string) []string {
	var words []string
	var currentWord strings.Builder

	for _, r := range text {
		if unicode.IsLetter(r) {
			currentWord.WriteRune(r)
		} else {
			// If encounter a non-letter, complete the current word
			if currentWord.Len() > 0 {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
		}
	}

	// Add the last word if there is one
	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words
}

// 'GetTopWords' returns the top N most frequent words
func GetTopWords(wordFrequency map[string]int, topN int) []WordFrequency {
	// Convert the map into a slice for sorting
	var frequencies []WordFrequency
	for word, count := range wordFrequency {
		frequencies = append(frequencies, WordFrequency{
			Word:  word,
			Count: count,
		})
	}

	// Sort by descending frequency
	sort.Slice(frequencies, func(i, j int) bool {
		// First by frequency (descending)
		if frequencies[i].Count != frequencies[j].Count {
			return frequencies[i].Count > frequencies[j].Count
		}
		// With the same frequency - alphabetically
		return frequencies[i].Word < frequencies[j].Word
	})

	// Return only the top N
	if topN > len(frequencies) {
		topN = len(frequencies)
	}

	return frequencies[:topN]
}
