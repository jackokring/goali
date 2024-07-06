;; tree-sitter handling
(setq major-mode-remap-alist
      '(
; redirect to *-ts-mode modes
	(js2-mode . javascript-ts-mode)
	(lua-mode . lua-ts-mode)
	(markdown-mode . markdown-ts-mode)
	(nix-mode . nix-ts-mode)
	(rust-mode . rust-ts-mode)
	(go-mode . go-ts-mode)
	(json-mode . json-ts-mode)
	(python-mode . python-ts-mode)
	(yaml-mode . yaml-ts-mode)
; close setq
	))

;; Rerun hooks
; Mahmoud Adam - Emacs Wiki
(defun run-non-ts-hooks ()
  (let ((major-name (symbol-name major-mode)))
    (when (string-match-p ".*-ts-mode" major-name)
      (run-hooks (intern (concat (replace-regexp-in-string "-ts" "" major-name) "-hook"))))))

(add-hook 'prog-mode-hook 'run-non-ts-hooks)
