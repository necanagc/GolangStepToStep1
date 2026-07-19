package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// func ToLowerSlice(t []string) []string {
// 	for i, r := range t {
// 		t[i] = strings.ToLower(r)
// 	}
// 	return t
// }

// func findMaxWord(t map[string]int) (string, int) {
// 	max := -1
// 	var maxWord string

// 	for key, _ := range t {
// 		if t[key] > max {
// 			max = t[key]
// 			maxWord = key
// 		}
// 	}

// 	return maxWord, max
// }

// func getTopWords(wordMap map[string]int, n int) []string {
// 	var res []string
// 	var maxWord string
// 	for i := 0; i < n; i++ {
// 		maxWord, _ = findMaxWord(wordMap)
// 		res = append(res, maxWord)
// 		delete(wordMap, maxWord)
// 	}

// 	return res
// }

// func AnalyzeText(text string) {

// 	t := strings.FieldsFunc(text, func(r rune) bool {
// 		return unicode.IsPunct(r) || unicode.IsSpace(r)
// 	})
// 	t = ToLowerSlice(t)

// 	textMap := make(map[string]int)

// 	for _, key := range t {
// 		if _, ok := textMap[key]; !ok {
// 			textMap[key] = 1
// 		} else {
// 			textMap[key]++
// 		}
// 	}

// 	countUniq := len(textMap)

// 	maxWord, maxWordCount := findMaxWord(textMap)

// 	textMapCopy := make(map[string]int, len(textMap))
// 	for k, v := range textMap {
// 		textMapCopy[k] = v
// 	}

// 	top5res := getTopWords(textMapCopy, 5)

// 	fmt.Printf("Количество слов: %v\n", len(t))
// 	fmt.Printf("Количество уникальных слов: %v\n", countUniq)
// 	fmt.Printf("Самое часто встречающееся слово: \"%v\" (встречается %d раз)\n", maxWord, maxWordCount)

// 	fmt.Println("Топ-5 самых часто встречающихся слов:")
// 	for _, g := range top5res {
// 		fmt.Printf("\"%v\": %d раз\n", g, textMap[g])
// 	}

// }
