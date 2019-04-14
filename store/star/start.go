package star

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	fistStargazersURL = "https://api.github.com/repos/fanux/sealos/stargazers?page=%d&per_page=300"
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

//CheckFree is
func CheckFree(user string, pro string) bool {
	return IsStaredUnlimit(user) && (pro == "kubernetes1.14.1-HA")
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
