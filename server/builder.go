package server

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type (
	//ZombieChain struct to store zombie chain
	ZombieChain struct {
		Background string `json:"background"`
		Hat        string `json:"hat"`
		Eyes       string `json:"eyes"`
		Nose       string `json:"nose"`
		Skin       string `json:"skin"`
		Mouth      string `json:"mouth"`
		Chain      string `json:"chain"`
		Weapon     string `json:"weapon"`
		Clothing   string `json:"clothing"`
		Earrings   string `json:"earrings"`
	}

	//ZombieHunter struct to store zombie hunter
	ZombieHunter struct {
		Background  string `json:"background"`
		Gender      string `json:"gender"`
		Hat         string `json:"hat"`
		Eyes        string `json:"eyes"`
		Skin        string `json:"skin"`
		Mouth       string `json:"mouth"`
		Chain       string `json:"chain"`
		LeftWeapon  string `json:"leftWeapon"`
		RightWeapon string `json:"rightWeapon"`
		Clothing    string `json:"clothing"`
		Earrings    string `json:"earrings"`
		Swag        string `json:"swag"`
	}

	//ZombieFight struct to store fight
	ZombieFight struct {
		Background          string `json:"background"`
		ZombieChain         string `json:"zombieChain"`
		ZombieChainLifeBar  int    `json:"zcLifeBar"`
		ZombieHunter        string `json:"zombieHunter"`
		ZombieHunterLifeBar int    `json:"zhLifeBar"`
		Vs                  string `json:"vs"`
		ZombieRecord        string `json:"zombieRecord"`
		HunterRecord        string `json:"hunterRecord"`
		ZombieKO            bool   `json:"zombieKo"`
		ZombieBeatup        bool   `json:"zombieBeatup"`
		HunterKO            bool   `json:"hunterKo"`
		HunterBeatup        bool   `json:"hunterBeatup"`
		Width               int    `json:"width"`
		Height              int    `json:"height"`
	}

	//Alien struct to store alien
	Alien struct {
		Background string `json:"background"`
		Skin       string `json:"skin"`
		Clothes    string `json:"clothes"`
		Hat        string `json:"hat"`
		Hand       string `json:"hand"`
		Mouth      string `json:"mouth"`
		Eyes       string `json:"eyes"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
	}
)

//GetZombieChain build a zombie chain image
func (s Server) GetZombieChain(c echo.Context) (err error) {
	logrus.Infof("Building Zombie Chain image...")

	// bind incoming object
	name := c.Param("name")
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	omitBackground := false
	omitBackgroundParam := c.QueryParam("omitBackground")
	if omitBackgroundParam == "true" {
		omitBackground = true
	}

	// find meta for that name
	zcMeta, ok := s.ZombieChainMeta[name]
	if !ok {
		logrus.Errorf("No zc meta found for name %s", name)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	image, err := s.buildZombieChainImage(zcMeta, omitBackground)
	if err != nil {
		logrus.WithError(err).Error("Error building zombie chain image")
		return c.NoContent(http.StatusInternalServerError)
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, image); err != nil {
		logrus.WithError(err).Error("Error encoding to png")
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Stream(http.StatusOK, "image/png", buffer)
}

//GetZombieHunter build a zombie hunter image
func (s Server) GetZombieHunter(c echo.Context) (err error) {
	logrus.Infof("Building Zombie Hunter image...")

	// bind incoming object
	name := c.Param("name")
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// find meta for that name
	zhMeta, ok := s.ZombieHunterMeta[name]
	if !ok {
		logrus.Errorf("No zh meta found for name %s", name)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	image, err := s.buildZombieHunterImage(zhMeta)
	if err != nil {
		logrus.WithError(err).Error("Error building zombie chain image")
		return c.NoContent(http.StatusInternalServerError)
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, image); err != nil {
		logrus.WithError(err).Error("Error encoding to png")
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Stream(http.StatusOK, "image/png", buffer)
}

//BuildZombieFight build a zombie fight
func (s Server) BuildZombieFight(c echo.Context) (err error) {
	logrus.Infof("Building Zombie Fight image...")

	// bind incoming object
	zombieFight := new(ZombieFight)
	if err = c.Bind(zombieFight); err != nil {
		logrus.WithError(err).Errorf("Error binding incoming struct")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	image, background, err := s.BuildZombieFightImage(*zombieFight)
	if err != nil {
		logrus.WithError(err).Error("Error building zombie chain image")
		return c.NoContent(http.StatusInternalServerError)
	}

	c.Response().Header().Set("Background", background)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, &jpeg.Options{75}); err != nil {
		logrus.WithError(err).Error("Error encoding to png")
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Stream(http.StatusOK, "image/png", buffer)
}

// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
