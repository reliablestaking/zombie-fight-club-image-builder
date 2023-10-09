package main

import (
	"fmt"
	"image/jpeg"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/reliablestaking/zc-image-builder/server"
	"github.com/sirupsen/logrus"
)

func TestAlienImageBuilder(t *testing.T) {
	zombies, hunters, err := server.LoadMeta("metadata")
	if err != nil {
		t.Fatal("Error loading zc meta")
	}

	ar, err := server.LoadAlienRarity()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading alien rarity")
	}

	// init server
	s := server.Server{
		ZombieChainMeta:  zombies,
		ZombieHunterMeta: hunters,
		AlienRarity:      ar,
	}

	rand.Seed(time.Now().UnixNano())

	// build 10000 aliens
	t.Logf("Building aliens")
	for i := 1; i <= 5000; i++ {

		alien := s.CalculateRandomAlien()
		image, err := s.BuildAlienImage(alien)
		if err != nil {
			t.Fatalf("Error building zombie chain image %v", err)
		}

		//resize
		alienResized := server.ResizeImage(*image, 512, 512)
		f, err := os.Create(fmt.Sprintf("alien-examples/Alien%05d.jpeg", i))
		if err != nil {
			t.Fatalf("Error saving fight with alien %v with err %v", alien, err)
		}
		defer f.Close()
		err = jpeg.Encode(f, alienResized, &jpeg.Options{Quality: 70})
		if err != nil {
			t.Fatalf("Error encoding fight with alien %v with err %v", alien, err)
		}

	}
}
