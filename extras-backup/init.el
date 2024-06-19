;; N.B. Terminal mode Emacs can't capture all keybindings automatically

;; This symbolic link although suggested causes a duplication error
;(make-symbolic-link ".config/emacs" "~/.emacs.d")

;; Emacs manual online and mepla packages stable
; https://www.gnu.org/software/emacs/manual/html_node/emacs/index.html
; You will find this useless unless you have a very upto date Emacs version
(require 'package)
(add-to-list 'package-archives '("melpa-stable" . "https://stable.melpa.org/packages/"))

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
; technically all C-c should be user defined, but prefix C-b
(define-prefix-command 'custom-c-map)
(global-set-key (kbd "C-b") 'custom-b-map)
(global-set-key (kbd "C-S-c") 'custom-c-map) ; use shift pass thru
; Extend my custom-b-map ... tmux shortcut key too
; ^ terminals via stty -a => c \ u d q s z r w v o
; Those control keys might not be possible to feed into Emacs
; as terminal may filter them, so use ^B prefix
(define-key custom-b-map (kbd "c") 'custom-c-map) ; as C-c is copy
(define-key custom-b-map (kbd "x") (kbd "C-S-x") ; execute as C-x is cut
(define-key custom-b-map (kbd "\\") (kbd "C-\\"))
(define-key custom-b-map (kbd "u") (kbd "C-u"))
(define-key custom-b-map (kbd "d") (kbd "C-d"))
(define-key custom-b-map (kbd "q") (kbd "C-q")) ; quit
(define-key custom-b-map (kbd "s") (kbd "C-s")) ; save
(define-key custom-b-map (kbd "z") (kbd "C-z")) ; undo
(define-key custom-b-map (kbd "r") (kbd "C-r"))
(define-key custom-b-map (kbd "w") (kbd "C-w"))
(define-key custom-b-map (kbd "v") (kbd "C-v")) ; paste
(define-key custom-b-map (kbd "o") (kbd "C-o"))

;; User custom-c-map ^C usually but adapted for "user" commands (^B c)


;; my remaps for ^f, ^s and ^q
; find (f)orward -> (f)ind
(global-set-key (kbd "C-f") 'isearch-forward)
; save (s)earch -> (s)ave
(global-set-key (kbd "C-s") 'save-buffer)
; quit insert code -> (q)uit
(global-set-key (kbd "C-q") 'save-buffers-kill-terminal)

;; buffer navigation
(global-set-key (kbd "M-<left>") 'previous-buffer)
(global-set-key (kbd "M-<right>") 'next-buffer)
; a nice home list
(global-set-key (kbd "M-<up>") 'list-buffers)
; end some buffers
(global-set-key (kbd "M-<down>") 'kill-some-buffers)

;; with shift <up> is like backspace, <left> is like tab
; useful for managing window panes
(global-set-key (kbd "M-S-<left>") 'other-window)
(global-set-key (kbd "M-S-<right>") 'split-window-right)
(global-set-key (kbd "M-S-<up>") 'delete-window)
(global-set-key (kbd "M-S-<down>") 'split-window-below)
