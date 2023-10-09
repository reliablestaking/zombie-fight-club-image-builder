package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"testing"

	"github.com/reliablestaking/zc-image-builder/server"
)

func TestZfcImageBuilder(t *testing.T) {
	zombies, hunters, err := server.LoadMeta("metadata")
	if err != nil {
		t.Fatal("Error loading zc meta")
	}

	// init server
	s := server.Server{
		ZombieChainMeta:  zombies,
		ZombieHunterMeta: hunters,
	}

	zombieFight := server.ZombieFight{
		Background:          "Graveyard",
		Vs:                  "VS",
		ZombieRecord:        "1-0",
		HunterRecord:        "1-0",
		ZombieChainLifeBar:  10,
		ZombieHunterLifeBar: 10,
		ZombieChain:         "ZombieChains00001",
		ZombieHunter:        "ZombieHunter00001",
	}

	// build every zfc combination
	t.Logf("Building zombie chains")
	for i := 1; i <= 10000; i++ {
		zombieName := fmt.Sprintf("ZombieChains%05d", i)
		hunterName := fmt.Sprintf("ZombieHunter%05d", i)

		t.Logf("Buildling fight with zombie %s and hunter %s", zombieName, hunterName)
		zombieFight.ZombieChain = zombieName
		zombieFight.ZombieBeatup = false
		zombieFight.ZombieHunter = hunterName
		zombieFight.HunterBeatup = false
		fightImage, _, err := s.BuildZombieFightImage(zombieFight)
		if err != nil {
			t.Fatalf("Error building fight with zombie %s with err %v", zombieName, err)
		}

		//resize
		fightImageResized := server.ResizeImage(*fightImage, 600, 337)
		f, err := os.Create(fmt.Sprintf("examples/Fight%05d.jpeg", i))
		if err != nil {
			t.Fatalf("Error saving fight with zombie %s with err %v", zombieName, err)
		}
		defer f.Close()
		err = jpeg.Encode(f, fightImageResized, &jpeg.Options{Quality: 70})
		if err != nil {
			t.Fatalf("Error encoding fight with zombie %s with err %v", zombieName, err)
		}

		t.Logf("Buildling beatup fight with zombie %s", zombieName)
		zombieFight.ZombieBeatup = true
		zombieFight.HunterBeatup = true
		_, _, err = s.BuildZombieFightImage(zombieFight)
		if err != nil {
			t.Fatalf("Error building fight with zombie %s with err %v", zombieName, err)
		}

	}
}
