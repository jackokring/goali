#!/usr/bin/bash
pip freeze | tee requirements.txt
# copy some user file backups
pushd extras-backup
# $XDG
cp -r ~/.config/nano .
cp -r ~/.config/rofi .
cp -r ~/.config/neofetch .
# irregular
cp ~/.bashrc .
cp ~/.config/starship.toml .
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
gacp $1

