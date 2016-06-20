
set autoindent
set dir=.,~/tmp/vim
set ff=unix
setlocal ff=unix
set guifont=Bitstream\ Vera\ Sans\ Mono\ 10
set hlsearch
set noic
set linebreak
set matchtime=3
set number
set shellslash
set showmatch
set shiftwidth=4
set tabstop=4
set expandtab
set tags=/vobs/test/tags
set wildmode=longest,list
set wrap
set nowrapscan

ab come ! ct co -nc %
ab mkme ! mymk
ab mkk ! mymk

au BufRead,BufNewFile *.py set filetype=py
au BufEnter *.py set ai sw=4 ts=4 sta et fo=croql
au BufEnter *.cc,*.cpp,*.cxx,*.c,*.h set nu ai sw=4 ts=4 expandtab sta et fo=croql



