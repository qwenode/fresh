# Fresh


Fresh is a command line tool that builds and (re)starts your web application everytime you save a Go or template file.

If the web framework you are using supports the Fresh runner, it will show build errors on your browser.


## Installation

    go install github.com/qwenode/fresh@latest

## Usage

    cd /path/to/myapp

Start fresh:

    fresh config.ini

Fresh will watch for file events, and every time you create/modify/delete a file it will build and restart the application.
If `go build` returns an error, it will log it in the tmp folder.
