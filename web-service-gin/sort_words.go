package main

import (
	"sort"
)

var ordenation WordsOrder = WordsOrder{asc: "asc", desc: "desc"}

func SortWords(value []string, order string) []string {
	if order == ordenation.desc {
		return sortWordsDesc(value)
	}
	return sortWordsAsc(value)

}

func sortWordsAsc(value []string) []string {
	sort.Strings(value)
	return value
}

func sortWordsDesc(value []string) []string {
	sort.Slice(value, func(i, j int) bool {
		// retorno dá função anonima
		return value[i] > value[j]
	})
	return value
}
