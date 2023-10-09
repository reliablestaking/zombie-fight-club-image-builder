package server

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"strings"

	"github.com/nfnt/resize"
	"github.com/reliablestaking/zc-image-builder/util"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"golang.org/x/image/font/opentype"
)

func (s Server) buildZombieChainImage(zombieChain ZombieChain, omitBackground bool) (*image.RGBA, error) {
	baseImage := image.NewRGBA(image.Rect(0, 0, 1080, 1080))

	if !omitBackground {
		background, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/background/%s.jpg", formatName(zombieChain.Background)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, background)
	}

	weapon, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/weapon/%s.png", formatName(zombieChain.Weapon)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, weapon)

	skin, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/skin/%s.png", formatName(zombieChain.Skin)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, skin)

	clothing, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/clothing/%s.png", formatName(zombieChain.Clothing)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, clothing)

	necklace, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/necklace/%s.png", formatName(zombieChain.Chain)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, necklace)

	mouth, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/mouth/%s.png", formatName(zombieChain.Mouth)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, mouth)

	nose, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/nose/%s.png", formatName(zombieChain.Nose)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, nose)

	hat, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/hat/%s.png", formatName(zombieChain.Hat)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, hat)

	eye, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/eyes/%s.png", formatName(zombieChain.Eyes)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, eye)

	earring, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiechains/earrings/%s.png", formatName(zombieChain.Earrings)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, earring)

	return baseImage, nil
}

func (s Server) buildZfcZombieChainImage(zombieChain ZombieChain, beatup bool) (*image.RGBA, error) {
	baseImage := image.NewRGBA(image.Rect(0, 0, 600, 675))

	weapon, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/weapon/%s.png", formatName(zombieChain.Weapon)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, weapon)

	skin, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/skin/%s.png", formatName(zombieChain.Skin)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, skin)

	clothing, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/clothing/%s.png", formatName(zombieChain.Clothing)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, clothing)

	necklace, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/necklace/%s.png", formatName(zombieChain.Chain)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, necklace)

	hat, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/hat/%s.png", formatName(zombieChain.Hat)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, hat)

	if beatup {
		mouth, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/mouth-beatup/%s-Beatup.png", formatName(zombieChain.Mouth)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, mouth)
	} else {
		mouth, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/mouth/%s.png", formatName(zombieChain.Mouth)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, mouth)
	}

	nose, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/nose/%s.png", formatName(zombieChain.Nose)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, nose)

	if beatup {
		eye, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/eyes-beatup/%s-Beatup.png", formatName(zombieChain.Eyes)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, eye)
	} else {
		eye, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/eyes/%s.png", formatName(zombieChain.Eyes)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, eye)
	}

	if zombieChain.Eyes != "Visor Down" {
		earring, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiechains/earrings/%s.png", formatName(zombieChain.Earrings)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, earring)
	}

	return baseImage, nil
}

func (s Server) buildZombieHunterImage(zombieHunter ZombieHunter) (*image.RGBA, error) {
	baseImage := image.NewRGBA(image.Rect(0, 0, 1080, 1080))

	gender := strings.ToLower(zombieHunter.Gender)

	skinFolder := "skin"
	if zombieHunter.RightWeapon == "None" {
		skinFolder = "skin-fist"
	}
	skin, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/%s/%s.png", gender, skinFolder, formatName(zombieHunter.Skin)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, skin)

	clothing, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/clothing/%s.png", gender, formatName(zombieHunter.Clothing)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, clothing)

	leftWeapon, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/left-weapon/%s.png", gender, formatName(zombieHunter.LeftWeapon)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, leftWeapon)

	rightWeapon, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/right-weapon/%s.png", gender, formatName(zombieHunter.RightWeapon)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, rightWeapon)

	necklace, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/necklace/%s.png", gender, formatName(zombieHunter.Chain)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, necklace)

	mouth, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/mouth/%s.png", gender, formatName(zombieHunter.Mouth)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, mouth)

	hat, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/hat/%s.png", gender, formatName(zombieHunter.Hat)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, hat)

	eye, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/eyes/%s.png", gender, formatName(zombieHunter.Eyes)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, eye)

	earring, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zombiehunters/%s/earrings/%s.png", gender, formatName(zombieHunter.Earrings)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, earring)

	return baseImage, nil
}

func (s Server) buildZfcZombieHunterImage(zombieHunter ZombieHunter, beatup bool) (*image.RGBA, error) {
	baseImage := image.NewRGBA(image.Rect(0, 0, 600, 675))

	gender := strings.ToLower(zombieHunter.Gender)

	skinFolder := "skin"
	if zombieHunter.RightWeapon == "None" {
		skinFolder = "skin-fist"
	}
	skin, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/%s/%s.png", gender, skinFolder, formatName(zombieHunter.Skin)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, skin)

	clothing, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/clothing/%s.png", gender, formatName(zombieHunter.Clothing)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, clothing)

	leftWeapon, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/left-weapon/%s.png", gender, formatName(zombieHunter.LeftWeapon)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, leftWeapon)

	rightWeapon, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/right-weapon/%s.png", gender, formatName(zombieHunter.RightWeapon)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, rightWeapon)

	necklace, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/necklace/%s.png", gender, formatName(zombieHunter.Chain)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, necklace)

	hat, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/hat/%s.png", gender, formatName(zombieHunter.Hat)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, hat)

	if beatup {
		mouth, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/mouth-beatup/%s-Beatup.png", gender, formatName(zombieHunter.Mouth)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, mouth)
	} else {
		mouth, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/mouth/%s.png", gender, formatName(zombieHunter.Mouth)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, mouth)
	}

	if beatup {
		eye, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/eyes-beatup/%s-Beatup.png", gender, formatName(zombieHunter.Eyes)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, eye)
	} else {
		eye, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/eyes/%s.png", gender, formatName(zombieHunter.Eyes)))
		if err != nil {
			return nil, err
		}
		util.DrawOver(baseImage, eye)
	}

	earring, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc-zombiehunters/%s/earrings/%s.png", gender, formatName(zombieHunter.Earrings)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, earring)

	return baseImage, nil
}

func (s Server) BuildZombieFightImage(zombieFight ZombieFight) (*image.RGBA, string, error) {
	baseImage := image.NewRGBA(image.Rect(0, 0, 1200, 675))

	if zombieFight.Background == "" {
		//choose random if not provided
		zombieFight.Background = pickRandomTrait(s.ZfcBackgrounds)
	}

	background, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc/background/%s.jpg", formatName(zombieFight.Background)))
	if err != nil {
		return nil, "", err
	}
	util.DrawOver(baseImage, background)

	recordBox, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc/recordbox/%s.png", "Record-Box"))
	if err != nil {
		return nil, "", err
	}
	util.DrawOver(baseImage, recordBox)

	// find meta for zombie chain
	zcMeta, ok := s.ZombieChainMeta[zombieFight.ZombieChain]
	if !ok {
		return nil, "", fmt.Errorf("No zc meta found for name %s", zombieFight.ZombieChain)
	}
	zombieChain, err := s.buildZfcZombieChainImage(zcMeta, zombieFight.ZombieBeatup)
	if err != nil {
		return nil, "", err
	}
	util.DrawOverLocation(baseImage, zombieChain, 0, 0)

	// find meta for zombie hunter
	zhMeta, ok := s.ZombieHunterMeta[zombieFight.ZombieHunter]
	if !ok {
		return nil, "", fmt.Errorf("No zh meta found for name %s", zombieFight.ZombieHunter)
	}
	zombieHunter, err := s.buildZfcZombieHunterImage(zhMeta, zombieFight.HunterBeatup)
	if err != nil {
		return nil, "", err
	}
	util.DrawOverLocation(baseImage, zombieHunter, -600, 0)

	//add life bars
	zcLifeBar, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc/lifebars/Life Bar ZC %d.png", zombieFight.ZombieChainLifeBar))
	if err != nil {
		return nil, "", err
	}
	util.DrawOver(baseImage, zcLifeBar)

	zhLifeBar, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc/lifebars/Life Bar ZH %d.png", zombieFight.ZombieHunterLifeBar))
	if err != nil {
		return nil, "", err
	}
	util.DrawOver(baseImage, zhLifeBar)

	//add logo
	zfcLogo, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc/logos/ZFC Logo.png"))
	if err != nil {
		return nil, "", err
	}
	util.DrawOver(baseImage, zfcLogo)

	//add vs
	zfcVs, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc/vs/%s.png", zombieFight.Vs))
	if err != nil {
		return nil, "", err
	}
	util.DrawOver(baseImage, zfcVs)

	//add KO
	if zombieFight.ZombieKO {
		koLeft, err := util.GetImageFromFileNoCache("images/zfc/ko/KO Left.png")
		if err != nil {
			return nil, "", err
		}
		util.DrawOver(baseImage, koLeft)
	}
	if zombieFight.HunterKO {
		koRight, err := util.GetImageFromFileNoCache("images/zfc/ko/KO Right.png")
		if err != nil {
			return nil, "", err
		}
		util.DrawOver(baseImage, koRight)
	}

	//add record
	startingOffset := -445
	if len(zombieFight.ZombieRecord) == 3 {
		startingOffset = -473
	} else if len(zombieFight.ZombieRecord) == 5 {
		startingOffset = -459
	} else if len(zombieFight.ZombieRecord) >= 7 {
		startingOffset = -445
	} else {
		return nil, "", fmt.Errorf("Record is an invalid number of digits")
	}
	offset := 0

	// don't let go over 999...
	zcWins := ""
	zcLosses := ""
	zcRecordSplit := strings.Split(zombieFight.ZombieRecord, "-")
	if len(zcRecordSplit[0]) >= 4 {
		zcWins = "999"
	} else {
		zcWins = zcRecordSplit[0]
	}
	if len(zcRecordSplit[1]) >= 4 {
		zcLosses = "999"
	} else {
		zcLosses = zcRecordSplit[1]
	}
	zcRecord := fmt.Sprintf("%s-%s", zcWins, zcLosses)

	for _, d := range zcRecord {
		numberImage, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc/numbers/%s.png", string(d)))
		if err != nil {
			return nil, "", err
		}
		if string(d) == "-" {
			util.DrawOverLocation(baseImage, numberImage, startingOffset+offset-2, -62)
		} else {
			util.DrawOverLocation(baseImage, numberImage, startingOffset+offset, -52)
		}

		offset -= 14
	}

	startingOffset = -660
	if len(zombieFight.HunterRecord) == 3 {
		startingOffset = -688
	} else if len(zombieFight.HunterRecord) == 5 {
		startingOffset = -674
	} else if len(zombieFight.HunterRecord) >= 7 {
		startingOffset = -660
	} else {
		return nil, "", fmt.Errorf("Record is an invalid number of digits")
	}
	offset = 0

	// don't let go over 999...
	zhWins := ""
	zhLosses := ""
	zhRecordSplit := strings.Split(zombieFight.HunterRecord, "-")
	if len(zhRecordSplit[0]) >= 4 {
		zhWins = "999"
	} else {
		zhWins = zhRecordSplit[0]
	}
	if len(zhRecordSplit[1]) >= 4 {
		zhLosses = "999"
	} else {
		zhLosses = zhRecordSplit[1]
	}
	zhRecord := fmt.Sprintf("%s-%s", zhWins, zhLosses)

	for _, d := range zhRecord {
		numberImage, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/zfc/numbers/%s.png", string(d)))
		if err != nil {
			return nil, "", err
		}
		if string(d) == "-" {
			util.DrawOverLocation(baseImage, numberImage, startingOffset+offset-2, -62)
		} else {
			util.DrawOverLocation(baseImage, numberImage, startingOffset+offset, -52)
		}

		offset -= 14
	}

	// add message
	message := determineIfComboMessage(zcMeta, zhMeta)

	if message == "" {
		if zombieFight.ZombieChainLifeBar > zombieFight.ZombieHunterLifeBar {
			message = determineIfZombieWinMessage(zcMeta)
		} else if zombieFight.ZombieHunterLifeBar > zombieFight.ZombieChainLifeBar {
			message = determineIfHunterWinMessage(zhMeta)
		}
	}

	if message != "" {
		zfcMsg, err := util.GetImageFromFileNoCache(message)
		if err != nil {
			return nil, "", err
		}
		util.DrawOver(baseImage, zfcMsg)
	}

	//check if need to resize
	if zombieFight.Width != 0 && zombieFight.Height != 0 {
		resizedImage := ResizeImage(*baseImage, uint(zombieFight.Width), uint(zombieFight.Height))
		b := resizedImage.Bounds()
		m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(m, m.Bounds(), resizedImage, b.Min, draw.Src)
		return m, "", nil
	}

	return baseImage, zombieFight.Background, nil
}

func determineIfComboMessage(zcMeta ZombieChain, zhMeta ZombieHunter) string {
	if zcMeta.Mouth == "Gas Mask" && zhMeta.Mouth == "Gas Mask" {
		return "images/zfc/messages/both/_Whoever-Smelt-It,--Dealt-It!.png"
	} else if zcMeta.Clothing == "Oktoberfest" && zhMeta.LeftWeapon == "Beer" {
		return "images/zfc/messages/both/_Beer-Brawl!.png"
	} else if zcMeta.Hat == "Bull Horns" && zhMeta.Hat == "Horns" {
		return "images/zfc/messages/both/Bull-Fight.png"
	} else if zcMeta.Mouth == "Hockey" && zhMeta.Mouth == "Hockey Teeth" {
		return "images/zfc/messages/both/_Trip-to-the-Dentist.png"
	} else if zcMeta.Skin == "Tattoo" && strings.Contains(zhMeta.Skin, "Tattoo") {
		return "images/zfc/messages/both/_Inked.png"
	} else if zcMeta.Hat == "Crown" && strings.Contains(zhMeta.Hat, "Crown") {
		return "images/zfc/messages/both/_Kings-and-Queens.png"
	} else if zcMeta.Earrings == "Earbud" && zhMeta.Earrings == "Bluetooth" {
		return "images/zfc/messages/both/_Business-Meeting--Brawl!.png"
	} else if zcMeta.Skin == "Decay" && zhMeta.LeftWeapon == "Z Virus" {
		return "images/zfc/messages/both/_The-Cure.png"
	} else if zcMeta.Mouth == "Tongue" && zhMeta.Mouth == "Tongue" {
		return "images/zfc/messages/both/_French-Kiss.png"
	} else if zcMeta.Hat == "Mohawk" && zhMeta.Hat == "Punk" {
		return "images/zfc/messages/both/_This-Fight's--for-Vohn!.png"
	} else if zcMeta.Eyes == "Night Vision" && zhMeta.Eyes == "Night Vision Goggles" {
		return "images/zfc/messages/both/_After-Dark.png"
	} else if zcMeta.Eyes == "Laser Beam" && zhMeta.Eyes == "Laser" {
		return "images/zfc/messages/both/_Laser-Eyes!!!.png"
	} else if zcMeta.Eyes == "3D" && zhMeta.Eyes == "3d Glasses" {
		return "images/zfc/messages/both/_Movie-Massacre.png"
	} else if zcMeta.Eyes == "Cyclops" && zhMeta.Eyes == "Cyclops" {
		return "images/zfc/messages/both/_Fight-of--the-Cyclops!.png"
	} else if zcMeta.Skin == "Cyborg" && zhMeta.Skin == "Cyborg" {
		return "images/zfc/messages/both/_Robot-War!.png"
	} else if zcMeta.Clothing == "None" && zhMeta.Clothing == "None" {
		return "images/zfc/messages/both/_Up-Close-and-Personal.png"
	} else if zcMeta.Hat == "Santa" && zhMeta.Hat == "Santa" {
		return "images/zfc/messages/both/_Who's-Your-Santa!-.png"
	} else if zcMeta.Hat == "Headphones" && zhMeta.Hat == "Headphones" {
		return "images/zfc/messages/both/_Your-Music-Sucks!.png"
	} else if zcMeta.Weapon == "None" && zhMeta.LeftWeapon == "None" && zhMeta.RightWeapon == "None" {
		return "images/zfc/messages/both/_Bare-Knuckle-Brawl.png"
	} else if zcMeta.Clothing == "Cheerleader" && zhMeta.Clothing == "Cheerleader" {
		return "images/zfc/messages/both/_Let's-Go-Zombats!.png"
	} else if zcMeta.Hat == "Astronaut" && zhMeta.Hat == "Space Helmet" {
		return "images/zfc/messages/both/_Fight-...-In-Space!.png"
	} else if zcMeta.Weapon == "Angel Fully Fledged" && zhMeta.Clothing == "Demon" {
		return "images/zfc/messages/both/_Armageddon--Face-Off!.png"
	} else if zcMeta.Clothing == "Superzombie" && zhMeta.Clothing == "Superhero" {
		return "images/zfc/messages/both/_Super-Fight!.png"
	}

	return ""
}

func determineIfHunterWinMessage(zhMeta ZombieHunter) string {
	if zhMeta.Hat == "Jar Brain" {
		return "images/zfc/messages/hunter-wins/Big Brain FTW.png"
	} else if zhMeta.LeftWeapon == "Diamond Hands" {
		return "images/zfc/messages/hunter-wins/HODL for the Win!.png"
	} else if zhMeta.Mouth == "Vampire Fangs" {
		return "images/zfc/messages/hunter-wins/Drained!.png"
	} else if zhMeta.Eyes == "Visor" {
		return "images/zfc/messages/hunter-wins/Khaaannnn!.png"
	} else if zhMeta.RightWeapon == "Boomstick" {
		return "images/zfc/messages/hunter-wins/Say hello to my Boomstick!.png"
	} else if zhMeta.Hat == "Helmet" {
		return "images/zfc/messages/hunter-wins/I love the smell  of burnt Zombies  in the morning!.png"
	} else if zhMeta.Hat == "Viking" {
		return "images/zfc/messages/hunter-wins/To Valhalla!.png"
	} else if zhMeta.RightWeapon == "Banana" {
		return "images/zfc/messages/hunter-wins/Potassium FTW.png"
	} else if zhMeta.Hat == "Scientist" {
		return "images/zfc/messages/hunter-wins/'Schwifty Victory'.png"
	} else if zhMeta.Swag == "Coins" {
		return "images/zfc/messages/hunter-wins/Don't Touch My Coins!.png"
	} else if zhMeta.Swag == "Retro Computer" {
		return "images/zfc/messages/hunter-wins/Crypto Trader FTW.png"
	} else if zhMeta.Swag == "Rocket" {
		return "images/zfc/messages/hunter-wins/To the Moon!.png"
	} else if zhMeta.Swag == "Hourglass" {
		return "images/zfc/messages/hunter-wins/Time Travel.png"
	} else if zhMeta.Swag == "Garden Gnome" {
		return "images/zfc/messages/hunter-wins/Don't Mess with  my Gnome!.png"
	} else if zhMeta.RightWeapon == "Light Gun" {
		return "images/zfc/messages/hunter-wins/8Bit Victory.png"
	} else if zhMeta.RightWeapon == "Powerful Glove" {
		return "images/zfc/messages/hunter-wins/Radical Victory.png"
	} else if zhMeta.Clothing == "Police Officer" {
		return "images/zfc/messages/hunter-wins/Law Enforcer!.png"
	} else if zhMeta.RightWeapon == "Chainsaw" {
		return "images/zfc/messages/hunter-wins/Groovy.png"
	} else if zhMeta.LeftWeapon == "Lucy Bat" {
		return "images/zfc/messages/hunter-wins/Grand Slam!.png"
	} else if zhMeta.RightWeapon == "Flamethrower" {
		return "images/zfc/messages/hunter-wins/Crispy!.png"
	} else if zhMeta.RightWeapon == "Foam Finger" {
		return "images/zfc/messages/hunter-wins/You're Number One!.png"
	} else if zhMeta.LeftWeapon == "Butcher Knife" {
		return "images/zfc/messages/hunter-wins/Minced Meated.png"
	} else if zhMeta.Clothing == "Pirate" {
		return "images/zfc/messages/hunter-wins/Ahoy!.png"
	} else if zhMeta.RightWeapon == "Eternity Glove" {
		return "images/zfc/messages/hunter-wins/Oh Snap!.png"
	} else if zhMeta.RightWeapon == "Patriot Shield" {
		return "images/zfc/messages/hunter-wins/Avenged!.png"
	} else if zhMeta.Hat == "Headphones" {
		return "images/zfc/messages/hunter-wins/Head Banger.png"
	} else if zhMeta.Clothing == "Hungry Arrow Girl" {
		return "images/zfc/messages/hunter-wins/The Odds Were in  your favor!.png"
	} else if strings.Contains(zhMeta.LeftWeapon, "Space Wizard") {
		return "images/zfc/messages/hunter-wins/Space Wizard!.png"
	} else if strings.Contains(zhMeta.Skin, "Zombie") {
		return "images/zfc/messages/hunter-wins/Moral Victory!.png"
	} else if zhMeta.RightWeapon == "Pistol" || zhMeta.RightWeapon == "Revolver" {
		return "images/zfc/messages/hunter-wins/Packing Heat.png"
	} else if zhMeta.RightWeapon == "Fire Hand" || zhMeta.LeftWeapon == "Fire Hand" {
		return "images/zfc/messages/hunter-wins/_Scorched!.png"
	} else if zhMeta.Hat == "Chef" {
		return "images/zfc/messages/hunter-wins/_Delicious-Victory.png"
	}

	return ""
}

func determineIfZombieWinMessage(zcMeta ZombieChain) string {
	if zcMeta.Hat == "Pilot" {
		return "images/zfc/messages/zombie-wins/_Aces-High!.png"
	} else if zcMeta.Clothing == "Clown" && zcMeta.Eyes == "Clown" && zcMeta.Mouth == "Clown" {
		return "images/zfc/messages/zombie-wins/_Clown-Show.png"
	} else if zcMeta.Clothing == "UFO" {
		return "images/zfc/messages/zombie-wins/_Alien-Lover.png"
	} else if zcMeta.Eyes == "Disguise" {
		return "images/zfc/messages/zombie-wins/_Anonymous-Victory!.png"
	} else if zcMeta.Clothing == "Pirate" {
		return "images/zfc/messages/zombie-wins/_Arrrghhhh!!!.png"
	} else if zcMeta.Eyes == "Visor Down" {
		return "images/zfc/messages/zombie-wins/_Moon-Zombies-Rock.png"
	} else if strings.Contains(zcMeta.Weapon, "Baby") {
		return "images/zfc/messages/zombie-wins/_Don't-Baby-Me!.png"
	} else if strings.Contains(zcMeta.Weapon, "Rainbow") {
		return "images/zfc/messages/zombie-wins/_Double-Rainbow--Wings---what-does--it-mean-.png"
	} else if zcMeta.Hat == "Crown" {
		return "images/zfc/messages/zombie-wins/_Royal-Victory!.png"
	} else if zcMeta.Eyes == "Machine" {
		return "images/zfc/messages/zombie-wins/_Judgement-Day!.png"
	} else if zcMeta.Clothing == "Maximalist" {
		return "images/zfc/messages/zombie-wins/_Zombie-Maximalists--FTW!.png"
	} else if zcMeta.Hat == "Cat" {
		return "images/zfc/messages/zombie-wins/_Have-no-fear!.png"
	} else if zcMeta.Nose == "Bull Ring" {
		return "images/zfc/messages/zombie-wins/_Zombie-Bull.png"
	} else if zcMeta.Eyes == "Aviator" {
		return "images/zfc/messages/zombie-wins/_Top-Gun.png"
	} else if zcMeta.Hat == "Captain" {
		return "images/zfc/messages/zombie-wins/_I'm-the-Captain--Now!.png"
	} else if strings.Contains(zcMeta.Mouth, "Cigar") {
		return "images/zfc/messages/zombie-wins/_Smokin'!.png"
	} else if zcMeta.Hat == "Beret" {
		return "images/zfc/messages/zombie-wins/_victorieux!.png"
	} else if zcMeta.Hat == "Karate" {
		return "images/zfc/messages/zombie-wins/_Wax-ON!.png"
	} else if zcMeta.Mouth == "Vampire" {
		return "images/zfc/messages/zombie-wins/_Drained!.png"
	} else if zcMeta.Mouth == "Stitches" {
		return "images/zfc/messages/zombie-wins/_Silent-But-Deadly!.png"
	} else if zcMeta.Clothing == "Pimp" {
		return "images/zfc/messages/zombie-wins/_OG-Pimp!.png"
	} else if zcMeta.Hat == "Wizard" {
		return "images/zfc/messages/zombie-wins/_Magical-Victory.png"
	} else if zcMeta.Hat == "Pumpkin" {
		return "images/zfc/messages/zombie-wins/_Trick-or-Treat!.png"
	} else if zcMeta.Mouth == "Predator" {
		return "images/zfc/messages/zombie-wins/_Want-Some-Candy-.png"
	} else if zcMeta.Skin == "Giraffe" {
		return "images/zfc/messages/zombie-wins/_What-does-the--Giraffe-Say-.png"
	} else if zcMeta.Skin == "Decay" {
		return "images/zfc/messages/zombie-wins/_Infected!.png"
	} else if zcMeta.Hat == "Headphones" {
		return "images/zfc/messages/zombie-wins/_Head-Banger.png"
	} else if zcMeta.Clothing == "Clown" {
		return "images/zfc/messages/zombie-wins/_Zombie-don't--play-that!.png"
	} else if zcMeta.Weapon == "Arrows" {
		return "images/zfc/messages/zombie-wins/_Bullseye.png"
	} else if zcMeta.Hat == "Bowler" {
		return "images/zfc/messages/zombie-wins/Gentlemen's Duel.png"
	} else if zcMeta.Hat == "Plunger" {
		return "images/zfc/messages/zombie-wins/Flushed.png"
	} else if zcMeta.Chain != "None" { //this rule last
		return "images/zfc/messages/zombie-wins/_Zombie-Chains-FTW!.png"
	}

	return ""
}

func (s Server) BuildAlienImage(alien Alien) (*image.RGBA, error) {
	logrus.Infof("Building random alien %w", alien)

	baseImage := image.NewRGBA(image.Rect(0, 0, 1080, 1080))

	background, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/alien/background/%s.jpg", formatName(alien.Background)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, background)

	skin, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/alien/skin/%s.png", formatName(alien.Skin)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, skin)

	clothes, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/alien/clothes/%s.png", formatName(alien.Clothes)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, clothes)

	hat, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/alien/hat/%s.png", formatName(alien.Hat)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, hat)

	hand, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/alien/hand/%s.png", formatName(alien.Hand)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, hand)

	mouth, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/alien/mouth/%s.png", formatName(alien.Mouth)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, mouth)

	eyes, err := util.GetImageFromFileNoCache(fmt.Sprintf("images/alien/eyes/%s.png", formatName(alien.Eyes)))
	if err != nil {
		return nil, err
	}
	util.DrawOver(baseImage, eyes)

	//check if need to resize
	if alien.Width != 0 && alien.Height != 0 {
		resizedImage := ResizeImage(*baseImage, uint(alien.Width), uint(alien.Height))
		b := resizedImage.Bounds()
		m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(m, m.Bounds(), resizedImage, b.Min, draw.Src)
		return m, nil
	}

	return baseImage, nil
}

func ResizeImage(imageToResize image.RGBA, width, height uint) image.Image {
	return resize.Resize(width, height, &imageToResize, resize.Bicubic)
}

func formatName(name string) string {
	return strings.ReplaceAll(name, " ", "-")
}

func addLabel(img *image.RGBA, x, y int, label string) error {
	// col := color.RGBA{200, 100, 0, 255}
	// col = color.Black
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	fontBytes, err := ioutil.ReadFile("images/zfc/font/RoadRage.otf")
	if err != nil {
		return err
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: face,
		Dot:  point,
	}
	d.DrawString(label)

	return nil
}

// func addLabel(img *image.RGBA, text string) (*image.RGBA, error) {
// 	//initialize the context
// 	c := freetype.NewContext()

// 	//read font data
// 	fontBytes, err := ioutil.ReadFile("images/zfc/font/font.ttf")
// 	if err != nil {
// 		return nil, err
// 	}
// 	f, err := freetype.ParseFont(fontBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//set label configuration
// 	c.SetDPI(64)
// 	c.SetFont(f)
// 	c.SetFontSize(48)
// 	c.SetClip(img.Bounds())
// 	c.SetDst(img)
// 	c.SetSrc(image.Black)

// 	//positioning the label
// 	pt := freetype.Pt(200, 200)

// 	//draw the label on image
// 	_, err = c.DrawString(text, pt)
// 	if err != nil {
// 		log.Println(err)
// 		return img, nil
// 	}
// 	pt.Y += c.PointToFixed(24)

// 	return img, nil
// }
