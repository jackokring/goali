#!/usr/bin/bash
# Source this to set XDG variables
# set config base
export XDG_CONFIG_HOME="${XDG_CONFIG_HOME:-$HOME/.config}"
# Import all XDG set by xdg-config-update autorun
# THey're not exporting by default
. $XDG_CONFIG_HOME/user-dirs.dirs