package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mitchellh/mapstructure"
)

type Secure struct {
	Status string `json:"status"`
	Data   struct {
		Layer struct {
			Features []struct {
				Name            string `json:"Name"`
				VersionFormat   string `json:"VersionFormat"`
				NamespaceName   string `json:"NamespaceName"`
				AddedBy         string `json:"AddedBy"`
				Version         string `json:"Version"`
				Vulnerabilities []struct {
					Name          string      `json:"Name"`
					NamespaceName string      `json:"NamespaceName"`
					Link          string      `json:"Link"`
					FixedBy       string      `json:"FixedBy"`
					Description   string      `json:"Description"`
					Metadata      interface{} `json:"Metadata"`
					Severity      string      `json:"Severity"`
				} `json:"Vulnerabilities"`
			} `json:"Features"`
		} `json:"Layer"`
	} `json:"data"`
}

type Image struct {
	Organisation string `json:"organisation"`
	Repository   string `json:"repository"`
	Tag          string `json:"tag"`
}

func main() {
	var images []Image

	jsonFile, err := os.Open("input.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &images)

	// images getVulnerabilities
	var res []Secure
	for _, i := range images {
		res = append(res, getVulnerabilities(i))
	}

	// Printout result
	for _, i := range res {
		for _, sec := range i.Data.Layer.Features {
			fmt.Println(sec.Vulnerabilities)
		}
	}
}

func getVulnerabilities(image Image) Secure {

	manifest, err := getManifestDigest(image)
	if err != nil {
		panic("Error when get image manifest from https://quay.io. ")
	}

	vulnerabilities := "https://quay.io/api/v1/repository/" + image.Organisation + "/" + image.Repository + "/manifest/" + manifest + "/security" + "?vulnerabilities=true"

	resp, err := http.Get(vulnerabilities)
	if err != nil {
		panic(err)
	}

	// parse respon to map & unmarshal
	var data map[string]interface{}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		panic(err)
	}

	// parse to Secure structe
	var s Secure

	err = mapstructure.Decode(data, &s)
	if err != nil {
		panic(err)
	}

	return s
}

func getManifestDigest(image Image) (manifest string, err error) {
	// Get all tags
	url := "https://quay.io/api/v1/repository/" + image.Organisation + "/" + image.Repository + "/tag/"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// a map to save json
	var data map[string]interface{}

	// unmarshal
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		panic(err)
	}

	// find manifest digest of input image
	for _, item := range data["tags"].([]interface{}) {
		if item.(map[string]interface{})["name"] == image.Tag {
			return item.(map[string]interface{})["manifest_digest"].(string), nil
		}
	}
	return "", err
}
