; run once performing compile at install of tree-sitter
(mapc #'treesit-install-language-grammar (mapcar #'car treesit-language-source-alist))

