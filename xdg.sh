#!/usr/bin/bash
# Source this to set XDG variables
# set config base plus ...
export XDG_CONFIG_HOME="${XDG_CONFIG_HOME:-$HOME/.config}"
export XDG_DATA_HOME="${XDG_DATA_HOME:-$HOME/.local/share}"
export XDG_STATE_HOME="${XDG_STATE_HOME:-$HOME/.local/state}"
export XDG_CACHE_HOME="${XDG_CACHE_HOME:-$HOME/.cache}"
# Import all XDG set by xdg-config-update autorun
# THey're not exporting by default
. $XDG_CONFIG_HOME/user-dirs.dirs