# ipcmsg

**WIP:**
This is work in progress, do not use for anything serious.

The `parsey` package is meant to help supporting line-based configuration files similar to OpenBSD's look and feel.

It differs from `parse.y` in that it doesn't generate a parser at compile time,
but lets you configure the lexer and grammar dynamically.

For example of use,
see the [example program](https://github.com/poolpOrg/go-parsey/blob/main/example/example.go)