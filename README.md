# Gignore - Generate .gitinore files

## Install

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

It's possible to add your own custom templates, simply create a directory in `internal/registry/templates`. You'll need to rebuild the project before you can load the new templates.

Then pass the dir name as a flag, for example:

```bash
gignore -dir=custom go
```

You may set an environment variable `GIGNORE_TEMPLATE_DIR` to avoid passing the `-dir` flag each time.

If a template is requested but not found in the custom directory then the gitignoreio registry will act as a fallback.

## Special Thanks

[gitignore.io][gitignoreio] For providing such a useful .gitignore service

[mh-cbon][mh-cbon] For writing the [gigo][gigo] client library for gitignore.io


[task]: https://taskfile.dev/
[gitignoreio]: https://www.toptal.com/developers/gitignore
[mh-cbon]: https://github.com/mh-cbon
[gigo]: https://github.com/mh-cbon/gigo
[ignore]: https://github.com/neptship/ignore
