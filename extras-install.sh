#!/usr/bin/bash
. xdg.sh
install() {
if ../yes-no.sh "install ${1} config"
then
    mkdir -p $XDG_CONFIG_HOME/${1}
    cp -r $XDG_CONFIG_HOME/${1} ~/.mess
    cp -r ${1} $XDG_CONFIG_HOME
fi
}
# install some user file backups
pushd extras-backup
#$XDG
mkdir -p ~/.mess
install "nano"
install "rofi"
install "neofetch"
if ../yes-no.sh "install dwm window manager"
then
    cp ~/bin/dwm ~/.mess
    cp ~/bin/dmenu ~/.mess
    cp ~/bin/stest ~/.mess
    cp ~/bin/st ~/.mess
    cp ~/bin/sl ~/.mess
    cp ~/startwm.sh ~/.mess
    cp startwm.sh ~
    pushd dwm
	make install
	# best launch method
	echo "Delete ~/.xsession to go back to old window manager."
	echo "Edit it to add in more dwm features or auto-starts."
	# dmenu and st terminal also to ~/bin
	cd ../dmenu
	make install
	cd ../st
	make install
	cd ../slstatus
	make install
    popd
    # install nerd font used
    cp -r JetBrainsMono ~/.local/share/fonts
    fc-cache -rv >/dev/null
fi
if ../yes-no.sh "install bash config"
then
    cp ~/.bashrc ~/.mess
    cp .bashrc ~
fi
if ../yes-no.sh "install starship config"
then
    cp $XDG_CONFIG_HOME/starship.toml ~/.mess
    cp starship.toml $XDG_CONFIG_HOME
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


