package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	tfe "github.com/hashicorp/go-tfe"
)

const ()

func main() {

	token := flag.String("org_token", "deodoekdoekdo", "Organization Token")
	flag.Parse()
	for !flag.Parsed() {

	}

	var (
		TOKEN         = *token
		WORKPLACE     = "momo-land"
		Orgramization = "Momo"
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
