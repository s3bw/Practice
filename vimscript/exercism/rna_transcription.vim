"
" This function takes a DNA strand and returns its RNA complement.
"
"   G -> C
"   C -> G
"   T -> A
"   A -> U
"
" If the input is invalid, return an empty string.
"
" Example:
"
"   :echo ToRna('ACGTGGTCTTAA')
"   UGCACCAGAAUU
"

let s:rna = {
\   'G':'C',
\   'C': 'G',
\   'T': 'A',
\   'A': 'U',
\}

function! ToRna(strand) abort
    let a:complement = ''
    for s:item in split(a:strand, '\zs')
        if s:item !~? '[GCTA]'
            return ''
        endif
        let a:complement.=s:rna[s:item]
    endfor
    return a:complement
endfunction


" echom ToRna('CTA')
