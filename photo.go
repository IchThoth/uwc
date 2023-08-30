package uwc

import "time"

type ProfileImg struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type Photo struct {
	Id           string       `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Width        int          `json:"width"`
	Height       int          `json:"height"`
	Color        string       `json:"color"`
	BlurHash     string       `json:"blur_hash"`
	Likes        int          `json:"likes"`
	LikedByUser  bool         `json:"liked_by_user"`
	Description  string       `json:"description"`
	User         []User       `json:"user"`
	ProfileImage []ProfileImg `json:"profile_image"`
	Links        []Links      `json:"links"`
}
