/*
Разработайте консольную утилиту для анализа текстовых файлов, содержащих литературные произведения.
Утилита должна подсчитывать частоту встречаемости слов и выводить наиболее часто употребляемые слова в порядке убывания частоты.
Регистр букв при подсчёте не учитывается (например, "The" и "the" считаются одним и тем же словом).
Под словом подразумевается последовательность букв, а цифры и символы пунктуации не учитываются.
Поддерживаемые флаги командной строки:
1) -file или -f — обязательный флаг, указывающий путь к текстовому файлу.
Если флаг отсутствует или указанный файл не существует, программа должна завершиться с сообщением об ошибке.
2) -top или -t — необязательный флаг, определяющий количество слов, выводимых в результатах.
По умолчанию должно выводиться 10 наиболее частотных слов.
*/

/*
Запуск:
go run task_01.go -file Evgeniy_Onegin.txt
go run task_01.go -f Evgeniy_Onegin.txt

go run task_01.go -file Evgeniy_Onegin.txt -top 20
go run task_01.go -f Evgeniy_Onegin.txt -t 5
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"task_01/analyzer"
	"task_01/fileutils"
)

func main() {
	// Defining command line flags
	filePath := flag.String("file", "", "Path to the text file to analyze")
	filePathShort := flag.String("f", "", "Path to the text file to analyze (short)")
	topCount := flag.Int("top", 10, "Number of most frequent words to display")
	topCountShort := flag.Int("t", 10, "Number of most frequent words to display (short)")

	flag.Parse()

	// Determine which flag was used
	finalFilePath := *filePath
	if finalFilePath == "" {
		finalFilePath = *filePathShort
	}

	finalTopCount := *topCount
	if flag.Lookup("t").Value.String() != "10" {
		finalTopCount = *topCountShort
	}

	// Check if the file is specified
	if finalFilePath == "" {
		fmt.Fprintln(os.Stderr, "Error: file path is required. Use -file or -f flag.")
		flag.Usage()
		os.Exit(1)
	}

	// Reading the contents of the file
	content, err := fileutils.ReadFile(finalFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Analyzing the text
	wordFrequency := analyzer.AnalyzeText(content)

	// We get the top words
	topWords := analyzer.GetTopWords(wordFrequency, finalTopCount)

	// We display the results
	fmt.Printf("Top %d most frequent words:\n", finalTopCount)
	fmt.Println("-----------------------------")
	for i, word := range topWords {
		fmt.Printf("%d. %s - %d occurrences\n", i+1, word.Word, word.Count)
	}
}
