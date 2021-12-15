# TLDR Page Creator

TLDR Page Creator is a program designed to help users make TLDR pages, while avoiding syntax errors from TLDR-style markdown.

This doesn't substitute learning the syntax and you should definitely read the guides wrote by TLDR contributors and maintainers:

- https://github.com/tldr-pages/tldr/blob/main/contributing-guides/style-guide.md
- https://github.com/tldr-pages/tldr/blob/main/contributing-guides/git-terminal.md

However some of the following syntax is complete in this program:

- Title
- Punctuation (backticks, periods/full stops, dashes, greater-than symbol and colons)
- Preset for adding "More information:"
- Formatting of the page

Limitations:

- Lack of flags/arguments which can be used (planned to be added)
- "Token syntax"
- User specific errors (Open issues if required!)

## Usage

```text
Usage: tldr-page [options]

Options:
    -2: Use 2 lines of description in the page
```

## Installation

Requirements:

- A Go compiler

To Run:

`go run tldr-page.go`

To Install:

```shell
git clone https://github.com/CleanMachine1/tldr-page-creator.git
cd tldr-page-creator
go build tldr-page.go && sudo mv tldr-page /usr/local/bin/
```

Then after installation, to execute:

`tldr-page`

## Uninstallation

If you chose to install TLDR Page Creator and wish to uninstall:

`sudo rm /usr/local/bin/tldr-page`

## License

MIT
