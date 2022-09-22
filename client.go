package uwc

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	PhotoApi = "https://api.unsplash.com/photos"
)

type Client struct {
	Token  string
	apiUrl string
	hc     http.Client
}
type SearchResult struct {
	page int32 `json:"page"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token := os.Getenv("ACCESS_KEY")
	SecretToken := os.Getenv("SECRET_KEY")

}
