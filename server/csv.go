package server

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

//LoadZombieChainsMeta load zc meta
func LoadZombieChainsMeta(path string) (map[string]ZombieChain, error) {
	// open file
	f, err := os.Open(path + "/zombie-meta.csv")
	if err != nil {
		return nil, err
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// Parse the file
	r := csv.NewReader(f)

	zombies := make(map[string]ZombieChain)

	// skip first row
	_, err = r.Read()
	if err == io.EOF {
		return nil, err
	}
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.WithError(err).Fatal("Error reading row")
		}

		name := record[0]
		zombie := ZombieChain{
			Background: record[4],
			Weapon:     record[5],
			Skin:       record[6],
			Clothing:   record[7],
			Chain:      record[8],
			Mouth:      record[9],
			Nose:       record[10],
			Hat:        record[11],
			Eyes:       record[12],
			Earrings:   record[13],
		}

		zombies[name] = zombie
	}

	return zombies, nil
}

//LoadZombieHunterMeta load zh meta
func LoadZombieHunterMeta(path string) (map[string]ZombieHunter, error) {
	// open file
	f, err := os.Open(path + "/zombie-hunter-meta.csv")
	if err != nil {
		return nil, err
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// Parse the file
	r := csv.NewReader(f)

	zombies := make(map[string]ZombieHunter)

	// skip first row
	_, err = r.Read()
	if err == io.EOF {
		return nil, err
	}
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.WithError(err).Fatal("Error reading row")
		}

		name := record[0]
		zombie := ZombieHunter{
			Background:  record[6],
			Gender:      record[7],
			LeftWeapon:  record[10],
			RightWeapon: record[11],
			Skin:        record[8],
			Clothing:    record[9],
			Chain:       record[12],
			Mouth:       record[13],
			Hat:         record[14],
			Eyes:        record[15],
			Earrings:    record[16],
			Swag:        record[17],
		}

		zombies[name] = zombie
	}

	return zombies, nil
}
