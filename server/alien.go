package server

import (
	"bytes"
	"image/jpeg"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type (
	// AlienRarity store rarities
	AlienRarity struct {
		Backgrounds []string
		Skins       []string
		Clothes     []string
		Hats        []string
		Hands       []string
		Mouths      []string
		Eyes        []string
	}
)

//BuildAlien build an alien
func (s Server) BuildAlien(c echo.Context) (err error) {
	logrus.Infof("Building Alient image...")

	// bind incoming object
	alien := new(Alien)
	if err = c.Bind(alien); err != nil {
		logrus.WithError(err).Errorf("Error binding incoming struct")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	image, err := s.BuildAlienImage(*alien)
	if err != nil {
		logrus.WithError(err).Error("Error building alien image")
		return c.NoContent(http.StatusInternalServerError)
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, &jpeg.Options{75}); err != nil {
		logrus.WithError(err).Error("Error encoding to png")
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Stream(http.StatusOK, "image/png", buffer)
}

//GetRandomAlien build a random alient
func (s Server) GetRandomAlien(c echo.Context) (err error) {
	logrus.Infof("Building Random Alient image...")

	alien := s.CalculateRandomAlien()

	image, err := s.BuildAlienImage(alien)
	if err != nil {
		logrus.WithError(err).Error("Error building zombie chain image")
		return c.NoContent(http.StatusInternalServerError)
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, &jpeg.Options{70}); err != nil {
		logrus.WithError(err).Error("Error encoding to png")
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Stream(http.StatusOK, "image/jpeg", buffer)
}

func (s Server) CalculateRandomAlien() Alien {
	alien := Alien{}

	disallowedEyes := make([]string, 0)

	alien.Hat = pickRandomTrait(s.AlienRarity.Hats)

	// hat rules
	if alien.Hat == "Egg-Layer" {
		// not hat or eyes since this covers whole face
		alien.Mouth = "None"
		alien.Eyes = "None"
	} else {
		alien.Mouth = pickRandomTrait(s.AlienRarity.Mouths)

		//otherwise if tentacles, excludes some eyes
		if alien.Hat == "Tentacles" {
			disallowedEyes = append(disallowedEyes, "Night-Vision-Goggles")
			disallowedEyes = append(disallowedEyes, "Steampunk")
		}

		// pick eyes with potentially disallowed
		alien.Eyes = pickRandomTraitWithDisallowed(s.AlienRarity.Eyes, disallowedEyes)

		// if eyes are one of these, then certain hats are disallowed, so need to repick
		if alien.Eyes == "Shades" || alien.Eyes == "Goggles" || alien.Eyes == "Space-Glasses" || alien.Eyes == "3D" || alien.Eyes == "Aviator" || alien.Eyes == "Night-Vision-Goggles" || alien.Eyes == "Steampunk" || alien.Eyes == "Glasses" || alien.Eyes == "Visor-Two" {
			// need to repick hat
			disallowedHats := make([]string, 0)
			disallowedHats = append(disallowedHats, "Psychic-Helmet")
			disallowedHats = append(disallowedHats, "Melvin-The-Martian")
			disallowedHats = append(disallowedHats, "Cronenberg")
			disallowedHats = append(disallowedHats, "Illogical")
			disallowedHats = append(disallowedHats, "Headphones")
			disallowedHats = append(disallowedHats, "Zith-Lord")
			disallowedHats = append(disallowedHats, "Pharaoh")
			disallowedHats = append(disallowedHats, "Hunter-Dreadlocks")
			disallowedHats = append(disallowedHats, "Fins")
			disallowedHats = append(disallowedHats, "Egg-Layer")
			disallowedHats = append(disallowedHats, "Tentacles")

			alien.Hat = pickRandomTraitWithDisallowed(s.AlienRarity.Hats, disallowedHats)
		}
	}

	alien.Clothes = pickRandomTrait(s.AlienRarity.Clothes)
	// clothes rules
	if alien.Clothes == "Rainbow-Wings" {
		// if mouth isn't none, then repick
		if alien.Mouth != "None" {
			disallowedMouths := make([]string, 0)
			disallowedMouths = append(disallowedMouths, "Smoking-Cigar")
			disallowedMouths = append(disallowedMouths, "Parasite")

			alien.Mouth = pickRandomTraitWithDisallowed(s.AlienRarity.Mouths, disallowedMouths)
		}

		// repick hat
		if alien.Hat != "Egg-Layer" && alien.Hat != "Tentacles" {
			disallowedHats := make([]string, 0)
			disallowedHats = append(disallowedHats, "Hunter-Dreadlocks")
			disallowedHats = append(disallowedHats, "Pharaoh")
			disallowedHats = append(disallowedHats, "Zith-Lord")
			disallowedHats = append(disallowedHats, "Cronenberg")
			disallowedHats = append(disallowedHats, "Melvin-The-Martian")
			disallowedHats = append(disallowedHats, "Pschic-Helmet")
			disallowedHats = append(disallowedHats, "Tubes")
			disallowedHats = append(disallowedHats, "Egg-Layer")
			disallowedHats = append(disallowedHats, "Tentacles")

			alien.Hat = pickRandomTraitWithDisallowed(s.AlienRarity.Hats, disallowedHats)
		}
	}

	alien.Background = pickRandomTrait(s.AlienRarity.Backgrounds)
	alien.Hand = pickRandomTrait(s.AlienRarity.Hands)
	alien.Skin = pickRandomTrait(s.AlienRarity.Skins)

	return alien
}

func pickRandomTrait(names []string) string {
	randomName := names[rand.Intn(len(names))]
	return randomName
}

func pickRandomTraitWithDisallowed(names []string, disallowed []string) string {
	for {
		randomName := names[rand.Intn(len(names))]

		if contains(disallowed, randomName) {
			continue
		}

		return randomName
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
