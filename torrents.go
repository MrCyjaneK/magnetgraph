package main

import (
	"log"
	"os"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/storage"
)

var cl *torrent.Client
var dir string
var torrents = make(map[string]*torrent.Torrent)

func tinit() {
	var err error
	dir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dir = dir + "/.magnetgraph"
	os.MkdirAll(dir, 0750)
	config := torrent.NewDefaultClientConfig()
	config.DataDir = dir
	config.Seed = true
	config.HTTPUserAgent = "Magnetgraph v0.0.0 (posts over torrent)"
	config.ExtendedHandshakeClientVersion = "magnetgraph 20210802"
	config.Bep20 = "-MG0000-"
	config.DefaultStorage = storage.NewFileByInfoHash(dir)
	cl, err = torrent.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
}
