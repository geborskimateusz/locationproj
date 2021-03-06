package main

var usersNearby []User
var matches map[string]*Matches

// USERS SCHEMA

type User struct {
	Nickname    string      `json:"nickname"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func InitializeUsers() {
	firstUser := User{
		Nickname: "superuser123",
		Coordinates: Coordinates{
			X: 50.868529712179935,
			Y: 19.109939369334157,
		},
	}

	secondUser := User{
		Nickname: "anonuser",
		Coordinates: Coordinates{
			X: 50.866151341448436,
			Y: 19.11058980503063,
		},
	}

	thirdUser := User{
		Nickname: "thirdUSer",
		Coordinates: Coordinates{
			X: 50.856700985843815,
			Y: 19.108540597332587,
		},
	}
	fourthUser := User{
		Nickname: "JohnyFourthUserDoe",
		Coordinates: Coordinates{
			X: 50.8515960533141255,
			Y: 19.109704386867137,
		},
	}

	usersNearby = []User{}
	usersNearby = append(usersNearby, firstUser)
	usersNearby = append(usersNearby, secondUser)
	usersNearby = append(usersNearby, thirdUser)
	usersNearby = append(usersNearby, fourthUser)
}

func GetUsersNearby() []User {
	return usersNearby
}

// MATCHES SCHEMA

type PendingInvitations struct {
	Sent     []string `json:"sent"`
	Received []string `json:"received"`
}
type Matches struct {
	Pending PendingInvitations `json:"pending"`
	Matches []string           `json:"matches"`
}

func InitializeMatches() {

	firstMatch := Matches{
		Pending: PendingInvitations{
			Sent:     []string{},
			Received: []string{"fourthMatch"},
		},
		Matches: []string{"anonuser"},
	}

	secondMatch := Matches{
		Pending: PendingInvitations{
			Sent:     []string{},
			Received: []string{},
		},
		Matches: []string{"anonuser"},
	}

	thirdMatch := Matches{
		Pending: PendingInvitations{
			Sent:     []string{},
			Received: []string{"fourthMatch"},
		},
		Matches: []string{"superuser123", "JohnyFourthUserDoe"},
	}

	fourthMatch := Matches{
		Pending: PendingInvitations{
			Sent:     []string{"anonuser"},
			Received: []string{},
		},
		Matches: []string{},
	}

	matches = make(map[string]*Matches)
	matches["superuser123"] = &firstMatch
	matches["JohnyFourthUserDoe"] = &secondMatch
	matches["anonuser"] = &thirdMatch
	matches["fourthMatch"] = &fourthMatch
}

func GetMatches(username string) Matches {
	return *matches[username]
}

func UpdateMatches(username string, newMatches Matches) Matches {
	matches[username] = &newMatches
	return *matches[username]
}

func MatchRequest(username, invitedUser string) Matches {
	sent := &matches[username].Pending.Sent
	*sent = append(*sent, invitedUser)
	received := &matches[invitedUser].Pending.Received
	*received = append(*received, username)
	return *matches[username]
}

func CheckMatch(currentUser, requestedUser string) bool {
	currentUserMatches := &matches[currentUser].Matches
	requestedUserMatches := &matches[requestedUser].Matches

	requestedToCurrent := containsMatch(currentUserMatches, requestedUser)
	currentToRequested := containsMatch(requestedUserMatches, currentUser)

	if requestedToCurrent && currentToRequested {
		return true
	}
	return false
}

func containsMatch(matches *[]string, nickname string) bool {
	return contains(len(*matches), func(i int) bool {
		return (*matches)[i] == nickname
	})
}

func contains(n int, match func(i int) bool) bool {
	for i := 0; i < n; i++ {
		if match(i) {
			return true
		}
	}
	return false
}
