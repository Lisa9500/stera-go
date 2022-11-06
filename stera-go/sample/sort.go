package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	population := map[string]int{
		"Australia":  24982688,
		"Qatar":      2781677,
		"Wales":      3139000,
		"Burundi":    11175378,
		"Guinea":     12414318,
		"Niger":      22442948,
		"Brazil":     209469333,
		"Malta":      484630,
		"Peru":       31989256,
		"Yemen":      28498687,
		"Ireland":    4867309,
		"Kenya":      51393010,
		"Montserrat": 5900,
		"Cuba":       11338138,
		"Nicaragua":  6465513,
		"Jordan":     9956011,
		"Gabon":      2119275,
	}

	p := make(PairList, len(population))

	i := 0
	for k, v := range population {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	//p is sorted

	for _, k := range p {
		fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}
}
