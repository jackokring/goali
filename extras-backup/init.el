(make-symbolic-link ".config/emacs" "~/.emacs.d")
(require 'package)
(add-to-list 'package-archives '("melpa-stable" . "https://stable.melpa.org/packages/"))

;causal calc mode help add later version emacs
(require 'casual)
(define-key calc-mode-map (kbd "C-o") 'casual-main-menu)

; highlights max
(setq font-lock-maximum-decoration t)

(which-key-mode)
