package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

func GenerateUniqueChannel(currentUser, connectedTo string) string {
	arr := []string{currentUser, connectedTo}
	sort.Sort(sort.StringSlice(arr))
	result := strings.Join(arr, "")
	hash := sha256.Sum256([]byte(result))
	return fmt.Sprintf("%x", hash)
}
