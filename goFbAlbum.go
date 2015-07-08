package goFbAlbum

import (
	"encoding/json"
	"errors"
	"log"

	fb "github.com/huandu/facebook"
)

func init() {

}

type FbAlbum struct {
	Token string
}

// Constructor
func NewFbAlbum(token string) *FbAlbum {
	if token == "" {
		return nil
	}
	f := new(FbAlbum)
	f.Token = token
	return f
}

// Get my all albums
func (self *FbAlbum) GetMyAlbums() (*FBAlbums, error) {
	return self.GetAlbumsByUserId("me")
}

// Get all album using user id.
// note: this function only work if you provide a page id or page name. such as scottiepippen or 112743018776863.
func (self *FbAlbum) GetAlbumsByUserId(uid string) (*FBAlbums, error) {
	if uid == "" {
		return nil, errors.New("uid is empty")
	}
	resAlbum := self.RunFBGraphAPI("/" + uid + "/albums")
	retAlbum := FBAlbums{}
	ParseMapToStruct(resAlbum, &retAlbum)
	return &retAlbum, nil
}

// Get all photo from a album id, you can get album id from FBAlbums{} struct.
func (self *FbAlbum) GetPhotoByAlbum(albumId string) (*FBPhotos, error) {
	if albumId == "" {
		return nil, errors.New("albumId is empty")
	}
	photoRet := FBPhotos{}
	resPhoto := self.RunFBGraphAPI("/" + albumId + "/photos")
	ParseMapToStruct(resPhoto, &photoRet)
	return &photoRet, nil
}

// FaceBook Graph Query API.
func (self *FbAlbum) RunFBGraphAPI(query string) (queryResult interface{}) {
	res, err := fb.Get(query, fb.Params{
		"access_token": self.Token,
	})

	if err != nil {
		log.Fatalln("FB connect error, err=", err.Error())
	}
	return res
}

// Parse Graph API result and convert to specific interface.
func ParseMapToStruct(inData interface{}, decodeStruct interface{}) {
	jret, _ := json.Marshal(inData)
	err := json.Unmarshal(jret, &decodeStruct)
	if err != nil {
		log.Fatalln(err)
	}
}
