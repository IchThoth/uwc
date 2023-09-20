package uwc

import (
	"context"
	"fmt"
	"time"
)

type ProfileImg struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type Photo struct {
	Id           string     `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Color        string     `json:"color"`
	BlurHash     string     `json:"blur_hash"`
	Likes        int        `json:"likes"`
	LikedByUser  bool       `json:"liked_by_user"`
	Description  string     `json:"description"`
	User         []User     `json:"user"`
	ProfileImage ProfileImg `json:"profile_image"`
	Links        Links      `json:"links"`
}

func (c *Client) Photos(ctx context.Context, page int, perPage int, orderBy string) (*Photo, error) {
	unsplashUrl := fmt.Sprintf(c.BaseUrl+"/photos?page=%d&per_page=%d&order_by=%s", page, perPage, orderBy)

	var result Photo

	err := c.Get(ctx, unsplashUrl, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (c *Client) PhotoByID(ctx context.Context, id string) (*Photo, error) {
	unsplashUrl := fmt.Sprintf(c.BaseUrl+"/photos?id=%s", id)
	var result Photo
	err := c.Get(ctx, unsplashUrl, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (c *Client) LikedPhotos(ctx context.Context) bool {
	//do something
	unsplashUrl := fmt.Sprintf(c.BaseUrl)
}

func (c *Client) RandomPhoto(ctx context.Context, collection string,
	topics string, useraname string, query string,
	orientation string, contentFilter string, count string) ([]Photo, error) {
	//do someething :(
	unsplashUrl := fmt.Sprintf(c.BaseUrl + "/photos/random")
}

func PhotoStats() {

}

func TrackDownload() {

}

func UpdatePhoto() {

}

func LikePhoto() {

}

func UnlikePhoto() {

}
