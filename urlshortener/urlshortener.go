package urlshortener

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/michelemendel/genvaeg/entity"
	"github.com/michelemendel/genvaeg/repository"
)

type URLShortener struct {
	Charset     string
	Repo        *repository.Repo
	StartLength int
}

func NewURLShortener(charset string, repo *repository.Repo) *URLShortener {
	return &URLShortener{
		Charset:     charset,
		Repo:        repo,
		StartLength: 5,
	}
}

// Repeat to generates a short URL until we get a unique one
func (us *URLShortener) MakeShortURL(fullURL string, user entity.User) error {
	length := us.StartLength
	for {
		err := us.TryShortURL(length, fullURL, user)
		if err == nil {
			break
		}
		length++
	}
	return nil
}

func (us *URLShortener) TryShortURL(length int, fullURL string, user entity.User) error {
	shortURLPath := us.GenerateRandomPath(length)
	urlPair := entity.NewURLPair(fullURL, shortURLPath)
	fmt.Println(fullURL, shortURLPath)
	return us.Repo.CreateURLPair(*urlPair, user)
}

func (us *URLShortener) GenerateRandomPath(urlLength int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, urlLength)
	for i := range result {
		result[i] = us.Charset[random.Intn(len(us.Charset))]
	}
	return string(result)
}

func (us URLShortener) IsURLValid(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}
