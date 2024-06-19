#!/usr/bin/bash
# install some user file backups
pushd extras-backup
cp .bashrc ~
cp starship.toml ~/.config
cp *.el ~/.emacs.d
mkdir -p ~/.emacs.d/packages
pushd packages
cp *.el ~/.emacs.d/packages
popd
cp -r elpa ~/.emacs.d/
popd


