# install vim8 with ubuntu
```
sudo add-apt-repository ppa:jonathonf/vim
sudo apt update
sudo apt install vim
```

# install ctags

```
# https://askubuntu.com/questions/796408/installing-and-using-universal-ctags-instead-of-exuberant-ctags
git clone https://github.com/universal-ctags/ctags.git
cd ctags
./autogen.sh 
./configure
make
sudo make install

# https://github.com/shawncplus/phpcomplete.vim/wiki/Getting-better-tags
ctags -R --fields=+aimlS --languages=php 
```

# install plugins

```
sudo apt-get install silversearcher-ag
sudo apt install vim-gnome

git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go
git clone https://github.com/scrooloose/nerdtree.git ~/.vim/pack/plugins/start/nerdtree
git clone https://github.com/vim-scripts/mru.vim.git ~/.vim/pack/plugins/start/mru.vim
git clone https://github.com/kien/ctrlp.vim.git ~/.vim/pack/plugins/start/ctrlp.vim
git clone https://github.com/mileszs/ack.vim.git ~/.vim/pack/plugins/start/ack.vim
git clone https://github.com/ervandew/supertab.git ~/.vim/pack/plugins/start/supertab
git clone https://github.com/shawncplus/phpcomplete.vim.git ~/.vim/pack/plugins/start/phpcomplete.vim
git clone https://github.com/vim-php/phpctags.git ~/soft/phpctags
git clone https://github.com/Raimondi/delimitMate.git ~/.vim/pack/plugins/start/delimitMate
git clone https://github.com/rust-lang/rust.vim.git ~/.vim/pack/plugins/start/rust.vim
git clone https://github.com/racer-rust/vim-racer.git ~/.vim/pack/plugins/start/rust.vim-racer
git clone https://github.com/sickill/vim-monokai.git ~/.vim/pack/plugins/start/monokai
git clone https://github.com/easymotion/vim-easymotion.git ~/.vim/pack/plugins/start/vim-easymotion

npm -g install instant-markdown-d
git clone https://github.com/suan/vim-instant-markdown.git ~/.vim/pack/plugins/start/vim-instant-markdown

:GoInstallBinaries
```
