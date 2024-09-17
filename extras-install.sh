#!/usr/bin/bash
. xdg.sh
# yes it does better here
git submodule update --init --recursive
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
install "nvim"
if ../yes-no.sh "install dwm window manager"
then
    cp ~/bin/dwm ~/.mess
    cp ~/bin/dmenu ~/.mess
    cp ~/bin/stest ~/.mess
    cp ~/bin/st ~/.mess
    cp ~/bin/slstatus ~/.mess
    cp ~/startwm.sh ~/.mess
    cp startwm.sh ~
    pushd dwm
	rm config.h
	# now it won't keep old without patches
	make install
	# best launch method
	echo "Delete ~/startwm.sh to go back to old window manager."
	echo "Edit it to add in more dwm auto-starts."
	echo "This does not affect the local login, just XRDP sessions."
	# dmenu and st terminal also to ~/bin
	cd ../dmenu
	make install
	cd ../st
	make install
	cd ../slstatus
	make install
    popd
    # install nerd font used
    echo "Installing nerd font."
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
    rm -rf ~/.mess/.tmux
    cp -r ~/.tmux ~/.mess
    rm -rf ~/.tmux
    cp -r .tmux ~
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


