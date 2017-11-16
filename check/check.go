package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func main() {

	if len(os.Args) < 2 {
	}

	CheckRequest := CheckRequest{}

	err := json.NewDecoder(os.Stdin).Decode(&CheckRequest)
	if err != nil {
		panic(err)
	}

	var stderrBuffer bytes.Buffer
	var stdOutBuffer bytes.Buffer
	err = ioutil.WriteFile("cert.crt", []byte(CheckRequest.Source.CertAuthority), 0644)
	err = ioutil.WriteFile("client.crt", []byte(CheckRequest.Source.ClientCert), 0644)
	err = ioutil.WriteFile("client.key", []byte(CheckRequest.Source.ClientKey), 0644)

	fmt.Fprintln(os.Stderr, CheckRequest.Source.CertAuthority)
	fmt.Fprintln(os.Stderr, CheckRequest.Source.ClientCert)
	fmt.Fprintln(os.Stderr, CheckRequest.Source.ClientKey)
	cmd := exec.Command("kubectl", "--certificate-authority", "cert.crt", "--client-key",
		"client.key", "--client-certificate", "client.crt", "--server", CheckRequest.Source.ClusterURL, "get", "pods")
	cmd.Stderr = &stderrBuffer
	cmd.Stdout = &stdOutBuffer
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, stderrBuffer.String())
	fmt.Fprintln(os.Stderr, stdOutBuffer.String())

	versions := []Version{}
	versions = append(versions, Version{
		Version: time.Now().String(),
	})

	json.NewEncoder(os.Stdout).Encode(versions)
}

//CheckRequest ...
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

//Version ...
type Version struct {
	Version string `json:"version"`
}
