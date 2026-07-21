On Linux (or WSL):

==Note: on the second line here, if you're on COSMIC and want the frosted glass effect to work for emacs, replace `emacs` with `emacs-pgtk`==

```
sudo apt update && sudo apt upgrade -y
sudo apt install -y git emacs ripgrep fd-find cmake libtool-bin libvterm-dev
git clone --depth 1 https://github.com/doomemacs/core.git ~/.config/emacs
~/.config/emacs/bin/doom install
```

enter `y` for yes.

```
echo 'export PATH="$HOME/.config/emacs/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Now cd into `~/.config`

```
rm -rf doom
git clone https://github.com/RockLikeAmadeus/doom-emacs-config.git
mv doom-emacs-config doom
doom sync
```

Finally:

```
emacs &
```

# Set up emacs to run as server and client (linux)

Add the emacs daemon to the startup hook (the process differs by desktop environment, but ultimately you just want to run `emacs --daemon`)

Whenever running emacs, be sure and run emacsclient (or `emacsclient -c -a 'emacs'`) when the daemon is running.
