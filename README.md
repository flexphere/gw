# GW

gw is a simple wrapper around git worktree command.
it adds a custom configuration per repository to specify the worktree path and commands to run on worktree creation.


## Getting Started

### Installation
```sh
go install github.com/flexphere/gw@0.0.2
```

### Usage

#### Setup current repository
```sh
gw init \
--worktree-path ~/.worktree \         # setup current repository to create its worktrees under "~/.worktree"
--cmd "echo hello" --cmd "echo world" # run `echo hello` and `echo world` on worktree creation
```

#### Creating new worktree
```sh
gw new name-of-worktree
```

#### Editing the config
```sh
gw edit # opens the config file in the editor
```

#### Passthrough
every other command then above is redirected to `git worktree` command
```sh
gw list # executes `git worktree list`
gw prune # executes `git worktree prune`
...
```
