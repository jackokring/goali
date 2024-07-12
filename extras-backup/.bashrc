# ~/.bashrc: executed by bash(1) for non-login shells.
# see /usr/share/doc/bash/examples/startup-files (in the package bash-doc)
# for examples

# If not running interactively, don't do anything
case $- in
    *i*) ;;
      *) return;;
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

# set a fancy prompt (non-color, unless we know we "want" color)
case "$TERM" in
    xterm-color|*-256color) color_prompt=yes;;
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
xterm*|rxvt*)
    PS1="\[\e]0;${debian_chroot:+($debian_chroot)}\u@\h: \w\a\]$PS1"
    ;;
*)
    ;;
esac

# enable color support of ls and also add handy aliases
if [ -x /usr/bin/dircolors ]; then
    test -r ~/.dircolors && eval "$(dircolors -b ~/.dircolors)" || eval "$(dircolors -b)"
    #alias ls='ls --color=auto'
    #alias dir='dir --color=auto'
    #alias vdir='vdir --color=auto'

    #alias grep='grep --color=auto'
    #alias fgrep='fgrep --color=auto'
    #alias egrep='egrep --color=auto'
fi

# colored GCC warnings and errors
export GCC_COLORS='error=01;31:warning=01;35:note=01;36:caret=01;32:locus=01:quote=01'

# set PATH so it includes cargo bin if it exists
if [ -d "$HOME/.cargo/bin" ] ; then
    PATH="$HOME/.cargo/bin:$PATH"
else
    alias eza='exa'
fi

# some more ls aliases
alias ls='eza'
alias ll='eza -alh --git'
#alias la='ls -A'
#alias l='ls -CF'
#alias rm='trash'
alias top='htop'
alias la='eza -a'
alias tree='eza -al --tree --level=3'
alias cat='batcat'
alias gitfree='for b in $(git branch --merged | grep -v \*); do echo $b; git branch -D $b; done; git submodule foreach --recursive "for b in $(git branch --merged | grep -v \*); do echo $b; git branch -D $b; done"'
alias h='history'

# easy cd
alias ..="cd .."
alias ...="cd ../../"
alias ....="cd ../../../"

# further alias
alias apt='sudo nala'
alias update="apt upgrade -y; apt clean" # apt update; not req
alias venv='python -m venv'
alias mv='mv -i'
alias rm='trash -v'
alias mkdir='mkdir -p'
alias ps='ps auxf'
alias ping='ping -c 10'
alias less='less -R'
alias cls='clear'
alias tmux='tmux attach || tmux'
alias dmenu='rofi -dmenu'

e() {
	# emacs with files and no gtk error stream
	emacs "$@" 2>/dev/null&
}

#alias pgadmin='pgadmin4&'
alias tor='sudo systemctl restart tor'
alias n='nano'
#alias e='emacs'

# useful functions
s() { # do sudo, or sudo the last command if no argument given
    if [[ $# == 0 ]]; then
        sudo $(history -p '!!')
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
          *.cbt|*.tar.bz2|*.tar.gz|*.tar.xz|*.tbz2|*.tgz|*.txz|*.tar)
                       tar zxvf "$n"       ;;
          *.lzma)      unlzma ./"$n"      ;;
          *.bz2)       bunzip2 ./"$n"     ;;
          *.cbr|*.rar) unrar x -ad ./"$n" ;;
          *.gz)        gunzip ./"$n"      ;;
          *.cbz|*.epub|*.zip) unzip ./"$n"   ;;
          *.z)         uncompress ./"$n"  ;;
          *.7z|*.apk|*.arj|*.cab|*.cb7|*.chm|*.deb|*.iso|*.lzh|*.msi|*.pkg|*.rpm|*.udf|*.wim|*.xar|*.vhd)
                       7z x ./"$n"        ;;
          *.xz)        unxz ./"$n"        ;;
          *.exe)       cabextract ./"$n"  ;;
          *.cpio)      cpio -id < ./"$n"  ;;
          *.cba|*.ace) unace x ./"$n"     ;;
          *.zpaq)      zpaq x ./"$n"      ;;
          *.arc)       arc e ./"$n"       ;;
          *.cso)       ciso 0 ./"$n" ./"$n.iso" && \
                            extract "$n.iso" && \rm -f "$n" ;;
          *.zlib)      zlib-flate -uncompress < ./"$n" > ./"$n.tmp" && \
                            mv ./"$n.tmp" ./"${n%.*zlib}" && rm -f "$n"   ;;
          *.dmg)
                      hdiutil mount ./"$n" -mountpoint "./$n.mounted" ;;
          *.tar.zst)  tar -I zstd -xvf ./"$n"  ;;
          *.zst)      zstd -d ./"$n"  ;;
          *.lha)      lha x ./"$n"  ;;
          *)
                      echo "extract: '$n' - unknown archive method"
                      return 1
                      ;;
        esac
    done
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
    builtin popd "$@" && auto_activate_venv
}

# fast GPT 3.5
#export OPENAI_API_KEY=

# Alias definitions.
# You may want to put all your additions into a separate file like
# ~/.bash_aliases, instead of adding them here directly.
# See /usr/share/doc/bash-doc/examples in the bash-doc package.

if [ -f ~/.bash_aliases ]; then
    . ~/.bash_aliases
fi

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
alias cp='[ $# == 1 ] && cp -i $1 . || cp $*'

# Rack rebuild needs out of date
alias plugout='find ~/Rack/plugins -name "*.vcvplugin" -delete; '

export TZ="UTC"
# add commit push
gacp () {
	date=$(date +"%A %Y-%m-%d %H:%M:%S")
  message="${1:-$date}"
  git add . ; git commit -m "$message" ; git push
}

# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/bin" ] ; then
    PATH="$HOME/bin:$PATH"
fi

# set PATH so it includes pipx's bin if it exists
if [ -d "$HOME/.local/bin" ] ; then
    PATH="$HOME/.local/bin:$PATH"
fi

# Arm kit
#PATH="/usr/local/gcc-arm-none-eabi-8-2018-q4-major/bin:$PATH"

# color vars
export NONE='\e[0m'
export RED='\e[1;31m'
export GREEN='\e[1;32m'
export YELLOW='\e[1;33m'
export BLUE='\e[1;34m'
export MAGENTA='\e[1;35m'
export CYAN='\e[1;36m'
export WHITE='\e[1;37m'
#ensure escape processing
alias echo='echo -e'

# fzf in tmux
#alias fzf='fzf-tmux' # -p' # in float pane
# ${GREEN}tmux ^B ?$NONE terminal multiplex.

# notes (denoise echo vs. remove action => grouping)
printf "# command and location history search\n\
$CYAN^R$NONE is reverse command search. $CYAN^S$NONE is forward command\
 search (No XON/XOFF). Directory autojump ${GREEN}j$NONE (and ${GREEN}jc$NONE)\
 are installed. First parameter for match. ${GREEN}tldr$NONE for command help.\
 ${GREEN}fuck$NONE command corrector. Also ${GREEN}s$NONE last command sudo.\
 ${GREEN}h$NONE is for command history $RED!$NONE\n\n"
printf "# useful knowledge and additions\n\
$CYAN^D$NONE is end of stream terminate process. $CYAN^Z$NONE is process \
stop and ${GREEN}fg$NONE (and ${GREEN}bg$NONE) job control numbers.\
 ${GREEN}ll$NONE and ${GREEN}la$NONE do modified ${GREEN}ls$NONE types.\
 ${GREEN}espeak-ng$NONE for robot voice.\
 ${GREEN}entr$NONE file watcher command execute.\
 ${GREEN}extract$NONE archive type detection and extract.\
 ${GREEN}dragon$NONE CLI drag and drop manager.\
 ${CYAN}Shft+^V$NONE is paste in some contexts.\n\n"
printf "# modified behaviour\n\
Copy ${GREEN}cp$NONE has a single argument only automatic target of the\
 ${GREEN}pwd$NONE.\n\n"
printf "# code and data management\n\
${GREEN}gacp$NONE for git add/commit/push with optional message.\
 ${GREEN}gitfree$NONE for pruning a git repo. ${GREEN}fzf$NONE for fuzzy find.\
 ${GREEN}rg$NONE for ripgrep file word finder. ${GREEN}trash$NONE for trash\
 can management. ${GREEN}update$NONE does all the software updating in one\
 command. ${GREEN}ncdu$NONE is a disk usage analyzer.\n\n"
printf "# Rack\n\
${GREEN}plugout$NONE to remove all plugin ${GREEN}.vcvplugin$NONE files to\
 allow remaking them.\n\n"
printf "# $RED~/bin$NONE general user binaries.\n"
ls ~/bin
echo
printf "# $RED~/.local/bin$NONE for ${GREEN}pipx$NONE. You may need to allow\
 packages to use the global python context installed via ${GREEN}apt$NONE.\
 $RED~/.local/pipx/venvs/*/pyvenv.cfg$NONE\n"
ls ~/.local/bin
echo
# vscode seems to have tmux restart issue
printf "# can use ${GREEN}tmux ${CYAN}^B s <left/right/up/down>$NONE ...\n"
printf "# ${GREEN}pgadmin4$NONE in venv on http://127.0.0.1:5050\n"
single() {
# self and process?
	if test $(ps|grep $1|wc -l) != 2; then
		"$@" 2>&1 >/dev/null&
	fi
}
single pgadmin4

printf "# ${GREEN}tor$NONE on? socks4://127.0.0.1:9050\n"
printf "# ${GREEN}fluid$NONE FLTK GUI designer (C++ template tool)\n"
printf "# ${GREEN}glade$NONE Gtk GUI designer (XML template tool)\n"
echo
if [ -d "$HOME/.cargo/bin" ]; then
printf "# $RED~/.cargo/bin$NONE for rust binaries.\n"
ls ~/.cargo/bin
echo
fi
# continue by doing the reset of the .profile file
printf "# .profile for perhaps .NET\n"

# Install Ruby Gems to ~/gems (for jekyll.sh github.com docs)
export GEM_HOME="$HOME/gems"
export PATH="$HOME/gems/bin:$PATH"

# go env -w GOBIN=$HOME/bin

# starship (arm build on google drive)
eval "$(starship init bash)"

# activate virtual env after all path stuff
# autojump
. /usr/share/autojump/autojump.sh
# last, may include venv $PATH mash of added afterj

espeak-ng "What are you doing Dave? They're all dead Dave."
