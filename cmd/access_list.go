package main

import (
	"fmt"

	"cloud.google.com/go/bigquery"
)

// AccessList is a slice of bigquery.AccessEntry to facilitate operation on
type AccessList []*bigquery.AccessEntry

func (a AccessList) Len() int {
	return len(a)
}

func (a AccessList) Less(i, j int) bool {
	return a[i].Entity < a[j].Entity
}

func (a AccessList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Entries return a slice of string formatted Entity and Role
func (a AccessList) Entries() []string {
	var result []string
	for _, entry := range a {
		result = append(result, fmt.Sprintf("%s:%s\n", entry.Entity, entry.Role))
	}
	return result
}
