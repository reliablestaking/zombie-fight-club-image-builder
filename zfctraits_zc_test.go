package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/reliablestaking/zc-image-builder/server"
)

func TestBuildZfcTraits(t *testing.T) {
	zombies, _, err := server.LoadMeta("metadata")
	if err != nil {
		t.Fatal("Error loading zc meta")
	}

	backgrounds := make(map[string]int)
	hats := make(map[string]int)
	eyes := make(map[string]int)
	noses := make(map[string]int)
	skins := make(map[string]int)
	mouths := make(map[string]int)
	chains := make(map[string]int)
	weapons := make(map[string]int)
	clothings := make(map[string]int)
	earrings := make(map[string]int)

	for _, z := range zombies {
		updateMap(backgrounds, z.Background)
		updateMap(hats, z.Hat)
		updateMap(eyes, z.Eyes)
		updateMap(noses, z.Nose)
		updateMap(skins, z.Skin)
		updateMap(mouths, z.Mouth)
		updateMap(chains, z.Chain)
		updateMap(weapons, z.Weapon)
		updateMap(clothings, z.Clothing)
		updateMap(earrings, z.Earrings)
	}

	for key, element := range backgrounds {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	//create csv file
	f, err := os.Create("zc_trait_rarity.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()

	writeRarity(*w, backgrounds, "background")
	writeRarity(*w, hats, "hat")
	writeRarity(*w, eyes, "eyes")
	writeRarity(*w, noses, "nose")
	writeRarity(*w, skins, "skin")
	writeRarity(*w, mouths, "mouth")
	writeRarity(*w, chains, "chain")
	writeRarity(*w, weapons, "weapon")
	writeRarity(*w, clothings, "clothing")
	writeRarity(*w, earrings, "earring")

}
