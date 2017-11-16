package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {

	if len(os.Args) < 2 {
	}

	request := OutRequest{}
	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		panic(err)
	}

	kubectl := []string{"--certificate-authority", "cert.crt", "--client-key", "client.key",
		"--client-certificate", "client.crt", "--server", request.Source.ClusterURL, "apply", "-f", request.Params.File}

	var stderrBuffer bytes.Buffer
	var stdOutBuffer bytes.Buffer
	err = ioutil.WriteFile("cert.crt", []byte(request.Source.CertAuthority), 0644)
	err = ioutil.WriteFile("client.crt", []byte(request.Source.ClientCert), 0644)
	err = ioutil.WriteFile("client.key", []byte(request.Source.ClientKey), 0644)

	cmd := exec.Command("kubectl", kubectl...)
	cmd.Stderr = &stderrBuffer
	cmd.Stdout = &stdOutBuffer
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(stderrBuffer.String())
	fmt.Println(stdOutBuffer.String())

	//versions := []Version{}
	//versions = append(versions, Version{
	//	Version: time.Now().String(),
	//})

	//json.NewEncoder(os.Stdout).Encode(versions)
}

//OutRequest ...
type OutRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
	Params  Params  `json:"params"`
}

//Params ...
type Params struct {
	File string `json:"file"`
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
