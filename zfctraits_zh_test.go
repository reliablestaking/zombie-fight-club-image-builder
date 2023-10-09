package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/reliablestaking/zc-image-builder/server"
)

func TestBuildZfcZhTraits(t *testing.T) {
	_, hunters, err := server.LoadMeta("metadata")
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
	rightWeapons := make(map[string]int)
	leftWeapons := make(map[string]int)
	clothings := make(map[string]int)
	earrings := make(map[string]int)
	swag := make(map[string]int)

	for _, z := range hunters {
		updateMap(backgrounds, z.Background)
		updateMap(chains, z.Chain)
		updateMap(clothings, z.Clothing)
		updateMap(earrings, z.Earrings)
		updateMap(eyes, z.Eyes)
		updateMap(hats, z.Hat)
		updateMap(mouths, z.Mouth)
		updateMap(rightWeapons, z.RightWeapon)
		updateMap(leftWeapons, z.LeftWeapon)
		updateMap(skins, z.Skin)
		updateMap(swag, z.Swag)

	}

	for key, element := range backgrounds {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	//create csv file
	f, err := os.Create("zh_trait_rarity.csv")
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
	writeRarity(*w, leftWeapons, "left-weapon")
	writeRarity(*w, rightWeapons, "right-weapon")
	writeRarity(*w, clothings, "clothing")
	writeRarity(*w, earrings, "earring")
	writeRarity(*w, swag, "loot")
}

func writeRarity(w csv.Writer, traitMap map[string]int, traitType string) {
	for key, element := range traitMap {
		record := make([]string, 3)
		record[0] = traitType
		record[1] = key
		record[2] = strconv.Itoa(element)

		//write row
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}

func updateMap(traitMap map[string]int, value string) {
	if val, ok := traitMap[value]; ok {
		traitMap[value] = val + 1
	} else {
		traitMap[value] = 1
	}
}
