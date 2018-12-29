# Toggl Export

[![Go Report Card](https://goreportcard.com/badge/github.com/nylo-andry/toggl-export)](https://goreportcard.com/report/github.com/nylo-andry/toggl-export)

> A simple command-line tool that exports your toggl timesheet to CSV

## Getting started

1. Create a `config.toml` file by copying the `example.config.toml`.

    - Fill in three values (`api_token`, `workspace_id`, `user_name`).
    - You can find your API token in your [Profile settings](https://toggl.com/app/profile)
    - If you do not know your `workspace_id`, you can find it by running the `workspace` command-line executable. You will still need to provide your API Token in your configuration file before you can find it.

2. Run the `export` executable and provide the timeframe you want to export (e.g. `export -start=2018-12-17 -end=2018-12-21`)

# Todos

- [x] Better structure
- [x] Remove resty. I do not really need it.
- [x] Tests
- [x] Create a cmd that will tell the person what is their workspace id
- [ ] Create Makefile
- [x] Better doc
- [ ] Create first release with documentation
