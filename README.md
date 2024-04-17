# Git Repos

## About

An Alfred workflow for listing all git repos in a directory (including sub-directories).

Inspired by the [git-repos Raycast extension](https://github.com/raycast/extensions/tree/main/extensions/git-repos)

## Installation

Download the latest version from [releases](https://github.com/ghall89/alfred-git-workflow/releases/). When downloaded, unzip then double-click the file `Git Repos.alfredworkflow`, and follow the prompts to install and configure the Workflow to your preferences.

## Usage

When Alfred is active, type `repo`, and you'll get a list of all the repos inside the directory you chose during the installation process. You can continue to type to search for a specific repo.

When a result is highlighted:

- Pressing the right arrow key will bring up the standard Alfred folder actions.

- Pressing return/enter will trigger the default "open in" action you selected during installation.

## Go

This workflow's custom logic is handled by a simple utility I wrote in Go. The utility is compiled for performance and ease of use - you don't need Go installed on your machine to use it. However, I've included the source code for the utility in `src/getRepos` for anybody who is curious.
