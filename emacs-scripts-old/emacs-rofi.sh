#!/usr/bin/bash
. ./xdg.sh
emacs $XDG_CONFIG_HOME/rofi/* 2>/dev/null&
echo "Don't forget to ./freeze.sh the changes after."
