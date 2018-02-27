package twitter

import "github.com/dghubble/sling"

// MediaParams
type MediaParams struct {
	Media            string `json:"media"`
	MediaData        []byte
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

func newMediaService(sling *sling.Sling) *StatusService {
	return &StatusService{
		sling: sling.Path("media/"),
	}
}

type MediaInfo struct {
	ImageType string `json:"image_type"`
	Width     int16  `json:"w"`
	Height    int16  `json:"h"`
}

// MediaService
type MediaService struct {
	sling *sling.Sling
}

// POST https://upload.twitter.com/1.1/media/upload.json
func (m MediaService) uploadImage(params MediaParams) {

	media := new(Media)
	apiError := new(APIError)
	path := "upload.json"
	resp, err := m.sling.New().Post(path).BodyForm(params).Receive(media, apiError)
	// return tweet, resp, relevantError(err, *apiError)

	// if raw data
	// if media.MediaData {

	//  return
	// }
	// else base64

}
