#!/usr/bin/bash
install() {
if ../yes-no.sh "install ${1} config"
then
    mkdir -p ~/.config/${1}
    cp -r ~/.config/${1} ~/.mess
    cp -r ${1} ~/.config
fi
}
# install some user file backups
pushd extras-backup
#$XDG
mkdir -p ~/.mess
install "nano"
install "rofi"
install "neofetch"
if ../yes-no.sh "install bash config"
then
    cp ~/.bashrc ~/.mess
    cp .bashrc ~
fi
if ../yes-no.sh "install starship config"
then
    cp ~/.config/starship.toml ~/.mess
    cp starship.toml ~/.config
fi
if ../yes-no.sh "install tmux config" 
then
    cp -r ~/.tmux ~/.mess
    cp -r .tmux ~
    git submodule update --init --recursive
    cp ~/.tmux.conf ~/.mess
    cp .tmux.conf ~
fi
if ../yes-no.sh "install emacs config" 
then
    cp -r ~/.emacs.d ~/.mess
    cp *.el ~/.emacs.d
    mkdir -p ~/.emacs.d/packages
    pushd packages
    cp *.el ~/.emacs.d/packages
    popd
    cp -r elpa ~/.emacs.d/
fi
popd


