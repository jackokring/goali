-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here

-- <cmd> or : is command mode in an action
-- Some various shifts possibly free
-- Alt sends an escape prefix

local function nkey(seq, action)
  vim.keymap.set({ "n", "v" }, seq, action)
end

local function nikey(seq, action)
  vim.keymap.set({ "n", "v", "i" }, seq, action)
end

-- Bare Sparse Escape (Not in use)
-- abcdefghijklmnopqrstuvwxyz
-- ABCDEFGHIJKLMNOPQRSTUVWXYZ
nkey("\\a", "")

-- Leader Space (Many used, see use pressing <space> in normal mode)
-- adhijkmnoprtvyz
-- ABCFGHIJMNOPQRSTUVWXYZ
nkey("<Leader>a", "")

-- Control (Lowercase RESERVED for plugins with no control, uppercase free with no control but shifted)
-- Can be in insert mode sometimes as control and not <c-v> literal prefixed
-- ABCDEFHIJKLMOPQRSTUVWXYZ (with control as easiest to finger)
-- Perculiar shift combination needed
nikey("<C-_>", "")
-- NOT <C-N> or <C-G> but rest of controls and not lowercase
nikey("<C-\\><C-A>", "")
-- Uppercase with shift
nkey("<C-\\><S-A>", "")
