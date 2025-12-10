package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func getStats(text string) (map[string]int, []string, []string, error) {
	if text == "" {
		return nil, nil, nil, fmt.Errorf("текст пустой")
	}

	text = strings.ToLower(text)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(text, f)

	uniqueWords := make(map[string]int)

	for _, word := range words {
		uniqueWords[word]++
	}

	keys := make([]string, 0, len(uniqueWords))
	for key := range uniqueWords {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return uniqueWords[keys[i]] > uniqueWords[keys[j]]
	})
	return uniqueWords, keys[:min(5, len(keys))], words, nil
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите многострочный текст (конец - пустая строка):")

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	text := strings.Join(lines, " ")
	uniqueWords, top5, words, err := getStats(text)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("\nОбщее число слов:", len(words))
	fmt.Println("\nЧисло уникальных слов:", len(uniqueWords))
	fmt.Println("\nТоп-5 самых частых слов:")
	for _, word := range top5 {
		fmt.Println("  ", word, " - ", uniqueWords[word])
	}
}
