-- doris plugin loader for nvim
return {
  "jackokring/doris.nvim",
  -- **local build**
  dev = true,
  dir = "~/projects/doris.nvim",
  fallback = true,
  -- **build command**
  build = "./build.sh",
  -- **setup options**
  opts = {},
  -- **lazy load info**
  lazy = true,
  -- **on event**
  -- event = { "BufEnter", "BufEnter *.lua" },
  -- **on command use**
  -- cmd = { "cmd" },
  -- **on filetype**
  -- ft = { "lua" },
  -- **on keys**
  -- keys = {
  --    -- key tables
  --    {
  --      "<leader>ft",
  --      -- "<cmd>Neotree toggle<cr>",
  --      -- desc = "NeoTree",
  --      -- mode = "n",
  --      -- ft = "lua"
  --    },
  --  },
}
