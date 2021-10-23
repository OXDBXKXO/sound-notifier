package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/itchyny/volume-go"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	soundnotifier = kingpin.New("sound-notifier", "sound controller for Linux that triggers notifications")
	silent = soundnotifier.Flag("silent", "Hide notification when volume changes").Bool()

	set = soundnotifier.Command("set", "Manually set volume (between 0 and 100)")
	setvolume = set.Arg("volume", "Volume delta to apply").Required().Int()

	up = soundnotifier.Command("up", "Increase volume")
	upvolume = up.Arg("volume", "Volume delta to apply").Required().Int()

	down = soundnotifier.Command("down", "Decrease volume")
	downvolume = down.Arg("volume", "Volume delta to apply").Required().Int()

	mute = soundnotifier.Command("mute", "Mute switch (automatic On/Off)")
)

func getVolume() (vol int) {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}

	return vol
}

func notifySend(title string, message string) error{
	var args []string
	args = append(args, title)
	args = append(args, message)
	args = append(args, "-t", "400")

	cmd := exec.Command("notify-send", args...)
	if cmd == nil {
		return fmt.Errorf("Malformed command!")
	}
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error running command: %s!", err)
	}
	return nil
}

func main() {
	switch kingpin.MustParse(soundnotifier.Parse(os.Args[1:])) {

	case set.FullCommand():
		if *setvolume < 0 {
			*setvolume = 0
		}
		if *setvolume > 100 {
			*setvolume = 100
		}

		err := volume.SetVolume(*setvolume)
		if err != nil {
			log.Fatalf("set volume failed: %+v", err)
		}

		if !*silent {
			notifySend("Volume Set", fmt.Sprintf("Volume changed to: %d", *setvolume))
		}

	case up.FullCommand():
		vol := getVolume()
		newVolume := vol + *upvolume
		if newVolume > 100 {
			newVolume = 100
		}

		err := volume.SetVolume(newVolume)
		if err != nil {
			log.Fatalf("set volume failed: %+v", err)
		}

		if !*silent {
			notifySend("Volume Increase", fmt.Sprintf("Volume changed to: %d", newVolume))
		}

	case down.FullCommand():
		vol := getVolume()
		newVolume := vol - *downvolume
		if newVolume < 0 {
			newVolume = 0
		}

		err := volume.SetVolume(newVolume)
		if err != nil {
			log.Fatalf("set volume failed: %+v", err)
		}

		if !*silent {
			notifySend("Volume Decrease", fmt.Sprintf("Volume changed to: %d", newVolume))
		}

	case mute.FullCommand():

		muted, err := volume.GetMuted()
		if err != nil {
			log.Fatalf("get mute state failed: %+v", err)
		}

		if muted {
			err = volume.Unmute()
			if err != nil {
				log.Fatalf("unmute failed: %+v", err)
			}

			if !*silent {
				notifySend("Unmute", "Sound is ON")
			}
		} else {
			err = volume.Mute()
			if err != nil {
				log.Fatalf("mute failed: %+v", err)
			}

			if !*silent {
				notifySend("Mute", "Sound is OFF")
			}
		}
	}
}
