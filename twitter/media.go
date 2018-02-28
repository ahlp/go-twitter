package twitter

import (
	"io"
	"net/http"

	"github.com/dghubble/sling"
)

// MediaParams
type MediaParams struct {
	ImageBase64      string `json:"media"`
	AdditionalOwners string `json:"additional_owners"`
}

// MediaTweet
type Media struct {
	ID int64 `json:"media_id"`
	//   ID string `json:"media_id_string"`
	Size     int32     `json:"size"`
	LifeTime int32     `json:"expires_after_secs"`
	Image    MediaInfo `json:"image"`
}

func newMediaService(sling *sling.Sling) *MediaService {
	return &MediaService{
		sling: sling.Base("https://upload.twitter.com/1.1/").Path("media/"),
	}
}

// infos about the media returned by the twitter API
type MediaInfo struct {
	ImageType string `json:"image_type"`
	Width     int16  `json:"w"`
	Height    int16  `json:"h"`
}

// MediaService
type MediaService struct {
	sling *sling.Sling
}

// Send image in body to twitter
// POST https://upload.twitter.com/1.1/media/upload.json
func (m MediaService) UploadImage(image io.Reader) (*Media, *http.Response, error) {
	media := new(Media)
	apiError := new(APIError)
	resp, err := m.sling.New().Post("upload.json").Body(image).Receive(media, apiError)
	return media, resp, relevantError(err, *apiError)
}

// Send image in base64 to twitter
// POST https://upload.twitter.com/1.1/media/upload.json
func (m MediaService) UploadImageBase64(params MediaParams) (*Media, *http.Response, error) {
	media := new(Media)
	apiError := new(APIError)
	resp, err := m.sling.New().Post("upload.json").BodyForm(params).Receive(media, apiError)
	return media, resp, relevantError(err, *apiError)
}
