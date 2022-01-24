# wezterm.nvim

Replacement of
- [vim-tmux-navigator](https://github.com/christoomey/vim-tmux-navigator) 
- [better-vim-tmux-resizer](https://github.com/RyanMillerC/better-vim-tmux-resizer)

for [wezterm](https://github.com/wez/wezterm).

## prerequisite

.bashrc / .zshrc
```
[ -n "$WEZTERM_PANE" ] && export NVIM_LISTEN_ADDRESS="/tmp/nvim$WEZTERM_PANE"
```

.config/fish/config.fish
```
if not set -q $WEZTERM_PANE
  set -x NVIM_LISTEN_ADDRESS "/tmp/nvim$WEZTERM_PANE"
end
```

I just use the original tmux plugins for vim configs.

- [vim-tmux-navigator](https://github.com/christoomey/vim-tmux-navigator) 
- [better-vim-tmux-resizer](https://github.com/RyanMillerC/better-vim-tmux-resizer)

```
vim.cmd([[
  packadd vim-tmux-navigator
  nnoremap <silent><c-h> <cmd>TmuxNavigateLeft<cr>
  nnoremap <silent><c-j> <cmd>TmuxNavigateDown<cr>
  nnoremap <silent><c-k> <cmd>TmuxNavigateUp<cr>
  nnoremap <silent><c-l> <cmd>TmuxNavigateRight<cr>

  tnoremap <c-h> <C-\><C-N><cmd>TmuxNavigateLeft<cr>
  tnoremap <c-j> <C-\><C-N><cmd>TmuxNavigateDown<cr>
  tnoremap <c-k> <C-\><C-N><cmd>TmuxNavigateUp<cr>
  tnoremap <c-l> <C-\><C-N><cmd>TmuxNavigateRight<cr>
]])

vim.cmd([[
  packadd better-vim-tmux-resizer
  let g:tmux_resizer_no_mappings = 1
  nnoremap <silent> <m-h> <cmd>TmuxResizeLeft<cr>
  nnoremap <silent> <m-j> <cmd>TmuxResizeDown<cr>
  nnoremap <silent> <m-k> <cmd>TmuxResizeUp<cr>
  nnoremap <silent> <m-l> <cmd>TmuxResizeRight<cr>
]])
```


## navigator

Currently wezterm doesn't provide api to execute wezterm action inside neovim. So we execute commandline program in wezterm which 
uses neovim remote msgpack api. 

Checkout https://github.com/wez/wezterm/discussions/995 for details/updates.

---

install
```
cd wezterm.nvim.navigator && go install
```


wezterm config
```
local wezterm = require("wezterm")
local os = require("os")

local move_around = function(window, pane, direction_wez, direction_nvim)
  local result = os.execute("env NVIM_LISTEN_ADDRESS=/tmp/nvim" .. pane:pane_id() ..  " wezterm.nvim.navigator " .. direction_nvim)
  if result then
		window:perform_action(wezterm.action({ SendString = "\x17" .. direction_nvim }), pane)
  else
		window:perform_action(wezterm.action({ ActivatePaneDirection = direction_wez }), pane)
  end
end

wezterm.on("move-left", function(window, pane)
	move_around(window, pane, "Left", "h")
end)

wezterm.on("move-right", function(window, pane)
	move_around(window, pane, "Right", "l")
end)

wezterm.on("move-up", function(window, pane)
	move_around(window, pane, "Up", "k")
end)

wezterm.on("move-down", function(window, pane)
	move_around(window, pane, "Down", "j")
end)
```


wezterm mapping
```
-- pane move(vim aware)
{ key = "h", mods = "CTRL", action = wezterm.action({ EmitEvent = "move-left" }) },
{ key = "l", mods = "CTRL", action = wezterm.action({ EmitEvent = "move-right" }) },
{ key = "k", mods = "CTRL", action = wezterm.action({ EmitEvent = "move-up" }) },
{ key = "j", mods = "CTRL", action = wezterm.action({ EmitEvent = "move-down" }) },
```

## resizer

```
local vim_resize = function(window, pane, direction_wez, direction_nvim)
	local result = os.execute(
		"env NVIM_LISTEN_ADDRESS=/tmp/nvim"
			.. pane:pane_id()
			.. " "
			.. homedir
			.. "/bin/"
			.. "wezterm.nvim.navigator "
			.. direction_nvim
	)
	if result then
		window:perform_action(wezterm.action({ SendString = "\x1b" .. direction_nvim }), pane)
	else
		window:perform_action(wezterm.action({ ActivatePaneDirection = direction_wez }), pane)
	end
end

wezterm.on("resize-left", function(window, pane)
	vim_resize(window, pane, "Left", "h")
end)

wezterm.on("resize-right", function(window, pane)
	vim_resize(window, pane, "Right", "l")
end)

wezterm.on("resize-up", function(window, pane)
	vim_resize(window, pane, "Up", "k")
end)

wezterm.on("resize-down", function(window, pane)
	vim_resize(window, pane, "Down", "j")
end)
```

```
{ key = "h", mods = "ALT", action = wezterm.action({ EmitEvent = "resize-left" }) },
{ key = "l", mods = "ALT", action = wezterm.action({ EmitEvent = "resize-right" }) },
{ key = "k", mods = "ALT", action = wezterm.action({ EmitEvent = "resize-up" }) },
{ key = "j", mods = "ALT", action = wezterm.action({ EmitEvent = "resize-down" }) },
```

