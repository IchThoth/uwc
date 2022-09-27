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
	Token          string
	Secret         string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string, secret string) *Client {
	c := http.Client{}
	return &Client{
		Token:  token,
		Secret: secret,
		hc:     c,
	}
}

//Requests that return multiple items (a list of photos, for example) will
//be paginated into pages of 10 items by default, up to a maximum of 30.
//The optional page and per_page query parameters can be supplied to define
//which page and the number of items per page to be returned, respectively.
//If page is not supplied, the first page will be returned.
//URL’s for the first, last, next, and previous pages are supplied,
//if applicable. They are comma-separated and differentiated with a rel attribute.

func main() {
	//load environment variables
	err := godotenv.Load("uwc.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AccessToken := os.Getenv("ACCESS_KEY")
	SecretKey := os.Getenv("SECRET_KEY ")

	var client = NewClient(AccessToken, SecretKey)

}
