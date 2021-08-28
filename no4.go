package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	listString := strings.Split(s, "")
	sort.Strings(listString)

	return strings.Join(listString, "")

}

func main() {
	var arrWord = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	var listAnagram = make(map[string][]string)

	for _, sentence := range arrWord {
		key := sortString(sentence)
		listAnagram[key] = append(listAnagram[key], sentence)
	}

	for _, sentences := range listAnagram {
		for _, w := range sentences {
			fmt.Println("[", w, "]")
		}
	}

	// fmt.Println(arrWord)
}
