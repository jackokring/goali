-- Options are automatically loaded before lazy.nvim startup
-- Default options that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/options.lua
-- Add any additional options here

-- column line of 81 to set limit line
vim.o.colorcolumn = "81"
-- vim.opt.colorcolumn = 80
vim.o.timeoutlen = 500
-- for key sequence determinacy
vim.o.scrolloff = 16
-- kept screen cursor in centre when scrolling
-- that really help when expanding the which key help
