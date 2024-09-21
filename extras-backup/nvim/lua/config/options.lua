-- Options are automatically loaded before lazy.nvim startup
-- Default options that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/options.lua
-- Add any additional options here

local o = vim.o
-- column line of 81 to set limit line
o.colorcolumn = "81"
-- vim.opt.colorcolumn = 80
o.timeoutlen = 500
-- for key sequence determinacy
o.scrolloff = 16
-- kept screen cursor in centre when scrolling
-- that really helps when expanding the which key help
