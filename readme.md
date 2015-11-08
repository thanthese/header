# What's it do?

Before

```md
# First order headline

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis
nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

## Second order headline

# First order headline
```

After

```md
# First order headline =========================================================

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis
nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

## Second order headline -------------------------------------------------------

# First order headline =========================================================
```

# Installation

Assuming you have [go](https://golang.org/) and [git](https://git-scm.com/) installed:

```sh
go get http://github.com/thanthese/header
```

# Usage

Operates on `STDIN`. Intended use is to pipe the current line in your text editor through the tool and back into the active buffer. With vim that would be something like `!!header` from normal mode. A helpful mapping might be:

```viml
nmap <c-h> :.!header<cr>
```

To apply to every headline in the current document use `:g/^#/.!header`.

Use `-h` to see full options list:

    Usage of header:
      -d depth
            (optional) Header depth 0-2, or -1 to derive from input. (default -1)
      -p plain
            (optional) Don't include banner -- "plain". (default with banners)
      -w width
            (optional) Header width. (default 75)

# Rationale

Markdown's `#` for headlines is wonderfully easy to type and easy for programs to parse. However, contrary to markdown's original intentions, it can make headings difficult to spot in long documents. The underlining syntax solves the visibility problem, but at the expense of both other benefits. This system is the best of all worlds -- a system that's easy to type, easy to parse, and easy to see in a plain text document.

Try this neat trick in vim: `:g/^#` to see an outline of the document. Use `:grep` or [fzf.vim](https://github.com/junegunn/fzf.vim) to making viewing the outline even easier.

# Why are there only two headline levels?

If you need more than two levels, you're doing it wrong. Create more, smaller files. If you really, really need the extra depth, consider using a tool that supports that complexity -- something with folding capabilities.
