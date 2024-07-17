#!/bin/bash
# edited for bash, as sh link to dash maybe looses some env

# xrdp X session start script (c) 2015, 2017, 2021 mirabilos
# published under The MirOS Licence

# Rely on /etc/pam.d/xrdp-sesman using pam_env to load both
# /etc/environment and /etc/default/locale to initialise the
# locale and the user environment properly.

# is there any system profile needed?
if test -r /etc/profile; then
	. /etc/profile
fi

if test -r ~/.profile; then
	. ~/.profile
fi

# maybe $BASH_VERSION is not set right for .profile
# N.B.  >>>> sh -c "echo \$BASH_VERSION" <<<<
#export PATH="$HOME/bin:$PATH"

# run ~/bin/dwm instead of /usr/bin/dwm

# dbus messages
if which dunst ; then
	dunst&
fi

# desktop background
if which feh ; then
	if test -r ~/.fehbg; then
		~/.fehbg&
	else
		feh --bg-scale --randomize --recursive ~/Pictures&
	fi
fi

# status tray
slstatus&

# launch WM
exec dwm
