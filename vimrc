" Indent next line to match the current line's indentation.
set autoindent
" Dir list to store vim's temp files.
set dir=.,~/tmp/vim
" line ending as in Unix (line feed only, no carriage return from DOS)
set ff=unix
setlocal ff=unix

" Nice programmer font
set guifont=Bitstream\ Vera\ Sans\ Mono\ 10
" Highlight found strings
set hlsearch
" Search distinguishes between 'a' and 'A'
set noignorecase
set linebreak
set matchtime=3
" Number lines
set number
set shellslash
" Matching () {} [] shown
set showmatch

" Tabs of 4
set shiftwidth=4
set tabstop=4
set expandtab
set smarttab

set tags=/path/to/tag/file/tags
" In ':e' tab expands to list of all files that match.
set wildmode=longest,list
set wrap
" Stop scan at end of buffer (and at beginning in reverse search)
set nowrapscan

" Some macros
ab come ! ct co -nc %
ab mkme ! mymk
ab mkk ! mymk
ab p4e ! p4 edit %
ab pe ! p4 edit %

" Buffer specific commands for: python, C/C++, makefiles
au BufRead,BufNewFile *.py set filetype=py
au BufEnter *.py set ai sw=4 ts=4 sta et fo=croql

au BufEnter,BufNewFile *.cc,*.cpp,*.cxx,*.c,*.h set nu ai sw=4 ts=4 expandtab sta et fo=croql

" Makefiles should keep tabs
au BufRead,BufNewFile,BufEnter make*,Make*  set noexpandtab

command Bash ConqueTerm bash
command Rc ConqueTerm rc

