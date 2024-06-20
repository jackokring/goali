;; N.B. Terminal mode Emacs can't capture all keybindings automatically

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
(global-unset-key (kbd "C-S-v")) ; unset shell paste

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

; ^n map (n)ext -> (n)ew
(global-set-key (kbd "C-n") 'find-file)

; ^p map (p)revious -> (p)ackages
(global-set-key (kbd "C-p") 'list-packages)

; My ^b map (b)ack -> (b)old -> (b)e
; This is not "user" commands, but modified CUA easy extensions
(define-prefix-command 'custom-b-map)
; Seems escape is strange as a first character
(define-prefix-command 'custom-escape-map)
; C-x = system, C-c = user "letters only", C-\ = as a prefix, it's got previous backing
(global-set-key (kbd "C-\\") 'custom-escape-map)
; technically all C-c should be user defined, but prefix C-b
; B, X and C are handled different as a SIGINT, plus 2 passthroughs
; for Bad Buffer, eXecute and Command after C-b
; the rest are for easy intercept requirements
(global-set-key (kbd "C-b") 'custom-b-map)
(global-set-key (kbd "C-w") 'move-beginning-of-line) ; WE remap beginning/end of line
(global-set-key (kbd "C-y") (kbd "C-M-c")) ; Y combinator exit recursion remap
; Extend my custom-b-map ... tmux shortcut key too
; ^ terminals via stty -a => c \ u d q s z r w v o
; Those control keys might not be possible to feed into Emacs
; as terminal may filter them, so use ^B prefix
(define-key custom-b-map (kbd "c") mode-specific-map) ; as C-c is copy
(define-key custom-b-map (kbd "x") ctl-x-map) ; execute as C-x is cut and no '
; terminal catch extras beyond C-c SIGINT -> simpler terminal requirement
(define-key custom-b-map (kbd "\\") 'custom-escape-map) ; C mode line endings appears like only found C-x C-\
(define-key custom-b-map (kbd "u") (kbd "C-u")) ; universal-argument
(define-key custom-b-map (kbd "d") (kbd "C-d")) ; cursor delete right
(define-key custom-b-map (kbd "q") (kbd "C-q")) ; quit
(define-key custom-b-map (kbd "s") (kbd "C-s")) ; save
(define-key custom-b-map (kbd "z") (kbd "C-z")) ; undo
(define-key custom-b-map (kbd "r") (kbd "C-r")) ; recursive edit, exit by C-M-c
(define-key custom-b-map (kbd "w") (kbd "C-w")) ; REMAP of C-a beginning of line, copy? C-y is paste too
(define-key custom-b-map (kbd "v") (kbd "C-v")) ; paste
(define-key custom-b-map (kbd "o") (kbd "C-o")) ; open line
; encourage recursive edits by making exit easier?
(define-key custom-b-map (kbd "y") (kbd "C-y")) ; end recursive edit, Y combinator hint
; a few extra conveiniences
(define-key custom-b-map (kbd "b") 'kill-buffer) ; C-b b kill "bad" buffer (bibi gun)
; no ijlm circling the K-ill to end of line
;; End of the B map

;; User custom-c-map ^C usually but adapted for "user" commands (^B c)
; N.B. "mode-specific-map" replaces custom-c-map for C-c
; also help-map for C-h
; It's the common commands of the global mode-specific-map set/unset


;; Custom escape map ^\ seems open for use too
; apparently terminal mode does not pass C-\ through

;; my remaps for ^f, ^s and ^q
; find (f)orward -> (f)ind
(global-set-key (kbd "C-f") 'isearch-forward)
; save (s)earch -> (s)ave
(global-set-key (kbd "C-s") 'save-buffer)
; quit insert code -> (q)uit
(global-set-key (kbd "C-q") 'save-buffers-kill-terminal)

;; buffer navigation sometimes intercepted
; left/right sometimes won't work (depends on buffer modes)
(global-set-key (kbd "M-<up>") 'previous-buffer)
(global-set-key (kbd "M-<down>") 'next-buffer)

;; with shift <up> is like backspace, <left> is like tab (key pointing arrows)
; useful for managing window panes
(global-set-key (kbd "M-S-<left>") 'other-window)
(global-set-key (kbd "M-S-<right>") 'split-window-right)
(global-set-key (kbd "M-S-<up>") 'delete-window)
(global-set-key (kbd "M-S-<down>") 'split-window-below)

;;; Externals
(add-to-list 'load-path "~/.emacs.d/packages")
(load "externals.el")
(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(cua-mode t)
 '(custom-safe-themes
   '("72ed8b6bffe0bfa8d097810649fd57d2b598deef47c992920aef8b5d9599eefe" default))
 '(package-selected-packages
   '(autothemer gruvbox-theme helm-core ini-mode js2-mode lua-mode markdown-mode nix-mode org pdf-tools rust-mode which-key go-mode html-to-markdown json-mode python-mode yaml-mode rainbow-delimiters)))
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 )
