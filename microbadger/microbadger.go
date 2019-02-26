package microbadger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type MicroBadgerResponse struct {
	ImageName  string `json:"ImageName"`
	WebhookURL string `json:"WebhookURL"`
	Labels     struct {
		OrgLabelSchemaVersion       string    `json:"org.label-schema.version"`
		OrgLabelSchemaDescription   string    `json:"org.label-schema.description"`
		OrgLabelSchemaVendor        string    `json:"org.label-schema.vendor"`
		OrgLabelSchemaVcsURL        string    `json:"org.label-schema.vcs-url"`
		OrgLabelSchemaURL           string    `json:"org.label-schema.url"`
		OrgLabelSchemaVcsRef        string    `json:"org.label-schema.vcs-ref"`
		OrgLabelSchemaName          string    `json:"org.label-schema.name"`
		OrgLabelSchemaSchemaVersion string    `json:"org.label-schema.schema-version"`
		OrgLabelSchemaBuildDate     time.Time `json:"org.label-schema.build-date"`
	} `json:"Labels"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	ImageURL    string    `json:"ImageURL"`
	ID          string    `json:"Id"`
	LastUpdated time.Time `json:"LastUpdated"`
	Description string    `json:"Description"`
	LatestSHA   string    `json:"LatestSHA"`
	Versions    []struct {
		VersionControl struct {
			Commit string `json:"Commit"`
			URL    string `json:"URL"`
			Type   string `json:"Type"`
		} `json:"VersionControl"`
		ImageName string    `json:"ImageName"`
		Created   time.Time `json:"Created"`
		License   struct {
			Code string `json:"Code"`
			URL  string `json:"URL"`
		} `json:"License"`
		SHA    string `json:"SHA"`
		Labels struct {
			ComMicroscalingLicense          string    `json:"com.microscaling.license"`
			OrgLabelSchemaVersion           string    `json:"org.label-schema.version"`
			OrgLabelSchemaDescription       string    `json:"org.label-schema.description"`
			OrgLabelSchemaURL               string    `json:"org.label-schema.url"`
			OrgLabelSchemaVcsURL            string    `json:"org.label-schema.vcs-url"`
			OrgLabelSchemaVendor            string    `json:"org.label-schema.vendor"`
			OrgLabelSchemaName              string    `json:"org.label-schema.name"`
			OrgLabelSchemaVcsRef            string    `json:"org.label-schema.vcs-ref"`
			OrgLabelSchemaSchemaVersion     string    `json:"org.label-schema.schema-version"`
			ComMicroscalingDockerDockerfile string    `json:"com.microscaling.docker.dockerfile"`
			OrgLabelSchemaBuildDate         time.Time `json:"org.label-schema.build-date"`
		} `json:"Labels"`
		DownloadSize int `json:"DownloadSize"`
		Tags         []struct {
			Tag string `json:"tag"`
		} `json:"Tags"`
		LayerCount int    `json:"LayerCount"`
		Author     string `json:"Author"`
	} `json:"Versions"`
	PullCount     int    `json:"PullCount"`
	DownloadSize  int    `json:"DownloadSize"`
	LatestVersion string `json:"LatestVersion"`
	LayerCount    int    `json:"LayerCount"`
	Author        string `json:"Author"`
}

func GetImage(path string) (MicroBadgerResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.microbadger.com/%s", path))
	if err != nil {
		return MicroBadgerResponse{}, fmt.Errorf("error getting image details from MicroBadger: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MicroBadgerResponse{}, fmt.Errorf("error reading Body of MicroBadger response: %v", err)
	}

	var imageResponse MicroBadgerResponse

	err = json.Unmarshal(body, &imageResponse)
	if err != nil {
		return MicroBadgerResponse{}, fmt.Errorf("error unmarshalling MicroBadger response to struct: %v", err)
	}

	return imageResponse, nil
}
