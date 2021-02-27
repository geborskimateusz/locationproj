package main

import (
	"testing"
)

func TestDB(t *testing.T) {
	InitializeMatches()
	InitializeUsers()

	t.Run("Match request", func(t *testing.T) {

		firstUser := "superuser123"
		secUser := "JohnyFourthUserDoe"

		MatchRequest(firstUser, secUser)

		firstUserMatches := GetMatches(firstUser)
		secondUserMatches := GetMatches(secUser)

		assertContains(t, firstUserMatches.Pending.Sent, secUser)
		assertContains(t, secondUserMatches.Pending.Received, firstUser)

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
