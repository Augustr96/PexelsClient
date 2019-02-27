package GoPexels

// result of search photos
type SearchResult struct {
	Page         float64 `json:"page"`
	PerPage      float64 `json:"per_page"`
	TotalResults float64 `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

// trending photos
type CuratedResult struct {
	Page     float64 `json:"page"`
	PerPage  float64 `json:"per_page"`
	NextPage string  `json:"next_page"`
	Photos   []Photo `json:"photos"`
}

// photo detail
type Photo struct {
	Id              float64     `json:"id"`
	Width           float64     `json:"width"`
	Height          float64     `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	Src             PhotoSource `json:"src"`
}

// url of photo
type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

// result of search video
type VideoSearchResult struct {
	Page         float64 `json:"page"`
	PerPage      float64 `json:"per_page"`
	TotalResults float64 `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Videos       []Video `json:"videos"`
}

// video detail
type Video struct {
	Id            float64         `json:"id"`
	Width         float64         `json:"width"`
	Height        float64         `json:"height"`
	Url           string          `json:"url"`
	Image         string          `json:"image"`
	FullRes       interface{}     `json:"full_res"` // fixme use?
	Duration      float64         `json:"duration"`
	VideoFiles    []VideoFiles    `json:"video_files"`
	VideoPictures []VideoPictures `json:"video_pictures"`
}

// video source
type VideoFiles struct {
	Id       float64 `json:"id"`
	Quality  string  `json:"quality"`
	FileType string  `json:"file_type"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Link     string  `json:"link"`
}

// video Schematic diagram
type VideoPictures struct {
	Id      float64 `json:"id"`
	Picture string  `json:"picture"`
	Nr      float64 `json:"nr"`
}

// popular photos
type PopularVideos struct {
	Page         float64 `json:"page"`
	PerPage      float64 `json:"per_page"`
	TotalResults float64 `json:"total_results"`
	Url          string  `json:"url"`
	Videos       []Video `json:"videos"`
}
