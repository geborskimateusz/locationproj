package main

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	InitializeMatches()
	InitializeUsers()

	t.Run("Match request", func(t *testing.T) {

		firstUser := "superuser123"
		secUser := "JohnyFourthUserDoe"

		a := GetMatches(firstUser)
		fmt.Println(a)
		MatchRequest(firstUser, secUser)

		firstUserMatches := GetMatches(firstUser)
		secondUserMatches := GetMatches(secUser)

		fmt.Println(firstUserMatches)
		// fmt.Println(secondUserMatches)

		assertContains(t, firstUserMatches.Pending.Sent, secUser)
		assertContains(t, secondUserMatches.Pending.Received, firstUser)

	})

	t.Run("CheckMatches", func(t *testing.T) {

		firstUser := "anonuser"
		secUser := "JohnyFourthUserDoe"

		containsFirstUser := CheckMatch(secUser, firstUser)
		secUserMatches := GetMatches(secUser)

		if !containsFirstUser {
			t.Errorf("Element %v does not exist in %v", firstUser, secUserMatches)
		}

	})
}

func TestGenerateChannel(t *testing.T) {

	t.Run("Match request", func(t *testing.T) {

		firstUser := "superuser123"
		secUser := "JohnyFourthUserDoe"

		firstHash := GenerateUniqueChannel(firstUser, secUser)
		secHash := GenerateUniqueChannel(firstUser, secUser)
		if firstHash != secHash {
			t.Errorf("Different channels %v and %v", firstHash, secHash)
		}

	})

}

func assertContains(t *testing.T, s []string, e string) {
	t.Helper()

	contains := false
	for _, a := range s {
		if a == e {
			contains = true
		}
	}
	if !contains {
		t.Errorf("Element %v does not exist in %v", e, s)
	}
}
