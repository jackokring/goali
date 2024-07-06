;; N.B. Terminal mode Emacs can't capture all keybindings automatically

; put this HERE
(add-to-list 'load-path "~/.emacs.d/packages")

;; This symbolic link although suggested causes a duplication error
;(make-symbolic-link ".config/emacs" "~/.emacs.d")

;; Emacs manual online and mepla packages stable
; https://www.gnu.org/software/emacs/manual/html_node/emacs/index.html
; This is the only require required for bootstrap of packages
; custom-set-variables at the end of the file sets up autoload
; for all the packages installed
(require 'package)
(add-to-list 'package-archives '("melpa-stable" . "https://stable.melpa.org/packages/"))

; ===========================================================
;;; Externals related to other external not included packages
; keeps the basic init.el clean of code which is not default
; due to early kbd-command-find binding, load externals first
(load "externals.el")
; for when it's available
(load "tree-sitter.el")

; C- and ^ are used interchangably as the control key modifier throughout this file

;; Sensible modern defaults
(setq font-lock-maximum-decoration t)
(show-paren-mode 1)
; Can use C-S-x and C-S-c for execute and user keybinds
; PgUp and PgDn replace C-v or use C-<up> or C-<down>
(cua-mode t)
; Set up a default set of tabs for differing layouts
; You can still change the buffer in each tab
(tab-bar-mode)
; A generic tab line in "the" buffer? Interactive?
; (tab-line-mode)

;; Not sure if this might change or has a newer version
(defun kbd-command-find (bind)
  "Find function bound to."
  (key-binding (kbd bind)))

;; mapping functions as defined https://www.gnu.org/software/emacs/manual/html_node/emacs/Init-Rebinding.html
(defun keymap-set (map key bound)
  "Do map key bind."
  (define-key map (kbd key)
    (if (stringp bound) (kbd-command-find bound) bound)
  ))

(defun keymap-global-set (key bound)
  "Do global key bind."
  (keymap-set (current-global-map) key bound))

(keymap-global-set "C-S-v" nil) ; unset "shell paste" bad to encourage stuff not work in terminal

;; Define C-/ to comment and uncomment regions and lines
(defun comment-or-uncomment-line-or-region ()
  "Comments or uncomments the region or the current line if there's no active region."
  (interactive)
  (let (beg end)
    (if (region-active-p)
      (setq beg (region-beginning) end (region-end))
      (setq beg (line-beginning-position) end (line-end-position)))
    (comment-or-uncomment-region beg end)
    (next-line)))
(keymap-global-set "C-/" 'comment-or-uncomment-line-or-region)

;; C-a marks the whole buffer (a)ll
(keymap-global-set "C-a" 'mark-whole-buffer)

;; cursor remaps (b)ack, (f)orward, (p)revious, (n)ext reusable cursors?
; This is an obvious improvement as the cursor keys and mouse work fine
; muscle memory just makes for a good experience

; ^n map (n)ext -> (n)ew
(keymap-global-set "C-n" 'find-file) ; also does open style functionality
; yes ... editing these comments really put C-o on the map ;D

; ^p map (p)revious -> (p)ackages
(keymap-global-set "C-p" 'list-packages) ; easier than the manual
(keymap-global-set "C-S-p" 'treesit-install-language-grammar) ; get grammar

;; Custom maps used
; My ^b map (b)ack -> (b)old -> (b)e
; This is not "user" commands, but modified CUA easy extensions
(define-prefix-command 'custom-b-map)
; Seems escape is strange as a first character
(define-prefix-command 'custom-escape-map)
(define-prefix-command 'custom-v-map)

; C-x = system, C-c = user "letters only", C-\ = as a prefix, it's got previous backing
; M-v -> don't encourage an opposite now working as paste
(keymap-global-set "C-\\" 'custom-escape-map)
; an extra map prefix
(keymap-global-set "M-v" 'custom-v-map) ; page up opposite to C-v page down replaced by paste

;; technically all C-c should be user defined, but prefix C-b
; B, X and C are handled different as a SIGINT, plus 2 passthroughs
; for Bad Buffer, eXecute and Command after C-b
; the rest are for easy intercept requirements
(keymap-global-set "C-b" 'custom-b-map)

; might as well have a new beginning of line as select all CUA on C-a
(keymap-global-set "C-w" 'move-beginning-of-line) ; WE remap beginning/end of line
(keymap-global-set "M-w" 'backward-sentence) ; sentence motion

; stop those pasty yanking fingers
; It's also the first macro style mapping. Beware recursive errors of calling oneself
(keymap-global-set "C-y" 'exit-recursive-edit) ; Y combinator exit recursive edit remap Yλ upside down of a join for a longer end?

; ===========================================
;; Some other more logical keys for CUA users
;; my remaps for ^f, ^s and ^q
; find (f)orward -> (f)ind
(keymap-global-set "C-f" 'isearch-forward)
; special for use insode incremental active searches
(keymap-set isearch-mode-map "C-f" 'isearch-repeat-forward)

; save (s)earch -> (s)ave
(keymap-global-set "C-s" 'save-buffer)
; and a save as, as I find it useful for saving a new template with select all DEL
(keymap-global-set "C-S-s" 'write-file)

; quit insert code -> (q)uit
(keymap-global-set "C-q" 'save-buffers-kill-terminal)

;; Extend my custom-b-map ... tmux shortcut key too
; ^ terminals via stty -a => c \ u d q s z r w v o
; Those control keys might not be possible to feed into Emacs
; as terminal may filter them, so use ^B prefix
(keymap-set custom-b-map "C-c" mode-specific-map) ; as C-c is copy
(keymap-set custom-b-map "C-x" ctl-x-map) ; execute as C-x is cut
; a few extra conveiniences
(keymap-set custom-b-map "C-b" 'kill-buffer) ; C-b C-b kill "bad" buffer (bibi gun), tmux -> (C-b)^4
; terminal catch extras beyond C-c SIGINT -> simpler terminal requirement
(keymap-set custom-b-map "\\" 'custom-escape-map) ; C mode line endings appears like only found C-\ use

; C-i -> TAB
; C-j -> LFD new line
; C-l -> FF (global-form-feed-mode) section seperations
; C-m -> RET enter implies a new line is entered
; C-[ -> ESC also is a M- prefix

; use macro form as direct binds to actions wouldn't map on changing that "under" bound
; add any funny shell business here to retarget control key combinations
; use string implies early bind to command on key replacement
(keymap-set custom-b-map "u" "C-u") ; universal-argument
(keymap-set custom-b-map "d" "C-d") ; cursor delete right
(keymap-set custom-b-map "q" "C-q") ; quit
(keymap-set custom-b-map "s" "C-s") ; save
(keymap-set custom-b-map "z" "C-z") ; undo
(keymap-set custom-b-map "r" "C-r") ; recursive edit, exit by C-M-c -> C-y
(keymap-set custom-b-map "w" "C-w") ; REMAP of C-a beginning of line, copy? C-y is paste too
(keymap-set custom-b-map "v" "C-v") ; paste
(keymap-set custom-b-map "o" "C-o") ; open line
; no ijlm circling the K-ill to end of line
; I mean some might try binds for cursor movement, but I'm sure that needs a raw keyboard terminal map ^M = CR
;; End of the B map?

; So still possible binds without 
; C-@ -> NUL
; C-^ -> RS
; C-\] -> GS
; C-_ -> US

; ====================================================================
;; User custom-c-map ^C usually but adapted for "user" commands (^B c)
; N.B. "mode-specific-map" replaces custom-c-map for C-c
; also help-map for C-h
; It's the common commands of the global mode-specific-map set/unset
(load "c-map.el")

;; Open custom-escape-map ^\ seems open for use too
; apparently terminal mode does not pass C-\ through
(load "escape-map.el")

;; Open custom-v-map M-v seems open for use too
; not to encourage an opposite of a taken by paste
(load "v-map.el")

; ================================
;; Buffer and window manipulations 
; buffer navigation sometimes intercepted
; left/right sometimes won't work (depends on buffer modes)
(keymap-global-set "M-<up>" 'previous-buffer)
(keymap-global-set "M-<down>" 'next-buffer)

; alreagy mapped?
;(keymap-global-set "C-M-<up>" 'backward-list)
;(keymap-global-set "C-M-<down>" 'forward-list)
;(keymap-global-set "C-M-<left>" 'backward-sexp)
;(keymap-global-set "C-M-<right>" 'forward-sexp)

;; with shift <up> is like backspace, <left> is like tab (key pointing arrows)
; useful for managing window panes
(keymap-global-set "M-S-<left>" 'other-window)
(keymap-global-set "M-S-<right>" 'split-window-right)
(keymap-global-set "M-S-<up>" 'delete-window)
(keymap-global-set "M-S-<down>" 'split-window-below)

; ========================================================================
;;; Custom addition by Emacs DON'T EDIT BELOW, IT'S AUTOMAGICALLY INSERTED

(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(cua-mode t)
 '(custom-safe-themes
   '("72ed8b6bffe0bfa8d097810649fd57d2b598deef47c992920aef8b5d9599eefe" default))
 '(package-selected-packages
   '(nix-ts-mode form-feed rainbow-mode autothemer gruvbox-theme helm-core ini-mode js2-mode lua-mode markdown-mode nix-mode org pdf-tools rust-mode which-key go-mode html-to-markdown json-mode python-mode yaml-mode rainbow-delimiters)))
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 )
