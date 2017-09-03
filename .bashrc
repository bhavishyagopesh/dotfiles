#
# ~/.bashrc
#

# If not running interactively, don't do anything
[[ $- != *i* ]] && return

export LFS=/mnt/lfs

alias ls='ls --color=auto'
# Git branch in prompt.
parse_git_branch() {
    git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/ (\1)/'
}

export PS1="\[\e[0;36m[\u\e[m\]@\h \[\e[0;33m\W\e[m\]\$(parse_git_branch)\[\033[00m\]]$ "
export DISPLAY=:0.0
PATH=$PATH:/usr/local/bin:/home/bhavishya/somewhere/arcanist/bin/
xhost + >/dev/null
DEFAULT="[37;40m"
PINK="[35;40m"
GREEN="[32;40m"
ORANGE="[33;40m"

#hg_dirty() {
    #hg status --no-color 2> /dev/null \
    #| awk '$1 == "?" { unknown = 1 }
           #$1 != "?" { changed = 1 }
           #END {
             #if (changed) printf "!"
             #else if (unknown) printf "?"
           #}'
#}

#hg_branch() {
    #hg branch 2> /dev/null | \
        #awk '{ printf "\033[37;0m on \033[35;40m" $1 }'
    #hg bookmarks 2> /dev/null | \
        #awk '/\*/ { printf "\033[37;0m at \033[33;40m" $2 }'
#}

#export PS1='\n\e${PINK}\u \
#\e${DEFAULT}at \e${ORANGE}\h \
#\e${DEFAULT}in \e${GREEN}\w\
#$(hg_branch)\e${GREEN}$(hg_dirty)\
#\e${DEFAULT}\n$ '

