# install vim8 with ubuntu
```
sudo add-apt-repository ppa:jonathonf/vim
sudo apt update
sudo apt install vim
```

# install plugins

```
sudo apt-get install silversearcher-ag
sudo apt install vim-gnome

git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go
git clone https://github.com/scrooloose/nerdtree.git ~/.vim/pack/plugins/start/nerdtree
git clone https://github.com/kien/ctrlp.vim.git ~/.vim/pack/plugins/start/ctrlp.vim
git clone https://github.com/mileszs/ack.vim.git ~/.vim/pack/plugins/start/ack.vim
git clone https://github.com/ervandew/supertab.git ~/.vim/pack/plugins/start/supertab
git clone https://github.com/shawncplus/phpcomplete.vim.git ~/.vim/pack/plugins/start/phpcomplete.vim
git clone https://github.com/vim-php/phpctags.git ~/soft/phpctags
git clone https://github.com/Raimondi/delimitMate.git ~/.vim/pack/plugins/start/delimitMate
git clone https://github.com/rust-lang/rust.vim.git ~/.vim/pack/plugins/start/rust.vim
git clone https://github.com/racer-rust/vim-racer.git ~/.vim/pack/plugins/start/rust.vim-racer

:GoInstallBinaries
```
