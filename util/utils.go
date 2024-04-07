package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var once sync.Once

func InitEnv() {
	once.Do(InitEnvExec)
}

func InitEnvExec() {
	envFile := ".env"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("[utils]:error loading env file:", envFile)
	}
}

func GenerateUUID() string {
	return uuid.NewString()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		slog.Error("Error hashing password", "error", err)
		// Just create some random data, so the password is not empty and can be compromised.
		return string(GenerateUUID()), err
	}
	return string(bytes), nil
}

func ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Generate a random password with n charactes.
func GeneratePassword() string {
	aLongString := base64.StdEncoding.EncodeToString([]byte(GenerateUUID()))
	return strings.ToLower(aLongString[0:8])
}

// Generate a random string with n charactes.
// Used for the URL shortening.
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func String2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Int2String(i int) string {
	return strconv.Itoa(i)
}

func String2Bool(s string) bool {
	return s != ""
}

func Bool2String(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func PP(s any) {
	res, err := PrettyStruct(s)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(res)
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
