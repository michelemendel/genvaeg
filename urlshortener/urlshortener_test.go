package urlshortener

import (
	"fmt"
	"testing"

	"github.com/michelemendel/genvaeg/entity"
	"github.com/michelemendel/genvaeg/util"
)

const (
	CHAR_SET_SHORT = "ab" //We make this very short to quickly hit the limit of unique short URLs, to easier test that the algorithm works.
	BASE_URL       = "http://genvaeg.com"
)

func TestUrlShortener(t *testing.T) {
	repo := InitTest()

	// Create user
	hPw, _ := util.HashPassword("test")
	user := entity.NewUser("abe", hPw)
	err := repo.CreateUser(user)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	us := NewURLShortener(CHAR_SET_SHORT, repo)
	us.StartLength = 2

	var fullURL string

	// Create URLs
	for i := 0; i < 5; i++ {
		fullURL = makeFullURL(i)
		us.MakeShortURL(fullURL, user)
	}

	// Verify that the short URLs are unique
	shorts := []string{}
	urlPairs, _ := repo.GetAllURLPairs()
	for _, urlPair := range urlPairs {
		shorts = append(shorts, urlPair.ShortURL)
	}
	removeDuplicate(shorts)
	if len(shorts) != len(urlPairs) {
		t.Errorf("Short URLs are not unique")
	}
}

//--------------------------------------------------------------------------------
// Test helpers

func removeDuplicate(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func makeFullURL(nr int) string {
	return fmt.Sprintf("http://site_%d.com", nr)
}
