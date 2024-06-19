# Emacs `init.el`

It's in the `extras-backup` folder linked on the `README.md`.

## Control Codes
 * `C-a` Select all
 * `C-b` Special prefix (see below)
 * `C-c` Copy
 * `C-S-c ?` User command ?
 * `C-f` Find
 * `C-n` New/Open (C-o does blank line inserts OK)
 * `C-p` List packages
 * `C-q` Quit
 * `C-s` Save
 * `C-v` Paste
 * `C-x` Cut
 * `C-S-x` Execute
 * `M-<up/down>` Buffers *Up*: previous, *Down*: next
 * `M-S-<cursor>` Windows *Up*: close, *Down*: new below, *Left*: next, *Right*: new right 

## After `C-b`
 * `C-b` Kill buffer
 * `c ?` User command ? (the usual command can be `C-S-c ?` also, as "copy" takes over `C-c`)
 * `x` Execute
 * Any of `q s z v` to do `C-<any>` without terminal intercept (inbuilt control commands)
 * Any of `\ u d r w` to do `C-<any>` without terminal intercept

### I hope you find it makes Emacs better, not e `vi` l.
Maybe I should add extra things to make terminal mode better, but ... and `tmux`.
