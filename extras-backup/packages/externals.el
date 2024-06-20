;;; Externals
; mild bracket tinting: rainbow-delimiters
(add-hook 'prog-mode-hook #'rainbow-delimiters-mode)

; theme
(load-theme 'gruvbox-dark-medium t)

; enable which key help browsing
(require 'which-key)
(which-key-mode)
