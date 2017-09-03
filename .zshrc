#
# Executes commands at the start of an interactive session.
#
# Authors:
#   Sorin Ionescu <sorin.ionescu@gmail.com>
#

# Source Prezto.
if [[ -s "${ZDOTDIR:-$HOME}/.zprezto/init.zsh" ]]; then
  source "${ZDOTDIR:-$HOME}/.zprezto/init.zsh"
fi
if [ -d "/home/bhavishya/somewhere/arcanist/bin/" ] ; then
    export PATH="/home/bhavishya/somewhere/arcanist/bin/:$PATH"
fi
# This way the completion script does not have to parse Bazel's options
# repeatedly.  The directory in cache-path must be created manually.
zstyle ':completion:*' use-cache on
zstyle ':completion:*' cache-path ~/.zsh/cache
# Customize to your needs...
alias bazel=/home/bhavishya/Documents/bazel/output/bazel
alias python=/usr/bin/python3.6
alias g='git'
alias gi='git init'
alias ga='git add'
alias gs='git status'
alias gaa='git add --all'
alias gbn='git checkout -n'
alias gc='git commit'
alias gco='git checkout'
alias gp='git push'
alias gra='git remote add'
alias grv='git remote -v'

#For file viewing
alias O='xdg-open'
alias ..='cd ..'
#for xscreen-saver
alias L='xscreensaver-command --lock'
alias D='xscreensaver-command --demo'
source ~/.zpyi/zpyi.zsh 
