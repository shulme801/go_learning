# ENVIRONMENT VARIABLES

# simple prompt
export PS1="\W ==> "

export EDITOR=vim

# Make versions of tools installed with Homebrew take precedence over other versions
export PATH="/usr/local/bin:/usr/local/sbin:~/bin:$PATH"
export EDITOR="code -w"
#export GOPATH="/Users/shulme801/Dropbox/code2019/src/gopl.io"
export GOPATH="/Users/shulme801/Dropbox/code2019/src"

# Aliases
alias la="ls -la"
alias ll="ls -ll"
alias rm="rm -i"
alias mv="mv -i"
alias cp="cp -i"
alias ..="cd .."
alias ...="cd ../.."
alias cd..="cd .."
alias python="python3"
alias gocode="cd /Users/shulme801/Dropbox/code2019/src"
alias gonative="cd /Users/shulme801/Dropbox/code2019/src/docker/shulme801/cloudnatived"
alias ld="ls -ld"
