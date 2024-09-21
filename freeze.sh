#!/usr/bin/bash
. xdg.sh
pip freeze | tee requirements.txt
# remove *~ files
rm -rf *~
rm -rf \#*\#
# copy some user file backups
pushd extras-backup
# $XDG
cp -r $XDG_CONFIG_HOME/nano .
cp -r $XDG_CONFIG_HOME/rofi .
cp -r $XDG_CONFIG_HOME/neofetch .
cp -r $XDG_CONFIG_HOME/nvim .
# irregular
cp -r ~/.local/share/fonts/JetBrainsMono .
cp ~/.bashrc .
cp $XDG_CONFIG_HOME/starship.toml .
cp ~/.emacs.d/*.el .
cp -r ~/.tmux .
cp ~/.tmux.conf .
mkdir -p packages
pushd packages
cp ~/.emacs.d/packages/*.el .
popd
mkdir -p elpa
cp -r ~/.emacs.d/elpa .
popd
# add commit push
gacp () {
	date=$(date +"%A %Y-%m-%d %H:%M:%S")
  message="${1:-$date}"
  git add . ; git commit -m "$message" ; git push
}
gacp "$1"

