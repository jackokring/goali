# ~/.bashrc: executed by bash(1) for non-login shells.
# see /usr/share/doc/bash/examples/startup-files (in the package bash-doc)
# for examples

# If not running interactively, don't do anything
case $- in
*i*) ;;
*) return ;;
esac

# don't put duplicate lines or lines starting with space in the history.
# See bash(1) for more options
HISTCONTROL=ignoreboth
HISTTIMEFORMAT="%Y-%m-%d %T "

# append to the history file, don't overwrite it
shopt -s histappend

# for setting history length see HISTSIZE and HISTFILESIZE in bash(1)
HISTSIZE=1000
HISTFILESIZE=2000

# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

# If set, the pattern "**" used in a pathname expansion context will
# match all files and zero or more directories and subdirectories.
#shopt -s globstar

# make less more friendly for non-text input files, see lesspipe(1)
#[ -x /usr/bin/lesspipe ] && eval "$(SHELL=/bin/sh lesspipe)"

# set variable identifying the chroot you work in (used in the prompt below)
if [ -z "${debian_chroot:-}" ] && [ -r /etc/debian_chroot ]; then
	debian_chroot=$(cat /etc/debian_chroot)
fi

# set better for nerdfont and nvim
export TERM=st-256color

# set a fancy prompt (non-color, unless we know we "want" color)
case "$TERM" in
xterm-color | *-256color) color_prompt=yes ;;
esac

# uncomment for a colored prompt, if the terminal has the capability; turned
# off by default to not distract the user: the focus in a terminal window
# should be on the output of commands, not on the prompt
#force_color_prompt=yes

if [ -n "$force_color_prompt" ]; then
	if [ -x /usr/bin/tput ] && tput setaf 1 >&/dev/null; then
		# We have color support; assume it's compliant with Ecma-48
		# (ISO/IEC-6429). (Lack of such support is extremely rare, and such
		# a case would tend to support setf rather than setaf.)
		color_prompt=yes
	else
		color_prompt=
	fi
fi

if [ "$color_prompt" = yes ]; then
	PS1='${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '
else
	PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w\$ '
fi
unset color_prompt force_color_prompt

# If this is an xterm set the title to user@host:dir
case "$TERM" in
xterm* | rxvt*)
	PS1="\[\e]0;${debian_chroot:+($debian_chroot)}\u@\h: \w\a\]$PS1"
	;;
*) ;;
esac

# enable color support of ls and also add handy aliases
if [ -x /usr/bin/dircolors ]; then
	test -r ~/.dircolors && eval "$(dircolors -b ~/.dircolors)" || eval "$(dircolors -b)"
fi

# colored GCC warnings and errors
export GCC_COLORS='error=01;31:warning=01;35:note=01;36:caret=01;32:locus=01:quote=01'

# set PATH so it includes cargo bin if it exists
if [ -d "$HOME/.cargo/bin" ]; then
	PATH="$HOME/.cargo/bin:$PATH"
else
	# chromebook to small for rust
	alias eza='exa'
fi

# a spelling to just true/false a command exists
witch() {
	which "$@" >/dev/null 2>/dev/null
}

# some more ls aliases
alias ls='eza'
alias ll='eza -alh --git'
alias top='htop'
alias la='eza -a'
alias tree='eza -al --tree --level=3'
if witch bat; then
	alias cat='bat'
else
	alias cat='batcat'
fi
alias h='history'

# easy cd
alias ..="cd .."
alias ...="cd ../../"
alias ....="cd ../../../"

if ! witch sudo; then
	alias sudo=''
fi
# further alias
alias apt='sudo nala'
alias update="apt upgrade -y; apt clean" # apt update; not req
alias venv='python -m venv'
alias mv='mv -i'
if witch trash; then
	alias rm='trash -v'
fi
alias mkdir='mkdir -p'
alias ps='ps auxf'
alias ping='ping -c 10'
alias less='less -R'
alias cls='clear'
alias tmux='tmux attach || tmux'
alias dmenu='rofi -dmenu -normal-window'
# what a package botch
alias fd='fdfind'

v() {
	# nvim via the st terminal (nerd font)
	# termux don't use st but ~/.termux/font.ttf
	st nvim "$@" 2>/dev/null &
}

ard() {
	# use pipwire-jack alsa midi
	pw-jack ardour9 2>/dev/null &
}

carla() {
	# use pipewire-jack alsa midi
	pw-jack carla 2>/dev/null &
}

graph() {
	# to launch background and tray it
	qpwgraph 2>/dev/null &
}

crash() {
	# show module loading in nvim or other
	# to aid crash black screen of death
	# loading library?
	lsof -p $(pgrep "$@")
}

p() {
	# pet search -- command snippets
	pet search "$@"
}

pn() {
	# pet new -- command snippet
	pet new "$@"
}

#alias pgadmin='pgadmin4&'
alias tor='sudo systemctl restart tor'
alias n='nano'
alias did='history|grep'
alias ok='test $? == 0'
alias freeze='tmuxp freeze'

# enter arch container and list containers (rooted because)
alias arch='distrobox enter --root arch'
alias dbl='distrobox --root list'

# docker via podman
lazypod() {
	systemctl --user enable --now podman.socket
	lazydocker
}
alias docker=podman
export DOCKER_HOST=unix:///run/user/1000/podman/podman.sock

# useful functions
s() { # do sudo, or sudo the last command if no argument given
	if [[ $# == 0 ]]; then
		sudo "$(history -p '!!')"
	else
		sudo "$@"
	fi
}

extract() {
	# might add *.blwz from PiPy phinka module some time
	if [ $# -eq 0 ]; then
		# display usage if no parameters given
		echo "Usage: extract <path/file_name>.<zip|rar|bz2|gz|tar|tbz2|tgz|Z|7z|xz|ex|tar.bz2|tar.gz|tar.xz|.zlib|.cso|.zst|.lha>"
		echo "       extract <path/file_name_1.ext> [path/file_name_2.ext] [path/file_name_3.ext]"
	fi
	for n in "$@"; do
		if [ ! -f "$n" ]; then
			echo "'$n' - file doesn't exist"
			return 1
		fi

		case "${n%,}" in
		*.cbt | *.tar.bz2 | *.tar.gz | *.tar.xz | *.tbz2 | *.tgz | *.txz | *.tar)
			tar zxvf "$n"
			;;
		*.lzma) unlzma ./"$n" ;;
		*.bz2) bunzip2 ./"$n" ;;
		*.cbr | *.rar) unrar x -ad ./"$n" ;;
		*.gz) gunzip ./"$n" ;;
		*.cbz | *.epub | *.zip) unzip ./"$n" ;;
		*.z) uncompress ./"$n" ;;
		*.7z | *.apk | *.arj | *.cab | *.cb7 | *.chm | *.deb | *.iso | *.lzh | *.msi | *.pkg | *.rpm | *.udf | *.wim | *.xar | *.vhd)
			7z x ./"$n"
			;;
		*.xz) unxz ./"$n" ;;
		*.exe) cabextract ./"$n" ;;
		*.cpio) cpio -id <./"$n" ;;
		*.cba | *.ace) unace x ./"$n" ;;
		*.zpaq) zpaq x ./"$n" ;;
		*.arc) arc e ./"$n" ;;
		*.cso) ciso 0 ./"$n" ./"$n.iso" &&
			extract "$n.iso" && \rm -f "$n" ;;
		*.zlib) zlib-flate -uncompress <./"$n" >./"$n.tmp" &&
			mv ./"$n.tmp" ./"${n%.*zlib}" && rm -f "$n" ;;
		*.dmg)
			hdiutil mount ./"$n" -mountpoint "./$n.mounted"
			;;
		*.tar.zst) tar -I zstd -xvf ./"$n" ;;
		*.zst) zstd -d ./"$n" ;;
		*.lha) lha x ./"$n" ;;
		*)
			echo "extract: '$n' - unknown archive method"
			return 1
			;;
		esac
	done
}

# a whiptail alias set
export WHIP="whiptail"
# whiptail tail
export TAIL="0 0 3>&2 2>&1 1>&3"

# launch a process singleton
single() {
	# self and process?
	if test "$(ps | grep "$1" | wc -l)" != 2; then
		"$@" >/dev/null 2>&1 &
	fi
}

# Automatically activate Python venv if it exists
auto_activate_venv() {
	if [ -e "./bin/activate" ]; then
		source ./bin/activate
	fi
}

# Override the 'cd' command to call our function
cd() {
	builtin cd "$@" && auto_activate_venv
}

# If you use pushd/popd, you can override them too.
pushd() {
	builtin pushd "$@" && auto_activate_venv
}

popd() {
	builtin popd && auto_activate_venv
}

# quick almost shortcut
# root rofi
//() {
	rofi -show combi -normal-window &
}

# tab delimited autojump database directories
/() {
	DIR="$(awk -F '\t' '{print $2}' ~/.local/share/autojump/autojump.txt | rofi -dmenu -normal-window)"
	# if tmux session then jump not cd
	if [ -e "$DIR/.tmuxp.yaml" ]; then
		tmuxp load "$DIR"
	else
		cd "$DIR" || return
	fi
}

# cd -
-() {
	cd - || return
}

# fast GPT 3.5
#export OPENAI_API_KEY=

# enable programmable completion features (you don't need to enable
# this, if it's already enabled in /etc/bash.bashrc and /etc/profile
# sources /etc/bash.bashrc).
if ! shopt -oq posix; then
	if [ -f /usr/share/bash-completion/bash_completion ]; then
		. /usr/share/bash-completion/bash_completion
	elif [ -f /etc/bash_completion ]; then
		. /etc/bash_completion
	fi
fi

# Ctrl+S forward search to match Ctrl-R reverse search
# so turn off xon/xoff flow control
stty -ixon

export TZ="UTC"
# add commit push
gacp() {
	date=$(date +"%A %Y-%m-%d %H:%M:%S")
	message="${1:-$date}"
	git add .
	git commit -m "$message"
	git push
}

gacr() {
	date=$(date +"%A %Y-%m-%d %H:%M:%S")
  # avoid stash
	git add .
	git commit -m "[rebase] $date"
  # apply anything remote
	git pull --rebase
  BRANCH="$(git branch|sed -nr "s/^\* (.*)\\\$/\\1/p")"
  # get the import you want a rebase off
  git checkout "$1"
  # as your hopefully not editing it
  git pull
  # now got updates so back to edit branch
  git checkout "$BRANCH"
  # apply a rebase merge from the requested branch
  git rebase "$1"
}

# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/bin" ]; then
	PATH="$HOME/bin:$PATH"
fi

# set PATH so it includes pipx's bin if it exists
if [ -d "$HOME/.local/bin" ]; then
	PATH="$HOME/.local/bin:$PATH"
fi

#ensure escape processing
alias echo='echo -e'

# N.B. DETECT ARCH
if ls /etc/arch-release 2>/dev/null; then
	# in arch so do the .archrc and exit
	. "$HOME/.archrc"
#	exit 0
else
	# MINT OR TERMUX
	export ARCH=$(gcc -dumpmachine)
	export LV2_PATH=/usr/local/lib/$ARCH/lv2/

	# Arm kit
	#PATH="/usr/local/gcc-arm-none-eabi-8-2018-q4-major/bin:$PATH"

	# z88dk Z80 dev kit
	if [ -d ${HOME}/z88dk ]; then
		export PATH=${PATH}:${HOME}/z88dk/bin
		export ZCCCFG=${HOME}/z88dk/lib/config
		eval "$(perl -I ~/perl5/lib/perl5/ -Mlocal::lib)"
	fi

	# color vars
	export NONE='\e[0m'
	export RED='\e[1;31m'
	export GREEN='\e[1;32m'
	export YELLOW='\e[1;33m'
	export BLUE='\e[1;34m'
	export MAGENTA='\e[1;35m'
	export CYAN='\e[1;36m'
	export WHITE='\e[1;37m'

	# notes (denoise echo vs. remove action => grouping)
	echo "# command and location history search\n\
$CYAN^R$NONE is reverse command search. $CYAN^S$NONE is forward command\
 search (No XON/XOFF). Directory autojump ${GREEN}j$NONE (and ${GREEN}jc$NONE)\
 are installed. First parameter for match. Also ${GREEN}s$NONE last command sudo.\
 ${GREEN}h$NONE is for command history, also ${GREEN}did$NONE $RED!$NONE\n"
	echo "# useful knowledge and additions\n\
$CYAN^D$NONE is end of stream terminate process. $CYAN^Z$NONE is process \
stop and ${GREEN}fg$NONE (and ${GREEN}bg$NONE) job control numbers.\
 ${GREEN}ll$NONE and ${GREEN}la$NONE do modified ${GREEN}ls$NONE types.\
 ${GREEN}espeak-ng$NONE for robot voice.\
 ${GREEN}entr$NONE file watcher command execute.\
 ${GREEN}extract$NONE archive type detection and extract.\n"
	echo "# code and data management\n\
${GREEN}gacp$NONE for git add/commit/push with optional message.\
 ${GREEN}fzf$NONE for fuzzy find.\
 ${GREEN}rg$NONE for ripgrep file word finder. ${GREEN}update$NONE does all the software updating in one\
 command. ${GREEN}ncdu$NONE is a disk usage analyzer.\n"
	echo "# $RED~/bin$NONE general user binaries. Go binaries."
	ls ~/bin
	echo
	echo "# $RED~/.local/bin$NONE for ${GREEN}pipx$NONE. You may need to allow\
 packages to use the global python context installed via ${GREEN}apt$NONE.\
 $RED~/.local/pipx/venvs/*/pyvenv.cfg$NONE"
	ls ~/.local/bin
	echo
	echo "Maybe:"
	# vscode seems to have tmux restart issue
	echo "# can use ${GREEN}tmux ${CYAN}^B s <left/right/up/down>, c <new win>, & <kill win>, number <select win>, <space> <menu>$NONE"
	echo "# ${GREEN}pgadmin4$NONE in venv on http://127.0.0.1:5050"
	if witch pgadmin4; then
		single pgadmin4
	fi
	echo "# ${GREEN}tor$NONE on? socks4://127.0.0.1:9050"
	echo "# ${GREEN}fluid$NONE FLTK GUI designer (C++ template tool)"
	echo "# ${GREEN}glade$NONE Gtk GUI designer (XML template tool)"
	echo "# ${GREEN}//$NONE process launcher (rofi tool)"
	echo "# ${GREEN}/$NONE cd to commonly used (rofi tool)"
	echo "# ${GREEN}v$NONE neovim in st session"
	echo "# ${GREEN}p$NONE, ${GREEN}pn$NONE pet search and pet new (command snippets)"
	echo "# ${GREEN}freeze$NONE freeze tmux seesion for /"
	echo "# ${GREEN}carla$NONE, ${GREEN}ardour/ard(pw)$NONE and ${GREEN}graph$NONE for audio makers"
	echo "# ${GREEN}tldr$NONE for command help"
	echo "# ${GREEN}fuck$NONE command corrector"
	echo
	if [ -d "$HOME/.cargo/bin" ]; then
		echo "# $RED~/.cargo/bin$NONE for rust binaries."
		ls ~/.cargo/bin
		echo
	fi

	export GOBIN="$HOME/bin"
	# Install Ruby Gems to ~/gems (for jekyll.sh github.com docs)
	export GEM_HOME="$HOME/gems"
	export PATH="$HOME/gems/bin:$PATH"

	# starship (arm build on google drive)
	eval "$(starship init bash)"

	# activate virtual env after all path stuff
	# autojump
	PREFIX=${PREFIX:-/usr}
	. "$PREFIX/share/autojump/autojump.bash"
	# it's the j and jc aliases todoo
	# autojump >/dev/null
	# venv do
	cd .
	# last, may include venv $PATH mash of added afterj

	# Set up fzf key bindings and fuzzy completion
	eval "$(fzf --bash)"
	#get X display
	if [ "$PREFIX" == "/usr" ]; then
		export DISPLAY=:0
	else
		#vnc default display
		export DISPLAY=:1
	fi
	espeak-ng "What are you doing Dave? They're all dead Dave." &
fi
