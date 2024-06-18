;(make-symbolic-link ".config/emacs" "~/.emacs.d")
(require 'package)
(add-to-list 'package-archives '("melpa-stable" . "https://stable.melpa.org/packages/"))

;; defaults
(setq font-lock-maximum-decoration t)
(show-paren-mode 1)
(cua-mode t)

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
;; C-a marks the whole buffer
(global-set-key (kbd "C-a") 'mark-whole-buffer)

;; just mine (b)ack, (p)revious, (n)ext reusable cursors?

; n map
(define-prefix-command 'custom-n-map)
(global-set-key (kbd "C-n") 'custom-n-map)

; p map
(define-prefix-command 'custom-p-map)
(global-set-key (kbd "C-p") 'custom-p-map)

; extend p as right hand extender (user (p)rogramming)
(define-key custom-p-map (kbd "C-a") 'mark-whole-buffer) ; Select-all.
(define-key custom-p-map (kbd "C-v") 'yank) ; Paste.
(define-key custom-p-map (kbd "C-x") 'kill-region) ; Cut.
(define-key custom-p-map (kbd "C-c") 'kill-ring-save) ; Copy.

; b map
(define-prefix-command 'custom-b-map)
(global-set-key (kbd "C-b") 'custom-b-map)

;; my remap of search and save (f)orward -> (f)ind
; find
(global-set-key (kbd "C-f") 'isearch-forward)
; save
(global-set-key (kbd "C-s") 'save-buffer)
; quit
(global-set-key (kbd "C-q") 'save-buffers-kill-terminal)
