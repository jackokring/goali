-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here

-- <cmd> or : is command mode in an action
-- Some various shifts possibly free
-- Alt sends an escape prefix

-- N.B. Builtins use vim.fn prefix
-- lua_ls LSP is slightly late binding on warning of unsed function name in dict

local function iwrap(action)
  -- stop insert mode for a bit
  -- umm, stopinsert() is a normal mode thing ...
  -- not sure if there's a reason for needing <esc>
  local cur = vim.fn.getcurpos()
  local mode = vim.fn.mode(0)
  -- kind of like the oddball insert mode : command
  if mode == "i" then
    vim.cmd("stopinsert")
  end
  vim.fn.call(action, { mode, cur })
  -- restore, ah yes, the end of line by desired column?
  if mode == "i" then
    vim.cmd("startinsert")
  end
  vim.fn.setpos(".", cur)
end

-- action is function or key string (maybe recursive, careful)
local function nkey(seq, desc, action)
  vim.keymap.set({ "n", "v" }, seq, action, { desc = desc })
end

-- action must be a function for this one
local function nikey(seq, desc, action)
  vim.keymap.set({ "n", "v", "i" }, seq, function()
    iwrap(action)
  end, { desc = desc })
end

-- Bare Sparse Escape (Not in use)
-- abcdefghijklmnopqrstuvwxyz
-- ABCDEFGHIJKLMNOPQRSTUVWXYZ
nkey("\\a", "", "")

-- Leader Space (Many used, see use by pressing <space> in normal mode)
-- adijkmnopvyz
-- ABCFGIJMNOPQRSTUVWXYZ
nkey("<Leader>r", "Open Rofi Combi", "<cmd>!rofi -show combi<cr>")
nkey("<Leader>t", "Terminal", "<cmd>term<cr>i")

-- Control (Lowercase RESERVED for plugins with no control)
-- (uppercase free with no control but shifted)
-- Can be in insert mode sometimes as control and not <c-v> literal prefixed
-- ABCDEFHIJKLMOPQRSTUVWXYZ (with control as easiest to finger)
-- Perculiar shift combination needed singleton
nikey("<C-_>", "", function(mode, cur)
  -- nil
end)
-- NOT <C-N> or <C-G> but rest of controls and not lowercase
nikey("<C-\\><C-A>", "", function(mode, cur)
  -- nil
end)
-- Uppercase and symbols with shift unused
nkey("<C-\\><S-A>", "", "")
