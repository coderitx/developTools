## vim 环境配置

```shell
curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```

## vimrc 配置
```shell
set number
set ignorecase
set showmatch
set encoding=utf-8
set showcmd

set tabstop=4
syntax enable

set background=dark

set ai

vnoremap y "+y

let g:ycm_complete_in_strings = 1
let g:ycm_show_diagnostics_ui = 0
let g:ycm_complete_in_comments=1
let g:ycm_confirm_extra_conf=0
let g:ycm_collect_identifiers_from_tags_files=1
let g:ycm_min_num_of_chars_for_completion=1
let g:ycm_cache_omnifunc=0
let g:ycm_seed_identifiers_with_syntax=1


let g:go_fmt_command = "goimports"
let g:go_debug=['shell-commands']
let g:go_metalinter_command = "golangci-lint"
let g:go_autodetect_gopath = 1
let g:go_list_type = "quickfix"

let g:go_version_warning = 1
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1

let g:godef_split=2

set nocompatible
filetype off

call plug#begin()
Plug 'neoclide/coc.nvim', {'branch': 'master'}
Plug 'rust-lang/rust.vim'
Plug 'prabirshrestha/vim-lsp'
Plug 'scrooloose/nerdtree'
Plug 'dense-analysis/ale'
Plug 'vim-airline/vim-airline'
let g:ale_linters = {
  \ 'rust': ['cargo', 'clippy']
  \ }

let g:ale_fixers = {
  \ 'rust': ['cargo', 'rustfmt']
  \ }

call plug#end()

filetype plugin indent on

let mapleader = ","

let g:rustfmt_autosave = 1 " 保存时自动格式化
let g:rustfmt_command = "rustfmt" " 自定义格式化命令

let g:ale_rust_cargo_use_clippy = 1

let g:airline#extensions#ale#enabled = 1

autocmd vimenter * NERDTree
vnoremap <leader>ft :RustFmtRange<CR>
nnoremap <leader>ft :RustFmt<CR>
nnoremap <M-r> :RustRun<CR>
nnoremap <M-t> :RustTest<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
```

## 安装插件
vim ~/.vimrc

命令行模式下执行:
:PlugInstall  # 安装插件
## vim 环境配置


```shell
curl -fLo ~/.vim/autoload/plug.vim --create-dirs https:raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```


## vimrc 配置
```shell
set number
set ignorecase
set showmatch
set encoding=utf-8
set showcmd


set tabstop=4
syntax enable


set background=dark


set ai


vnoremap y "+y


let g:ycm_complete_in_strings = 1
let g:ycm_show_diagnostics_ui = 0
let g:ycm_complete_in_comments=1
let g:ycm_confirm_extra_conf=0
let g:ycm_collect_identifiers_from_tags_files=1
let g:ycm_min_num_of_chars_for_completion=1
let g:ycm_cache_omnifunc=0
let g:ycm_seed_identifiers_with_syntax=1




let g:go_fmt_command = "goimports"
let g:go_debug=['shell-commands']
let g:go_metalinter_command = "golangci-lint"
let g:go_autodetect_gopath = 1
let g:go_list_type = "quickfix"


let g:go_version_warning = 1
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1


let g:godef_split=2


set nocompatible
filetype off


call plug#begin()
Plug 'neoclide/coc.nvim', {'branch': 'master'}
Plug 'rust-lang/rust.vim'
Plug 'prabirshrestha/vim-lsp'
Plug 'scrooloose/nerdtree'
Plug 'dense-analysis/ale'
Plug 'vim-airline/vim-airline'
let g:ale_linters = {
  \ 'rust': ['cargo', 'clippy']
  \ }


let g:ale_fixers = {
  \ 'rust': ['cargo', 'rustfmt']
  \ }


call plug#end()


filetype plugin indent on


let mapleader = ","


let g:rustfmt_autosave = 1 " 保存时自动格式化
let g:rustfmt_command = "rustfmt" " 自定义格式化命令


let g:ale_rust_cargo_use_clippy = 1


let g:airline#extensions#ale#enabled = 1


autocmd vimenter * NERDTree
vnoremap <leader>ft :RustFmtRange<CR>
nnoremap <leader>ft :RustFmt<CR>
nnoremap <M-r> :RustRun<CR>
nnoremap <M-t> :RustTest<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
```


## 安装插件
vim ~/.vimrc


命令行模式下执行:
:PlugInstall  # 安装插件
:plugStatus ## vim 环境配置

```shell
curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```

## vimrc 配置
```shell
set number
set ignorecase
set showmatch
set encoding=utf-8
set showcmd

set tabstop=4
syntax enable

set background=dark

set ai

vnoremap y "+y

let g:ycm_complete_in_strings = 1
let g:ycm_show_diagnostics_ui = 0
let g:ycm_complete_in_comments=1
let g:ycm_confirm_extra_conf=0
let g:ycm_collect_identifiers_from_tags_files=1
let g:ycm_min_num_of_chars_for_completion=1
let g:ycm_cache_omnifunc=0
let g:ycm_seed_identifiers_with_syntax=1


let g:go_fmt_command = "goimports"
let g:go_debug=['shell-commands']
let g:go_metalinter_command = "golangci-lint"
let g:go_autodetect_gopath = 1
let g:go_list_type = "quickfix"

let g:go_version_warning = 1
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1

let g:godef_split=2

set nocompatible
filetype off

call plug#begin()
Plug 'neoclide/coc.nvim', {'branch': 'master'}
Plug 'rust-lang/rust.vim'
Plug 'prabirshrestha/vim-lsp'
Plug 'scrooloose/nerdtree'
Plug 'dense-analysis/ale'
Plug 'vim-airline/vim-airline'
let g:ale_linters = {
  \ 'rust': ['cargo', 'clippy']
  \ }

let g:ale_fixers = {
  \ 'rust': ['cargo', 'rustfmt']
  \ }

call plug#end()

filetype plugin indent on

let mapleader = ","

let g:rustfmt_autosave = 1 " 保存时自动格式化
let g:rustfmt_command = "rustfmt" " 自定义格式化命令

let g:ale_rust_cargo_use_clippy = 1

let g:airline#extensions#ale#enabled = 1

autocmd vimenter * NERDTree
vnoremap <leader>ft :RustFmtRange<CR>
nnoremap <leader>ft :RustFmt<CR>
nnoremap <M-r> :RustRun<CR>
nnoremap <M-t> :RustTest<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
```

## 安装插件
vim ~/.vimrc

命令行模式下执行:
:PlugInstall  # 安装插件
:PlugStatus   # 查看插件
light g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
easy g:go_highlight_operators = 1
easy g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1


easy g:godef_split=2


Set Nocompatible
filetype off


call plug#begin()
Plug 'neoclide/coc.nvim', {'branch': 'master'}
Plug 'rust-lang/rust.vim'
Plug 'prabirshrestha/vim-lsp'
Plug 'scrooloose/nerdtree'
Plug 'dense-analysis/ale'
Plug 'vim-airline/vim-airline'
let g:ale_linters = {
  \'rust': ['cargo', 'clippy']
  \ }


let g:ale_fixers = {
  \ 'rust': ['cargo', 'rustfmt']
  \ }


call plug#end()


filetype plugin indent on


easy mapleader = ","


Let G:rustfmt_autosave = 1" 保存时自动格式化
let g:rustfmt_command = "rustfmt" " 自定义格式化命令


Let g:ale_rust_cargo_use_clippy = 1


easy g:airline#extensions#ale#enabled = 1


autocmd vimenter * NERDTree
vnoremap <leader>ft :RustFmtRange<CR>
nnoremap <leader>ft :RustFmt<CR>
nnoremap <M-r> :RustRun<CR>
nnoremap <M-t> :RustTest<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
```


## 安装插件
vim ~/.vimrc


命令行模式下执行:
:P lugInstall #安装插件
:P lugStatus #查看插件
:P lugUpdate let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1

let g:godef_split=2

set nocompatible
filetype off

call plug#begin()
Plug 'neoclide/coc.nvim', {'branch': 'master'}
Plug 'rust-lang/rust.vim'
Plug 'prabirshrestha/vim-lsp'
Plug 'scrooloose/nerdtree'
Plug 'dense-analysis/ale'
Plug 'vim-airline/vim-airline'
let g:ale_linters = {
  \ 'rust': ['cargo', 'clippy']
  \ }

let g:ale_fixers = {
  \ 'rust': ['cargo', 'rustfmt']
  \ }

call plug#end()

filetype plugin indent on

let mapleader = ","

let g:rustfmt_autosave = 1 " 保存时自动格式化
let g:rustfmt_command = "rustfmt" " 自定义格式化命令

let g:ale_rust_cargo_use_clippy = 1

let g:airline#extensions#ale#enabled = 1

autocmd vimenter * NERDTree
vnoremap <leader>ft :RustFmtRange<CR>
nnoremap <leader>ft :RustFmt<CR>
nnoremap <M-r> :RustRun<CR>
nnoremap <M-t> :RustTest<CR>
nnoremap <C-t> :NERDTreeToggle<CR>
```

## 安装插件
vim ~/.vimrc

命令行模式下执行:
:PlugInstall  # 安装插件
:PlugStatus   # 查看插件
:PlugUpdate   # 更新插件
