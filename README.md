# Tree Sitter Prisma

[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

![](https://github.com/victorhqc/tree-sitter-prisma/workflows/Publish%20CI/badge.svg)

## Introduction

This is an **unofficial** Prisma language parsing. More information about the language and specs
can be found here:

- [vscode-prisma](https://github.com/prisma/vscode-prisma)
- [prisma2-schema-file](https://www.prisma.io/docs/concepts/components/prisma-schema)
- [prisma2-data-modeling](https://github.com/prisma/prisma2/blob/master/docs/data-modeling.md)

If you notice any bug or problem, please submit an issue or make a pull request. Any contribution
is welcomed and needed.

## Development

**Requirements:**

- Rust >= 1.36
- node >= 22

All the parsing logic is specified in `grammar.js` at the root level. To see if the changes made to
it are working, run the tests and compare the results.

```sh
npm test
```

A playground is also available to test the parser (though it needs Docker to run)

```sh
npm start
```

More information about how to write or use the tree parser can be found here:
[http://tree-sitter.github.io/tree-sitter/](http://tree-sitter.github.io/tree-sitter/)

Many parts of the code were scavenged from these repositories:

- [tree-sitter-javascript](https://github.com/tree-sitter/tree-sitter-javascript)
- [tree-sitter-typescript](https://github.com/tree-sitter/tree-sitter-typescript)

I'm grateful to the authors and contributors of those repositories, without them this parser would
be a lot times worse. Thank you for having such a good documentation and code.

## Neovim Usage

For neovim users, [`nvim-treesitter`](https://github.com/nvim-treesitter/nvim-treesitter)
is needed, please make sure it is installed.

A typical installation using [`vim-plug`]() would be like following

```lua
call plug#begin()

" ... Any other plugins

Plug 'nvim-treesitter/nvim-treesitter', {'do': ':TSUpdate'}

call plug#end()

lua <<EOF
require'nvim-treesitter.configs'.setup {
  ensure_installed = { "prisma" }, -- You can add other languages too
  sync_install = false,
  auto_install = true,
  highlight = {
    enable = true,
  },
}
EOF
```

And don't forget to run the following in nvim

```
:TSInstall
```
