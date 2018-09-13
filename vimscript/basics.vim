" Set assignment can only set to single literal values
set textwidth=100

" Let assignments can determine the values
let &textwidth = &textwidth + 10


" Local variable assignment with prefixes
let &l:number = 1


" Variable local to buffer
let b:hello = "world"


" Conditionals
if 0
    echom "if"
elseif "nope!"
    echom "elseif"
else
    echom "finally!"
endif


" Case Sensitivity is determined by the users settings
" the following will echo: vim is case sensitive
set noignorecase
if "foo" == "FOO"
    echom "vim is case insensitive"
elseif "foo" == "foo"
    echom "vim is case sensitive"
endif

" While this will echo: vim is case insensitive
set ignorecase
if "foo" == "FOO"
    echom "vim is case insensitive"
elseif "foo" == "foo"
    echom "vim is case sensitive"
endif

" Because of the above vim has two additional operators
" ==? Which is case-insensitive no matter what the user
" has set, and ==# which is case-sensitive.


" Stopped at: http://ay.stevelosh.com/chapters/23.html
