On Linux (or WSL):

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