# wezterm.nvim

This replaces [vim-tmux-navigator](https://github.com/christoomey/vim-tmux-navigator) for [wezterm](https://github.com/wez/wezterm).
Currently wezterm doesn't provide api to execute wezterm action inside neovim. So we execute commandline in wezterm which 
uses neovim remote msgpack api. 

Checkout https://github.com/wez/wezterm/discussions/995 for details/updates.

---

install
```
cd wezterm.nvim.navigator && go install
```


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
