package star

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	fistStargazersURL = "https://api.github.com/repos/fanux/fist/stargazers?page=%d&per_page=300"
	token             = "Bearer 450c54a01f33800cbb2598661f3d45f027a10faa"
)

//User is
type User struct {
	Login string `json:"login,omitempty"`
}

func isIn(user string, users []User) bool {
	for _, u := range users {
		if u.Login == user {
			return true
		}
	}
	return false
}

//IsStared is
func IsStared(user string) bool {
	resp, err := http.Get(fistStargazersURL)
	if err != nil {
		fmt.Println("error", err)
		return false
	}

	defer resp.Body.Close()
	us := &[]User{}
	err = json.NewDecoder(resp.Body).Decode(us)
	if err != nil {
		fmt.Println("json decode error: ", err)
		return false
	}

	return isIn(user, *us)
}

func starPage(user string, page int) bool {
	client := &http.Client{}

	url := fmt.Sprintf(fistStargazersURL, page)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error", err)
		fmt.Println("url", url)
		return false
	}

	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	us := &[]User{}
	err = json.NewDecoder(resp.Body).Decode(us)
	if err != nil {
		fmt.Println("json decode error: ", err)
		return false
	}
	return isIn(user, *us)
}

//IsStaredUnlimit is
func IsStaredUnlimit(user string) bool {
	for i := 1; i < 4; i++ {
		if starPage(user, i) {
			return true
		}
	}
	return false
}

func testStar() {
	/*
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: "7ff0cb681d70a19445f129197b088f88a130c07e"},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		// list all repositories for the authenticated user
		repos, _, err := client.Repositories.Get(ctx, "fanux", "fist")
		if err != nil {
			fmt.Println(err)
		}
	*/
	fmt.Println(IsStared("fanux"))
	fmt.Println(IsStared("aaa"))
}
