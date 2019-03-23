"
" This function takes a year and returns 1 if it's a leap year
" and 0 otherwise.
"
function! LeapYear(number) abort
    if (a:number%4)==0
        if (a:number%100)==0
            if (a:number%400)==0
                return 1
            endif
            return 0
        endif
        return 1
    endif
    return 0
endfunction
