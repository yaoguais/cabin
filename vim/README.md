## Install

```
first I prefer zsh
https://github.com/robbyrussell/oh-my-zsh

install zsh
# yum -y install zsh
# sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"

second install vim
cause by YouCompleteMe, use Vim 7.4.1578 with Python 2 or Python 3 support.
build vim from source: https://github.com/Valloric/YouCompleteMe/wiki/Building-Vim-from-source
# yum install -y ruby ruby-devel lua lua-devel luajit \
    luajit-devel ctags git python python-devel \
    python3 python3-devel tcl-devel \
    perl perl-devel perl-ExtUtils-ParseXS \
    perl-ExtUtils-XSpp perl-ExtUtils-CBuilder \
    perl-ExtUtils-Embed
# cd /usr/local
# git clone https://github.com/vim/vim.git
# cd vim
# ./configure --with-features=huge \
	--enable-multibyte \
	 --enable-rubyinterp=yes \
	--enable-pythoninterp=yes \
	--with-python-config-dir=/usr/bin \
	--enable-python3interp=yes \
	--with-python3-config-dir=/usr/lib/python3.5/config \
	--enable-perlinterp=yes \
	--enable-luainterp=yes \
	--enable-gui=gtk2 \
	--enable-cscope \
	--prefix=/usr/local
# make && make install

set PATH, VIMRUNTIME
# VIMRUNTIME=/usr/local/vim/runtime
# PATH=$PATH:/usr/local/vim/src
# vim --version

install spf13-vim3
# curl https://j.mp/spf13-vim3 -L > spf13-vim.sh && sh spf13-vim.sh

install vim-go, https://github.com/fatih/vim-go

edit configuration files

# cat .vimrc.local
set clipboard=unnamed
set clipboard=unnamedplus
" enable right mouse menu, but disabled NERDTree ???
" set mouse-=a
set nocompatible
filetype off
filetype plugin indent on
syntax on
au BufRead,BufNewFile *.go set filetype=go
colorscheme molokai

# cat .vimrc.bundles.local
Bundle 'fatih/vim-go'
Bundle 'Valloric/YouCompleteMe'

#  cat .vim.before.fork
let g:spf13_bundle_groups=['general', 'programming', 'misc', 'youcompleteme', 'go']

```

