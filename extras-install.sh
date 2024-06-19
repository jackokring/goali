#!/usr/bin/bash
# install some user file backups
pushd extras-backup
cp .bashrc ~
cp starship.toml ~/.config
cp init.el ~/.emacs.d
popd


