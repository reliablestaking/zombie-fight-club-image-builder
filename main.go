package main

// @title ZC Image Builder API
// @version 1.0
// @description This is the API to build ZC Images

import (
	"math/rand"
	"time"

	"github.com/reliablestaking/zc-image-builder/server"
	"github.com/sirupsen/logrus"
)

var (
	sha1ver   string // sha1 revision used to build the program
	buildTime string // when the executable was built
)

func main() {
	//seed random
	rand.Seed(time.Now().UnixNano())

	// load zombie meta
	zombies, hunters, err := server.LoadMeta("metadata")
	if err != nil {
		logrus.WithError(err).Fatal("Error loading zc meta")
	}

	ar, err := server.LoadAlienRarity()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading alien rarity")
	}

	backgrounds, err := server.LoadBackgroundRarity()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading alien rarity")
	}

	// init server
	server := server.Server{
		Sha1ver:          sha1ver,
		BuildTime:        buildTime,
		ZombieChainMeta:  zombies,
		ZombieHunterMeta: hunters,
		AlienRarity:      ar,
		ZfcBackgrounds:   backgrounds,
	}

	// start server
	server.Start()
}
