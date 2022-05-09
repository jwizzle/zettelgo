# Zettelgo

A go project that helps keeping notes on your unix system, in a zettelkasten-like fashion.
I mainly created it as a back-end for a [neovim plug-in](https://github.com/jwizzle/nd.nvim/). But it should be usable stand-alone.

The zettels package serves as the brains of the zettelkasten functionality. While cobra-based commands wrap it and serve it as a functional package.
The main goal writing this, was to call the commands from lua and capture json output. All functionality to let nd.nvim work correctly is currently present, and probably won't change. Yet it's still under development.

## Features

* List notes
* Show contents
* Start editting -> TODO
* Keep links in sync between notes -> TODO
* Create new notes by a template
* Filter results by tag, title, path, etc. (with json)
* Output as json

## TODO

* Remove note parsing from commands where it's not necessary
* Daemonize, switch to something better than capturing shell output for interfacing between other applications
