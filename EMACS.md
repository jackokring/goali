# Emacs `init.el`

It's in the `extras-backup` folder linked on the `README.md`. Try `./extras-install.sh`. *WARNING*: Read `EXTRAS-README.md`.

## Control Codes
 * `C-a` Select all
 * `C-b` Special prefix *see below*
 * `C-c` Copy
 * `C-S-c ?` User command `?`. Or `C-c ?` if there is no text selection. Prefix `mode-specific-map`. See `C-b C-c ?` and file `packages/c-map.el`
   * Only upper and lower case letters are allowed for "user" customizations
 * `C-f` Find. Also forward find again. `C-r` recusive edit in find. See `C-y`
 * `C-n` New/Open. `C-o` does blank line inserts
 * `C-p` List packages
 * `C-S-p` Install language grammar
 * `C-t` Start terminal application
 * `C-q` Quit
 * `C-s` Save
 * `C-v` Paste
 * `C-w` Beginning of line
 * `C-x` Cut
 * `C-S-x ?` Execute `?`. Or `C-x ?` if there is no text selection. See `C-b C-x ?`
   * After execute `\` behaves as though `C-b C-x C-\` was entered
 * `C-y` Exit recursive edit
 * `C-\` Another special prefix `custom-escape-map`. See file `packages/escape-map.el`
## Alt Codes
 * `M-c` Capitalize letter and advance to next word
 * `M-v` Another special prefix `custom-v-map`. See file `packages/v-map.el`
 * `M-x` It's still an *Emacs* classic. Ever needed to evaluate a *LISP* symbol? The fun starts here.
 * `M-<Up/Down>` Buffers *Up*: Previous, *Down*: Next
 * `M-S-<Cursor>` Windows *Up*: Close, *Down*: New Split Below, *Left*: Next, *Right*: New Split Right 
## After `C-b` Special Prefix
 * `C-b` Kill buffer
 * `C-c ?` User command `?` if text is selected
 * `C-x ?` Execute if text is selected
 * `\` Use special prefix `C-\`
 * Any difficult to enter terminal control codes `C-<key>` may map from `C-b <key>`
## Notes on Some keyboards
 * `AltGr` produces special characters on some national keymaps
 * On Chromebooks the `Everything` key can produce `Super` sequences `S-s-`, `C-s-` and `M-s-`if the other modifier is pressed first

### I hope you find it makes Emacs better, not e`vi`l.
Maybe I should add extra things to make terminal mode better, but ... and `tmux`.
Also a `.bashrc` is included.
