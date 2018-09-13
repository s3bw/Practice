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


" Functions must be defined with a capital letter.
function Meow()
    echom "Meow!"
endfunction

" Calling
call Meow()

" Returning
function GetMeow()
    return "Meow String!"
endfunction

echom GetMeow()

" Implicit returning, if nothing is returned it returns
" a 0
function TextwidthIsTooWide()
    if &l:textwidth ># 80
        return 1
    endif
endfunction

set textwidth=80
if TextwidthIsTooWide()
    echom "WARNING: Wide text!"
endif

" Arguments
function DisplayName(name)
    echom "Hello! My name is:"
    echom a:name
endfunction

" We provide the a: to represent the variable scope.
" otherwise vim won't be able to find the variable name.

" Varargs
function Varg(...)
    echom a:0
    echom a:1
    echom a:000
endfunction

call Varg("a", "b")

" The ellipsis tells vim that this function can take any
" number of arguments. When we echo a:0 this will display 2
" since this argument tells us of how many args were passed.
" a:1 displays the first extra argument. etc.
"
" Args in the a: scope can not be reassigned.

" Numbers can be resolved from hex notation using echom.
echom 0xff


" String concatenation
" Vim has a concatenation operator, using .
echom "Hello, " . "world"


" Vim String Functions
" split - splits by a separator
" split('one, two, three', ',')
"
" join - joins by a separator
" join(['foo", 'bar'], '...')
"
" tolower - lowercases
" tolower('Foo')
"
" toupper - uppercases
" toupper('Foo')
"
" Chapter 38
