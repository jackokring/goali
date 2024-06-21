# `./extras-install.sh`

The script will install extras, but will ask. *WARNING*: Some yes options will overwrite your local configuration.
 * Bash - `.bashrc` is overwritten. It might lead to your terminal setup not working nice, and errors.
 * Tmux - a full replacement config. Again if you have `tmux` setup your way already, select the `No` option.
 * Emacs - a full extended CUA implementation with extra MELPA. Select `No` to keep your current config.
A *backup is made* in `~/.mess` to assist if you have butter fingers.
