# Spotiterm

Spotiterm let's you control your spotify player from the terminal without leaving your workplace (**currently only on macOS**).

## Installation & Updating

Run the following command in your terminal to install the latest version for your platform:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/CanCodes/spotiterm/main/install.sh)"
```

Whenever there is a new version, you will be notified in your terminal.

## Usage

```
Usage: spotiterm [command]
Commands:
        toggle, t:              Toggles the playback state
        play, p:                Plays the current track
        pause, stop:            Pauses the current track
        next, n, skip, s:       Skips to the next track
        previous, back, b:      Skips to the previous track
        current, status:        Shows the current track
        volume, v:              Sets the volume
        quit, q:                Quits Spotify
        shuffle, sh:            Toggles shuffle
        repeat, r:              Toggles repeat
        help, h:                Shows this help

Examples:
        Play a Song From URL: spotiterm play <url>
```

## Contributing

Feel free to open an issue or pull request!

## License

Spotiterm is licensed under the MIT license.

## Legal

If you have any legal issues, please contact me at [<https://twitter.com/cancodez>](https://twitter.com/cancodez).
