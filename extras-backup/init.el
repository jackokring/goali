;; N.B. Terminal mode Emacs can't capture all keybindings automatically

; put this HERE
(add-to-list 'load-path "~/.emacs.d/packages")

;; This symbolic link although suggested causes a duplication error
;(make-symbolic-link ".config/emacs" "~/.emacs.d")

;; Emacs manual online and mepla packages stable
; https://www.gnu.org/software/emacs/manual/html_node/emacs/index.html
; You will find this useless unless you have a very upto date Emacs version
(require 'package)
(add-to-list 'package-archives '("melpa-stable" . "https://stable.melpa.org/packages/"))

; C- and ^ are used interchangably as the control key modifier throughout this file

;; Sensible modern defaults
(setq font-lock-maximum-decoration t)
(show-paren-mode 1)
; Can use C-S-x and C-S-c for execute and user keybinds
; PgUp and PgDn replace C-v or use C-<up> or C-<down>
(cua-mode t)
(global-unset-key (kbd "C-S-v")) ; unset "shell paste" bad to encourage stuff not work in terminal 

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
(global-set-key (kbd "C-/") 'comment-or-uncomment-line-or-region)

;; C-a marks the whole buffer (a)ll
(global-set-key (kbd "C-a") 'mark-whole-buffer)

;; cursor remaps (b)ack, (f)orward, (p)revious, (n)ext reusable cursors?
; This is an obvious improvement as the cursor keys and mouse work fine
; muscle memory just makes for a good experience

; ^n map (n)ext -> (n)ew
(global-set-key (kbd "C-n") 'find-file) ; also does open style functionality
; yes ... editing these comments really put C-o on the map ;D

; ^p map (p)revious -> (p)ackages
(global-set-key (kbd "C-p") 'list-packages) ; easier than the manual

;; Custom maps used
; My ^b map (b)ack -> (b)old -> (b)e
; This is not "user" commands, but modified CUA easy extensions
(define-prefix-command 'custom-b-map)
; Seems escape is strange as a first character
(define-prefix-command 'custom-escape-map)
(define-prefix-command 'custom-v-map)
; C-x = system, C-c = user "letters only", C-\ = as a prefix, it's got previous backing
; M-v -> don't encourage an opposite now working as paste
(global-set-key (kbd "C-\\") 'custom-escape-map)
(global-set-key (kbd "M-v") 'custom-v-map) ; page up opposite to C-v page down replaced by paste

;; technically all C-c should be user defined, but prefix C-b
; B, X and C are handled different as a SIGINT, plus 2 passthroughs
; for Bad Buffer, eXecute and Command after C-b
; the rest are for easy intercept requirements
(global-set-key (kbd "C-b") 'custom-b-map)

; might as well have a new beginning of line as select all CUA on C-a
(global-set-key (kbd "C-w") 'move-beginning-of-line) ; WE remap beginning/end of line
; stop those pasty yanking fingers
(global-set-key (kbd "C-y") (kbd "C-M-c")) ; Y combinator exit recursive edit remap Yλ upside down of a join for a longer end?

;; Extend my custom-b-map ... tmux shortcut key too
; ^ terminals via stty -a => c \ u d q s z r w v o
; Those control keys might not be possible to feed into Emacs
; as terminal may filter them, so use ^B prefix
(define-key custom-b-map (kbd "C-c") mode-specific-map) ; as C-c is copy
(define-key custom-b-map (kbd "C-x") ctl-x-map) ; execute as C-x is cut
; a few extra conveiniences
(define-key custom-b-map (kbd "C-b") 'kill-buffer) ; C-b b kill "bad" buffer (bibi gun)

; terminal catch extras beyond C-c SIGINT -> simpler terminal requirement
(define-key custom-b-map (kbd "\\") 'custom-escape-map) ; C mode line endings appears like only found C-x C-\
(define-key ctl-x-map (kbd "\\") (kbd "C-x C-\\")) ; just to allow THE code through

; add any funny shell business here to retarget control key combinations
(define-key custom-b-map (kbd "u") (kbd "C-u")) ; universal-argument
(define-key custom-b-map (kbd "d") (kbd "C-d")) ; cursor delete right
(define-key custom-b-map (kbd "q") (kbd "C-q")) ; quit
(define-key custom-b-map (kbd "s") (kbd "C-s")) ; save
(define-key custom-b-map (kbd "z") (kbd "C-z")) ; undo
(define-key custom-b-map (kbd "r") (kbd "C-r")) ; recursive edit, exit by C-M-c
(define-key custom-b-map (kbd "w") (kbd "C-w")) ; REMAP of C-a beginning of line, copy? C-y is paste too
(define-key custom-b-map (kbd "v") (kbd "C-v")) ; paste
(define-key custom-b-map (kbd "o") (kbd "C-o")) ; open line
; no ijlm circling the K-ill to end of line
;; End of the B map

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

;; Some other more logical keys for CUA users
;; my remaps for ^f, ^s and ^q
; find (f)orward -> (f)ind
(global-set-key (kbd "C-f") 'isearch-forward)
; special for use insode incremental active searches
(define-key isearch-mode-map (kbd "C-f") 'isearch-repeat-forward)
; save (s)earch -> (s)ave
(global-set-key (kbd "C-s") 'save-buffer)
; quit insert code -> (q)uit
(global-set-key (kbd "C-q") 'save-buffers-kill-terminal)

;; Buffer and window manipulations 
; buffer navigation sometimes intercepted
; left/right sometimes won't work (depends on buffer modes)
(global-set-key (kbd "M-<up>") 'previous-buffer)
(global-set-key (kbd "M-<down>") 'next-buffer)

;; with shift <up> is like backspace, <left> is like tab (key pointing arrows)
; useful for managing window panes
(global-set-key (kbd "M-S-<left>") 'other-window)
(global-set-key (kbd "M-S-<right>") 'split-window-right)
(global-set-key (kbd "M-S-<up>") 'delete-window)
(global-set-key (kbd "M-S-<down>") 'split-window-below)

;;; Externals related to packages
(load "externals.el")

;;; Custom addition by Emacs DON'T EDIT BELOW
(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(cua-mode t)
 '(custom-safe-themes
   '("72ed8b6bffe0bfa8d097810649fd57d2b598deef47c992920aef8b5d9599eefe" default))
 '(package-selected-packages
   '(rainbow-mode autothemer gruvbox-theme helm-core ini-mode js2-mode lua-mode markdown-mode nix-mode org pdf-tools rust-mode which-key go-mode html-to-markdown json-mode python-mode yaml-mode rainbow-delimiters)))
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 )
