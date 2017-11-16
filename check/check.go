package main

import (
	"os"
	"encoding/json"
	"os/exec"
)

func main() {

	if len(os.Args) < 2 {
	}

	oreq := CheckRequest{}

	err := json.NewDecoder(os.Stdin).Decode(oreq)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("kubectl", "--certificate-authority", oreq.Source.CertAuthority, "--client-key",
		oreq.Source.ClientKey, "--client-certificate", oreq.Source.ClientCert, "--server", oreq.Source.ClusterURL, "get", "pods")

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

}

type CheckRequest struct {
	Source Source `json:"source"`
}

//Source ...
type Source struct {
	ClusterURL    string `json:"cluster_url"`
	CertAuthority string `json:"certificate_authority"`
	ClientKey     string `json:"client_key"`
	ClientCert    string `json:"client_certificate"`
	Namespace     string `json:"namespace"`
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
