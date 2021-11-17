# TLDR Page Creator

TLDR Page Creator is a program designed to help users make TLDR pages, while avoiding syntax errors from TLDR-style markdown.

**It's currently work-in-progress**

This doesn't substitute learning the syntax and you should definitely read the guides wrote by TLDR contributors and maintainers:

- https://github.com/tldr-pages/tldr/blob/main/contributing-guides/style-guide.md
- https://github.com/tldr-pages/tldr/blob/main/contributing-guides/git-terminal.md

However some of the following syntax is complete in this program:

- Title
- Punctuation (backticks, periods/full stops, dashes, greater-than symbol and colons)
- Preset for adding "More information:"
- Formatting of the page

Limitations:

- 1 line of description for the main command (will be fixed later)
- Checking for user-given punctuation (planned to be added)
- Lack of flags/arguments which can be used (planned to be added)
- "Token syntax"
- User specific errors (Open issues if required!)

## Installation/Usage

Requirements:

- A Go compiler

To Run:

`go run tldr-page.go`

To Install:

`go build tldr-page.go && sudo mv tldr-page /usr/local/bin/`

Then after installation:

`tldr-page`

## Uninstallation

If you chose to install TLDR Page Creator and wish to uninstall:

`sudo rm /usr/bin/local/tldr-page`

## License

MIT
