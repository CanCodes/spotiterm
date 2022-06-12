if [ `uname` != "Darwin" ]; then
    echo "Spotiterm is only available on Mac OS X for now."
    exit 1
fi

# Ask for root privileges

arch=$(uname -m)

url="https://github.com/cancodes/spotiterm/releases/latest/download/$arch-mac.zip"

mkdir spotiterm_cache && cd spotiterm_cache

echo "Downloading the latest version of Spotiterm..."
curl -L $url -o spotiterm.zip
# mv ../spotiterm.zip .
unzip spotiterm.zip

read -p "Would you like to rename command spotiterm to spotify? (y/n) " -n 1 -r REPLY && echo
cmd="spotiterm"
# If the user wants to rename the command, do it. Otherwise, just move the app to the /usr/local/bin directory.
if [[ $REPLY =~ ^[Yy]$ ]]; then
    sudo mv spotiterm /usr/local/bin/spotify && chmod +x /usr/local/bin/spotify
    cmd="spotify"
else
    sudo mv spotiterm /usr/local/bin/ && chmod +x /usr/local/bin/spotiterm
fi

echo "\nCleaning up..."
cd .. && rm -rf spotiterm_cache

echo "Spotiterm installed successfully!"
echo "Run '$cmd' to start."
