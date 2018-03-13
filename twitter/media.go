package twitter

import (
	"bytes"
	"net/http"

	"github.com/dghubble/sling"
)

// MediaParams paremeters for media request
type MediaParams struct {
	Image            []byte `url:"media,omitempty"`
	ImageBase64      string `url:"media_data,omitempty"`
	AdditionalOwners string `url:"additional_owners,omitempty"`
}

// Media return from twitter
type Media struct {
	ID       int64     `json:"media_id"`
	IDString string    `json:"media_id_string"`
	Size     int32     `json:"size"`
	LifeTime int32     `json:"expires_after_secs"`
	Image    MediaInfo `json:"image"`
}

func newMediaService(sling *sling.Sling) *MediaService {
	return &MediaService{
		sling: sling.Base("https://upload.twitter.com/1.1/").Path("media/"),
	}
}

// MediaInfo infos about the media returned by the twitter API
type MediaInfo struct {
	ImageType string `json:"image_type"`
	Width     int16  `json:"w"`
	Height    int16  `json:"h"`
}

// MediaService service stance
type MediaService struct {
	sling *sling.Sling
}

// UploadImage POST https://upload.twitter.com/1.1/media/upload.json, only suportting base64
func (m MediaService) UploadImage(params *MediaParams) (*Media, *http.Response, error) {
	media := new(Media)
	apiError := new(APIError)
	requestParams := &MediaParams{AdditionalOwners: params.AdditionalOwners}
	var resp *http.Response
	var err error
	if &params.Image != nil && len(params.Image) > 0 {
		resp, err = m.sling.New().Post("upload.json").Set("media_type", "image/png").Set(
			"total_bytes", string(len(params.Image))).Body(
			bytes.NewReader(params.Image)).BodyForm(requestParams).Receive(media, apiError)
	} else {
		requestParams.ImageBase64 = params.ImageBase64
		resp, err = m.sling.New().Post("upload.json").BodyForm(requestParams).Receive(media, apiError)
	}
	return media, resp, relevantError(err, *apiError)
}
