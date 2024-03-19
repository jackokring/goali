---
layout: post
title:  "What's Next?"
date:   2024-03-19 16:33:00 +0000
categories: feature update
---
So a few scripts `require.sh`/`freeze.sh` to automate the pub sub process. This should make the build process simpler. It's not `make` but should be enough to fork with.

## General Things
There's still installing the `goali` binary to do. The todo issues should be easy now the TODO bot is added. Using `@todo`/`@body` in code comments. The command help needs expanding and deciding if environment variables are to be used. Then there's how to handle command verboseness for some "level" of detail. Then maybe i18n translations.

## Commands
### `knap`
The web servia. In slow planning.

### `mickey`
The GUI. This might end up being a mini-game.

### `snake`
The embedded python is almost there. It still needs to be given options to use the TUI, and have some more `snake/__init__.py` useful stuff added. Then the TUI launch will be worked through within the goali package. The `gin` package could then be expanded with custom TUI controls.

### `unicorn`
The unicode mangler. Needs the raw IO and mapping decided and done.


