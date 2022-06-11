package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// convertDuration converts a duration in seconds to a readable format (e.g. 1:23)
func convertDuration(s int64) string {
	minutes := s / 60
	seconds := s % 60
	str := fmt.Sprintf("%d:%d", minutes, seconds)
	return str
}

// printHelp prints the help text (I am sure there is a better way to do this)
func printHelp() {
	println(`Usage: spotify [command]
Commands:
	toggle, t:		Toggles the playback state
	play, p:		Plays the current track
	pause, stop:		Pauses the current track
	next, n, skip, s:	Skips to the next track
	previous, back, b:	Skips to the previous track
	current, status:	Shows the current track
	volume, v:		Sets the volume
	quit, q:		Quits Spotify
	shuffle, sh:		Toggles shuffle
	repeat, r:		Toggles repeat
	help, h:		Shows this help
	`)

}

// TODO: add time elapsed and duration data
// status prints the current status of Spotify, including the current track and the current volume
func status() {
	state := execute("-e", "tell application \"Spotify\" to return player state as string")
	if state == "playing" {
		track := execute("-e", "tell application \"Spotify\" to return artist of current track & \" - \" & name of current track")
		volume := execute("-e", "tell application \"Spotify\" to return sound volume as integer")
		repeatAndShuffle := execute("-e", "tell application \"Spotify\" to return repeating as string & \" - \" & shuffling as string")
		println("Playing " + track)
		println("Volume is set to:", volume+"%")
		println("Repeat and Shuffle is set to:", repeatAndShuffle)
	} else {
		println("Spotify is not playing")
	}
}

// execute executes the given AppleScript command and returns the output
func execute(command ...string) string {
	out, err := exec.Command("osascript", command...).Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		exec.Command("osascript", "-e", "tell application \"Spotify\" to launch").Run()
		printHelp()
		return
	}

	switch args[0] {
	case "toggle", "t":
		execute("-e", "tell application \"Spotify\" to playpause")
		status()
		break

	case "play", "p":
		execute("-e", "tell application \"Spotify\" to play")
		status()
		break

	case "pause", "stop":
		execute("-e", "tell application \"Spotify\" to pause")
		status()
		break

	case "next", "n", "skip", "s":
		execute("-e", "tell application \"Spotify\" to next track")
		status()
		break

	case "previous", "back", "b":
		execute("-e", "tell application \"Spotify\" to previous track")
		status()
		break

	case "current", "current-track", "ct", "cs", "status":
		status()
		break

	case "volume", "v":
		if len(args) == 1 {
			out := execute("-e", "tell application \"Spotify\" to return sound volume as integer")
			println("Volume is set to:", out+"%")
		} else {
			execute("-e", "tell application \"Spotify\" to set sound volume to "+args[1])
			println("Volume set to:", args[1]+"%")
		}
		break

	case "quit", "q":
		execute("-e", "tell application \"Spotify\" to quit")
		break

	case "shuffle", "sh":
		if execute("-e", "tell application \"Spotify\" to return shuffling as boolean") == "true" {
			execute("-e", "tell application \"Spotify\" to set shuffling to false")
			println("Shuffle is now off")
		} else {
			execute("-e", "tell application \"Spotify\" to set shuffling to true")
			println("Shuffle is now on")
		}
		break

	case "repeat", "r":
		if execute("-e", "tell application \"Spotify\" to return repeating as boolean") == "true" {
			execute("-e", "tell application \"Spotify\" to set repeating to false")
			println("Repeat is now off")
		} else {
			execute("-e", "tell application \"Spotify\" to set repeating to true")
			println("Repeat is now on")
		}
		break

	default:
		println("Unknown command:\n", args[0])
		printHelp()
		break
	}
}
