
set autoindent
set dir=.,~/tmp/vim
set expandtab
set ff=unix
setlocal ff=unix
set guifont=Bitstream\ Vera\ Sans\ Mono\ 10
set hlsearch
set noignorecase
set linebreak
set matchtime=3
set number
set shellslash
set shiftwidth=4
set showmatch
set tabstop=4
set smarttab
set tags=/vobs/test/tags
set wildmode=longest,list
set wrap
set nowrapscan

ab come ! ct co -nc %
ab mkme ! mymk
ab mkk ! mymk
ab p4e ! p4 edit %
ab pe ! p4 edit %

au BufRead,BufNewFile *.py set filetype=py
au BufEnter *.py set ai sw=4 ts=4 sta et fo=croql
au BufEnter *.cc,*.cpp,*.cxx,*.c,*.h set nu ai sw=4 ts=4 expandtab sta et fo=croql

command Bash ConqueTerm bash
command Rc ConqueTerm rc


