package main

import (
	"os"
	"encoding/json"
)

func main() {

	if len(os.Args) < 2 {
	}

	oreq := OutRequest{}

	err := json.NewDecoder(os.Stdin).Decode(oreq)
	if err != nil {

	}

}

type OutRequest struct {
	Source Source `json:"source"`
}

//Source ...
type Source struct {
	ClusterURL string `json:"cluster_url"`
	ClusterCA  string `json:"cluster_ca"`
	AdminKey   string `json:"admin_key"`
	AdminCert  string `json:"admin_cert"`
}

type OutResponse struct {
	Version Version `json:""`
}

//Version ...
type Version struct {
	Version string `json:"version"`
}

//Metadata ...
type Metadata struct {
	Pods []string `json:"metadata"`
}

//MetadataField ...
type MetadataField struct {
	Key   string
	Value string
}
