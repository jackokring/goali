;(make-symbolic-link ".config/emacs" "~/.emacs.d")
(require 'package)
(add-to-list 'package-archives '("melpa" . "https://melpa.org/packages/"))

;causal calc mode help add later version emacs
(require 'casual)
(define-key calc-mode-map (kbd "C-o") 'casual-main-menu)

;;; Code:
(eval-when-compile (require 'use-package))

; ts-grammar
(setq treesit-language-source-alist
   '((bash "https://github.com/tree-sitter/tree-sitter-bash")
     (cmake "https://github.com/uyha/tree-sitter-cmake")
     (css "https://github.com/tree-sitter/tree-sitter-css")
     (elisp "https://github.com/Wilfred/tree-sitter-elisp")
     (go "https://github.com/tree-sitter/tree-sitter-go")
     (gomod "https://github.com/camdencheek/tree-sitter-go-mod")
     (html "https://github.com/tree-sitter/tree-sitter-html")
     (javascript "https://github.com/tree-sitter/tree-sitter-javascript" "master" "src")
     (json "https://github.com/tree-sitter/tree-sitter-json")
     (make "https://github.com/alemuller/tree-sitter-make")
     (markdown "https://github.com/ikatyang/tree-sitter-markdown")
     (python "https://github.com/tree-sitter/tree-sitter-python")
     (toml "https://github.com/tree-sitter/tree-sitter-toml")
     (tsx "https://github.com/tree-sitter/tree-sitter-typescript" "master" "tsx/src")
     (typescript "https://github.com/tree-sitter/tree-sitter-typescript" "master" "typescript/src")
     (yaml "https://github.com/ikatyang/tree-sitter-yaml")))

;; do once to build grammar .so files
;(mapc #'treesit-install-language-grammar (mapcar #'car treesit-language-source-alist))

; mode remaps
(setq major-mode-remap-alist
 '((yaml-mode . yaml-ts-mode)
   (bash-mode . bash-ts-mode)
   (js2-mode . js-ts-mode)
   (typescript-mode . typescript-ts-mode)
   (json-mode . json-ts-mode)
   (css-mode . css-ts-mode)
   (go-mode . go-ts-mode) ; added go
   (gomod-mode . gomod-ts-mode)
   (python-mode . python-ts-mode)))

; start server for emacsclient calls use C-x # for sve exit next edit
(server-start)
(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(desktop-save-mode t)
 '(package-selected-packages '(markdown-mode which-key go-mode casual)))
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 )

; highlights max
(setq font-lock-maximum-decoration t)

; icicles help S-TAB
;(add-to-list 'load-path "~/.emacs.d/icicles/")
;(require 'icicles)
;(icy-mode 1)
(which-key-mode)
