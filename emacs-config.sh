#!/usr/bin/bash
BASE="~/.emacs.d"
OPEN="EMACS.md $BASE/init.el $BASE/packages/externals.el $BASE/packages/escape-map.el $BASE/packages/c-map.el $BASE/packages/v-map.el $BASE/packages/tree-sitter.el"
emacs $OPEN 2>/dev/null&
echo "Don't forget to ./freeze.sh the changes after."
