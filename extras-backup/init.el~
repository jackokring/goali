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

;; C-a marks the whole buffer (a)ll
(global-set-key (kbd "C-a") 'mark-whole-buffer)

;; cursor remaps (b)ack, (f)orward, (p)revious, (n)ext reusable cursors?

; n map (n)ext -> (n)ew
(define-prefix-command 'find-file)

; p map (p)revious -> (p)rogramming
(define-prefix-command 'custom-p-map)
(global-set-key (kbd "C-p") 'custom-p-map)

; extend p as right hand extender NOT PRINTING
(define-key custom-p-map (kbd "C-a") 'mark-whole-buffer) ; Select-all.
(define-key custom-p-map (kbd "C-v") 'yank) ; Paste.
(define-key custom-p-map (kbd "C-x") 'kill-region) ; Cut.
(define-key custom-p-map (kbd "C-c") 'kill-ring-save) ; Copy.

; b map (b)ack -> (b)old -> ??
(define-prefix-command 'custom-b-map)
(global-set-key (kbd "C-b") 'custom-b-map)

;; my remaps
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
