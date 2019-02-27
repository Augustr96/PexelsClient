package GoPexels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

type Client struct {
	Token string
	hc    http.Client
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

// search photos
// query: Get photos related to this query. (required)
// per_page: Defines the number of results per page. (optional, default: 15, max: 80)
// page: Defines the number of the page. (optional, default: 1)
func (c *Client) SearchPhotos(query string, perPage, page int) ( *SearchResult, error){
	url := fmt.Sprintf(PhotoApi + "/search?query=%s&per_page=%d&page=%d", query, perPage, page)
	resp, err := c.requestDoWithAuth("GET", url)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil ,err
	}
	var result SearchResult
	err = json.Unmarshal(data, &result)
	return &result, err
}

// get trending photos
// per_page: Defines the number of results per page. (optional, default: 15, max: 80)
// page: Defines the number of the page. (optional, default: 1)
func (c *Client) CuratedPhotos(perPage, page int) (*CuratedResult, error) {
	url := fmt.Sprintf(PhotoApi + "/curated?per_page=%d&page=%d", perPage, page)
	resp , err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err  := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var resutl CuratedResult
	err = json.Unmarshal(data, &resutl)
	return &resutl, err
}

// get a photo
// id: Defines the photo id
func (c *Client) GetPhoto(id int32) (*Photo, error) {
	url := fmt.Sprintf(PhotoApi + "/photos/%d", id)
	resp, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	var result Photo
	err = json.Unmarshal(data, &result)
	return &result, err
}

// get a random photo
func (c *Client) GetRandomPhoto() (*Photo, error) {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1001)
	result, err := c.CuratedPhotos(1, randNum)
	if err == nil && len(result.Photos) == 1 {
		return &result.Photos[0], nil
	}
	return nil, err
}

// search videos
// query: Get videos related to this query. (required)
// per_page: Defines the number of results per page. (optional, default: 15, max: 80)
// page: Defines the number of the page. (optional, default: 1)
func (c *Client) SearchVideo(query string, perPage, page int) (*VideoSearchResult, error) {
	url := fmt.Sprintf(VideoApi + "/search?query=%s&per_page=%d&page=%d", query, perPage, page)
	resp, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result VideoSearchResult
	err = json.Unmarshal(data, &result)
	return &result, err
}

// get popular videos
// per_page: Defines the number of results per page. (optional, default: 15, max: 80)
// page: Defines the number of the page. (optional, default: 1)
func (c *Client) PopularVideo(perPage, page int) (*PopularVideos, error) {
	url := fmt.Sprintf(VideoApi + "/popular?per_page=%d&page=%d", perPage, page)
	resp, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data ,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result PopularVideos
	err = json.Unmarshal(data, &result)
	return &result, err
}

// get a random video
func (c *Client) GetRandomVideo() (*Video, error) {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1001)
	result, err := c.PopularVideo(1, randNum)
	if err == nil && len(result.Videos) == 1 {
		return &result.Videos[0], nil
	}
	return nil, err
}

func (c *Client) requestDoWithAuth(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	return c.hc.Do(req)
}
