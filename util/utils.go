package util

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

const projectDirName = "genvaeg"

var once sync.Once

func InitEnv() {
	once.Do(InitEnvExec)
}

func InitEnvExec() {
	envFile := filepath.Join(string(GetRootDir()), ".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("error loading env file:", envFile)
	}
}

// We have to dynamically find the project root directory, because
// it's different for tests and the main server.
func GetRootDir() string {
	projectRootDirRe := regexp.MustCompile(`^.*` + projectDirName + ``)
	cwd, _ := os.Getwd()
	projectRootDir := projectRootDirRe.Find([]byte(cwd))
	return string(projectRootDir)
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
