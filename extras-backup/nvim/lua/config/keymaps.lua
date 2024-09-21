-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set:
-- https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here

-- <cmd> or : is command mode in an action
-- Some various shifts possibly free
-- Alt sends an escape prefix but also seems to check modifier state
-- :map xxx<cr> looks up mapping of a binding to xxx

-- N.B. Builtins use vim.fn prefix
local f = vim.fn
local a = vim.api
local k = vim.keymap.set

-- lua_ls LSP is slightly late binding on warning of unsed function name in dict

-- first letter of name must be UPPERCASE
-- this then allows ":Com args<cr>"
-- this was considered better than allowing functions in nikey and ninkey
-- as it also allow manual execution of such functions
local function com(name, desc, func)
  a.nvim_create_user_command(name, func, { desc = desc })
end

-- for the func in the command registration com to get args
local function args(opts)
  -- return table of arg strings from opts argument
  return opts.fargs
end

-- action is function or key string (maybe recursive, careful)
local function nkey(seq, desc, action)
  k("n", seq, action, { desc = desc })
end

-- also defines for i but ends with n mode (no func use com)
local function ninkey(seq, desc, action)
  nkey(seq, desc, action)
  k("i", seq, "<esc>" .. action, { desc = desc })
end

-- remains in mode i if in i
local function nikey(seq, desc, action)
  -- can't rely on control codes below being nothing in n mode
  nkey(seq, desc, action)
  -- escape for one action step to normal mode
  k("i", seq, "<C-\\><C-O>" .. action, { desc = desc })
end

--==============================================================================
-- Bare Sparse Escape (Not in use)
-- abcdefghijklmnopqrstuvwxyz
-- ABCDEFGHIJKLMNOPQRSTUVWXYZ
nkey("\\a", "", "")

--==============================================================================
-- Leader Space (Many used, see used by pressing <space> in normal mode)
-- adijkmnopvyz
-- ABCFGIJMNOPQRSTUVWXYZ

--==============================================================================
-- Control
-- Can be in insert mode as wrapped <esc> .. i by <C-\><C-O> or just <esc>
-- Perculiar shift combination needed singleton
ninkey("<C-\\><C-\\>", "Revert Buffer to Baseline", ":e!<cr>")
-- apparently treminal built in does not do terminal <C-/> works
-- this maybe to do with Chromebook passthrough but it becomes <C-_>
-- ABCDEFGIMOPQTUVXYZ
-- normal mode free to launch rofi, as register recall in i mode
nkey("<C-R>", "Open Rofi Combi", ":!rofi -show combi<cr>")
-- like a fast save all <C-S> is just save one file
nikey("<C-W>", "Write Quick All", ":wall<cr>")
