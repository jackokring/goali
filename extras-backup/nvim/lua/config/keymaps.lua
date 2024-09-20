-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here

-- <cmd> or : is command mode in an action
-- Some various shifts possibly free
-- Alt sends an escape prefix

-- N.B. Builtins use vim.fn prefix
-- lua_ls LSP is slightly late binding on warning of unsed function name in dict

-- action is function or key string (maybe recursive, careful)
local function nkey(seq, desc, action)
  vim.keymap.set("n", seq, action, { desc = desc })
end

local function nikey(seq, desc, action)
  vim.keymap.set("n", seq, action, { desc = desc })
  -- escape for one action step to normal mode
  vim.keymap.set("i", seq, "<C-\\><C-O>" .. action, { desc = desc })
end

-- Bare Sparse Escape (Not in use)
-- abcdefghijklmnopqrstuvwxyz
-- ABCDEFGHIJKLMNOPQRSTUVWXYZ
nkey("\\a", "", "")

-- Leader Space (Many used, see use by pressing <space> in normal mode)
-- adijkmnopvyz
-- ABCFGIJMNOPQRSTUVWXYZ
nkey("<Leader>r", "Open Rofi Combi", ":!rofi -show combi<cr>")
nkey("<Leader>t", "Terminal", ":term<cr>i")

-- Control
-- Can be in insert mode as wrapped <esc> .. i by <C-\><C-O>
-- Perculiar shift combination needed singleton
nikey("<C-_>", "", "")
-- ABCDEFGHIJKLMNOPQRSTUVXYZ
nikey("<C-W>", "Write Quick", ":w<cr>")
