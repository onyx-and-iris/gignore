![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![macOS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=macos&logoColor=F0F0F0)


# Gignore - Generate .gitignore files

## Install

With Go tools:

```bash
go generate ./...
go install ./cmd/gignore
```

With [Task][task]:

```bash
task install
```

## Usage

```bash
Usage:
  gignore [flags]
  gignore [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create a new .gitignore file
  help        Help about any command
  list        List all .gitignore files in the current directory

Flags:
  -h, --help              help for gignore
  -l, --loglevel string   Log level (trace, debug, info, warn, error, fatal, panic) (default "info")
  -r, --root string       Root directory to search for .gitignore files (default "gitignoreio")

Use "gignore [command] --help" for more information about a command.
```

For example:

```bash
gignore create go
```

## Custom Templates

It's possible to add your own custom templates, simply create a directory in `internal/registry/templates`. You'll need to [reinstall](https://github.com/onyx-and-iris/gignore?tab=readme-ov-file#install) the project before you can load the new templates.

Then pass the dir name as a flag, for example:

```bash
gignore -root=custom create go
```

You may set an environment variable `GIGNORE_TEMPLATE_ROOT` to avoid passing the `-root` flag each time.

If a template is requested but not found in the custom directory then the gitignoreio registry will act as a fallback.

## Logging

The `-loglevel` flag allows you to control the verbosity of the application's logging output. 

Acceptable values for this flag are:

- `trace`
- `debug`
- `info`
- `warn`
- `error`
- `fatal`
- `panic`

For example, to set the log level to `debug`, you can use:

```bash
gignore -loglevel=debug create python
```

The default log level is `warn` if the flag is not specified.

## Special Thanks

[gitignore.io][gitignoreio] For providing such a useful .gitignore service

[cuonglm][cuonglm] For writing the [gogi][gogi] client library for gitignore.io


[task]: https://taskfile.dev/
[gitignoreio]: https://www.toptal.com/developers/gitignore
[cuonglm]: https://github.com/cuonglm
[gogi]: https://github.com/cuonglm/gogi
[ignore]: https://github.com/neptship/ignore
