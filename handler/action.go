package handler

import (
	"math"
	"math/rand"

	"github.com/mccune1224/betrayal-widget/data"
)

// Key value of queue order of action types
var (
	ActionPriority = map[string]int{
		"Alteration":                           1,
		"Reactive":                             2,
		"Redirection":                          3,
		"Visit Redirection":                    3,
		"Investigation":                        4,
		"Protection":                           5,
		"Visit Blocking":                       6,
		"Vote Immunity":                        7,
		"Vote Manipulation":                    8,
		"Vote Blocking":                        8,
		"Support":                              9,
		"Debuff":                               10,
		"Theft":                                11,
		"Healing":                              12,
		"Destruction":                          13,
		"Killing":                              14,
		"Last: (exception): Detective's Solve": 15,
	}
)

func insert[T any](items []T, item T, i int) {
	items = append(items[:i+1], items[i:]...)
	items[i] = item
}

func scramble[T any](items []T) []T {
	perms := rand.Perm(len(items))
	newItems := make([]T, len(items))
	for i := 0; i < len(items); i++ {
		p := perms[i]
		newItems[i] = items[p]
	}

	return newItems
}

func highestPriority(action data.Action) int {
	lowest := math.MaxInt
	for _, cat := range action.Categories {
		rank := ActionPriority[cat]
		if lowest > rank && rank != 0 {
			// have to check for 0 because of default value in maps
			lowest = rank
		}
	}
	return lowest
}

func flatten[T any](lists [][]T) []T {
	var res []T
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func shuffle[T any](items []T) []T {
	perms := rand.Perm(len(items))
	newItems := make([]T, len(items))
	for i := 0; i < len(items); i++ {
		p := perms[i]
		newItems[i] = items[p]
	}
	return newItems
}

func sortComplexPlayerRequest(actions []*data.ComplexPlayerRequest) []data.ComplexPlayerRequest {
	buckets := make([][]data.ComplexPlayerRequest, len(ActionPriority))

	for _, v := range actions {
		prio := highestPriority(v.A)
		buckets[prio] = append(buckets[prio], *v)
	}

	for i := range buckets {
		buckets[i] = shuffle(buckets[i])
	}

	return flatten(buckets)
}

// func sortactionbypriority(actions []data.action) []data.action {
// 	buckets := make([][]data.action, len(actionpriority))
//
// 	for _, v := range actions {
// 		prio := highestpriority(v)
// 		buckets[prio] = append(buckets[prio], v)
// 	}
//
// 	for i := range buckets {
// 		buckets[i] = shuffle(buckets[i])
// 	}
//
// 	return flatten(buckets)
// }
