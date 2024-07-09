;;; Externals
;; add-hook applied first executed last
;; so keeps the delimiters in rainbow mode
;; it appears autoloads are no (require ...) needed

; mild bracket tinting: rainbow-delimiters
(add-hook 'prog-mode-hook #'rainbow-delimiters-mode)

; rainbow mode inline color defs actual
(add-hook 'prog-mode-hook #'rainbow-mode)
(add-hook 'text-mode-hook #'rainbow-mode)
(add-hook 'elisp-mode-hook #'rainbow-mode)

; theme
(load-theme 'gruvbox-dark-medium t)

; enable which key help browsing
(which-key-mode)

; form feed separation ^L
(global-form-feed-mode)
