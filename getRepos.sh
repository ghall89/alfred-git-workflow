#!/bin/bash

# Function to check if a directory is a Git repository
is_git_repo() {
  git -C "$1" rev-parse --is-inside-work-tree &> /dev/null
}

# Function to find Git repositories recursively and store their properties in an array
find_git_repos() {
  local dir="$1"
  local repos=()

  # Loop through each item in the directory
  for item in "$dir"/*; do
    if [ -d "$item" ]; then
      if is_git_repo "$item"; then
        # Extracting UID, file type, title, and file path
        uid=$(basename "$item")
        title=$(basename "$item")
        type="file"
        arg="$item"
        repos+=("{\"uid\":\"$uid\",\"type\":\"$type\",\"title\":\"$title\",\"subtitle\":\"$arg\",\"arg\":\"$arg\"},")
      else
        # Recursively call the function for subdirectories
        repos+=($(find_git_repos "$item"))
      fi
    fi
  done

  echo "${repos[@]}"
}

# Starting directory to search
start_dir=~/Developer

# Check if the starting directory is provided, otherwise use the current directory
if [ -z "$start_dir" ]; then
  start_dir="."
fi

# Call the function to find Git repositories
git_repos_json="{\"items\": ["
git_repos_json+=$(find_git_repos "$start_dir")
git_repos_json+="]}"

# Print out the JSON formatted array of Git repositories
echo "$git_repos_json"
