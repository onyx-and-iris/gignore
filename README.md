![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![macOS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=macos&logoColor=F0F0F0)


# Gignore - Generate .gitinore files

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
Usage of gignore:
  gignore [flags] <template>

Flags:
  -dir string
        directory containing .gitignore templates (default "gitignoreio")
  -l int
        log level (shorthand) (default 3)
  -list
        list available templates
  -loglevel int
        log level (default 3)
  -ls
        list available templates (shorthand)

Example:
  gignore go
```

## Custom Templates

It's possible to add your own custom templates, simply create a directory in `internal/registry/templates`. You'll need to [reinstall](https://github.com/onyx-and-iris/gignore?tab=readme-ov-file#install) the project before you can load the new templates.

Then pass the dir name as a flag, for example:

```bash
gignore -dir=custom go
```

You may set an environment variable `GIGNORE_TEMPLATE_DIR` to avoid passing the `-dir` flag each time.

If a template is requested but not found in the custom directory then the gitignoreio registry will act as a fallback.

## Special Thanks

[gitignore.io][gitignoreio] For providing such a useful .gitignore service

[cuonglm][cuonglm] For writing the [gogi][gogi] client library for gitignore.io


[task]: https://taskfile.dev/
[gitignoreio]: https://www.toptal.com/developers/gitignore
[cuonglm]: https://github.com/cuonglm
[gogi]: https://github.com/cuonglm/gogi
[ignore]: https://github.com/neptship/ignore
