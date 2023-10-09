package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/mitchellh/hashstructure/v2"
	"github.com/reliablestaking/zc-image-builder/server"
	"github.com/sirupsen/logrus"
)

func TestAlienCsvBuilder(t *testing.T) {
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

	//create csv file
	f, err := os.Create("aliens.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()

	hashes := make([]uint64, 0)
	collisionCount := 0

	// build 10000 aliens
	t.Logf("Building aliens")
	i := 1
	for i <= 1000000 {
		alien := s.CalculateRandomAlien()

		// get hash
		hash, err := hashstructure.Hash(alien, hashstructure.FormatV2, nil)
		if err != nil {
			panic(err)
		}

		if !contains(hashes, hash) {

			hashes = append(hashes, hash)

			record := make([]string, 9)
			record[0] = fmt.Sprintf("Alien%d", i)
			record[1] = fmt.Sprintf("Alien #%d", i)
			record[2] = alien.Background
			record[3] = alien.Skin
			record[4] = alien.Clothes
			record[5] = alien.Eyes
			record[6] = alien.Mouth
			record[7] = alien.Hand
			record[8] = alien.Hat

			//write row
			if err := w.Write(record); err != nil {
				log.Fatalln("error writing record to file", err)
			}

			i++
		} else {
			logrus.Infof("Hash collision at index %d", i)
			collisionCount++
		}
	}

	logrus.Infof("Finished with %d collision", collisionCount)
}

func contains(s []uint64, e uint64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
