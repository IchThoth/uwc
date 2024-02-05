package unsplashauth

import "golang.org/x/oauth2"

//This is the URL that links to the OAuth2 endpoint for Unsplash
const (

	//AuthenticationURL is used by Unsplash accounts for Oauth
	AuthenticationURL = "https://unsplash.com/oauth/authorize"

	//TokenUrl is the endpoint to send bearer tokens
	TokenUrl = "https://unsplash.com/oauth/token"
)

// Scopes let you specify exactly which types of data your application wants to access.
// The set of scopes you pass in your authentication request determines what access the
// permissions the user is asked to grant.
const (
	// ScopePublic is Default.It Read public data.
	ScopePublic = "public"
	//ScopeReadUser is used to access user’s private data.
	ScopeReadUser = "read_user"
	// ScopeWriteUser is used to update the user’s profile.
	ScopeWriteUser = "write_user"
	// ScopeReadPhotos is used to read private data from the user’s photos.
	ScopeReadPhotos = "read_photos"
	//ScopeWritePhotos is used to update photos on the user’s behalf.
	ScopeWritePhotos = "write_photos"
	//ScopeWriteLikes is used to like or unlike a photo on the user’s behalf.
	ScopeWriteLikes = "write_likes"
	//ScopeWriteFollowers is used to follow or unfollow a user on the user’s behalf.
	ScopeWriteFollowers = "write_followers"
	//ScopeReadCollections is used to view a user’s private collections.
	ScopeReadCollections = "read_collections"
	// ScopeWriteCollections is used to create and update a user’s collections.
	ScopeWriteCollections = "write_collections"
)

type Authenticator struct {
	config *oauth2.Config
}

type Authenticate func(a *Authenticate)

func UseClientID()
