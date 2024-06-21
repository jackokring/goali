#!/usr/bin/bash
# install some user file backups
pushd extras-backup
../yes-no.sh "install bash shell user's config" && cp .bashrc ~
../yes-no.sh "install starship config" && cp starship.toml ~/.config
if ../yes-no.sh "install tmux config" 
then
cp -r .tmux ~
git submodule update --init --recursive
cp .tmux.conf ~
fi
if ../yes-no.sh "install emacs config" 
then
cp *.el ~/.emacs.d
mkdir -p ~/.emacs.d/packages
pushd packages
cp *.el ~/.emacs.d/packages
popd
cp -r elpa ~/.emacs.d/
fi
popd


