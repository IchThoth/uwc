package uwc

type PhotoUrls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}
type PhotoLinks struct {
	Self      string `json:"self"`
	Html      string `json:"html"`
	Photos    string `json:"photos"`
	Likes     string `json:"photos"`
	Portfolio string `json:"portfolio"`
}

//Photos have the following link relations:

//rel	Description
//self	API location of this photo.
//html	HTML location of this photo.
//download	Download location of this photo.
type PhotoLinksR struct {
	Self             string `json:"self"`
	Html             string `json:"html"`
	Download         string `json:"photos"`
	DownloadLocation string `json:"download_location"`
}

type Photo struct {
	Id                 string           `json:"id"`
	CreatedAt          string           `json:"created_at"`
	UpdatedAt          string           `json:"updated_at"`
	Width              int32            `json:"width"`
	Height             int32            `json:"height"`
	Color              string           `json:"color"`
	BlurHash           string           `json:"blur_hash"`
	Likes              int32            `json:"likes"`
	LikedByUser        bool             `json:"liked_by_user"`
	Description        string           `json:"description"`
	User               User             `json:"user"`
	Links              PhotoLinks       `json:"links"`
	CurrentCollections []UserCollection `json:"current_user_collections"`
	Urls               PhotoUrls        `json:"urls"`
	LinksR             PhotoLinksR      `json:"links"`
}
