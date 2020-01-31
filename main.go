package main

import (
	"context"
	"fmt"
	"log"
	"os"

	tfe "github.com/hashicorp/go-tfe"
)

const ()

func main() {

	var (
		TOKEN         = os.Getenv("TF_ORG_TOKEN")
		WORKPLACE     = os.Getenv("TF_WORKPLACE")
		Orgramization = os.Getenv("TF_ORGRAN")
	)
	config := tfe.DefaultConfig()
	config.Token = TOKEN

	cli, err := tfe.NewClient(config)
	handleErr(err)
	back := context.Background()
	op := tfe.StateVersionListOptions{
		Organization: &Orgramization,
		Workspace:    &WORKPLACE}

	list, err := cli.StateVersions.List(back, op)
	handleErr(err)
	ver := list.Items[0] //Latest State update

	bs, err := cli.StateVersions.Download(back, ver.DownloadURL)
	handleErr(err)
	fmt.Printf("%s", bs)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
