<h1 align="center">
    Vimscript
</h1>

## Material

[Learn Vimscript the Hardway](http://learnvimscriptthehardway.stevelosh.com)
[Exercism vim track](https://exercism.io/my/tracks/vimscript)

## Reminders

### Variables

| Type | Example |
| ---- | ------- |
| scalar | `let height = 165` |
| list | `let interests = ['Cinema', 'Literature', 101]` |
| dictionary | ` let phone = {'cell': 123, 'work': '?'}` |

>Variable types, once assigned, are permanent and strictly enforced
at runtime. As we set the interests as a list there will be error now.

```vim
let interests = 'unknown'
```

>" Error: variable type mismatch

### Variable scopes and prefix

| Prefix | Meaning |
| ------ | ------- |
| `g:varname` | The variable is global |
| `s:varname` | The variable is local to the current script file |
| `w:varname` | The variable is local to the current editor window |
| `t:varname` | The variable is local to the current editor tab |
| `b:varname` | The variable is local to the current editor buffer |
| `l:varname` | The variable is local to the current function |
| `a:varname` | The variable is a parameter of the current function |
| `v:varname` | The variable is one that Vim predefines |

Pseudovariables

| Prefix | Meaning |
| ------ | ------- |
| `&varname` | A Vim option (local option if defined, otherwise global) |
| `&l:varname` | A local vim option |
| `&g:varname` | A global vim option |
| `@varname` | A vim register |
| `$varname` | An environment variable |

### Operators

| Operation | Operator Syntax |
| --------- | --------------- |
| Assignment | `let var=expr` |
| Numeric-add-and-assign | `let var+=expr` |
| Numeric-subtract-and-assign | `let var-=expr` |
| String-concatenate-and-assign | `let var.=expr` |
| Ternary operator | `bool?expr-if-true:expr-if-false` |
| Logical OR | `bool||bool` |
| Logical AND | `bool&&bool` |
| Numeric or string equality | `expr==expr` |
| String case insensitive eq | `expr==?expr` |
| String case sensitive eq | `expr==#expr` |
| Numeric or string inequality | `expr!=expr` |
| Numeric or string greater-than | `expr>expr` |
| Numeric or string gr-or-eq | `expr>=expr` |
| Numeric or string less than | `expr<expr` |
| Numeric or string l-or-eq | `expr<=expr` |
| Numeric addition | `num+num` |
| Numeric subtraction | `num-num` |
| String concatenation | `str.str` |
| Numeric Multiplication | `num*num` |
| Numeric division | `num/num` |
| Numeric modulus | `num%num` |
| Convert to number | `+num` |
| Numeric negation | `-num` |
| Logical NOT | `!bool` |
| Parenthetical precedence | `(expr)` |

## Usage

```vim
" File execution, you can run vimscripts from the vim
" command line.
" :source <vimscript.vim>
```
