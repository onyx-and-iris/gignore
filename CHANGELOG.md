# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

# [0.6.0] - 2025-05-01

### Added

-   list now accepts pattern arguments, for example `gignore list python` to filter out python templates.
-   `--config`  flag added to the root command for setting the config path (if required).
-   tests for list/create commands added

### Changed

-   Viper now used for configuration management. This changes the following:
    -   env var `GIGNORE_TEMPLATE_ROOT` is now `GIGNORE_ROOT`.
    -   configuration may be loaded from a config.yaml (default path $HOME/.config/gignore/config.yaml)

# [0.5.0] - 2025-04-13

### Changed

-   CLI component migrated to Cobra. This introduces the following changes:
    -   `list` is now a subcommand.
    -   `create` has been added as a subcommand, use it to create a new .gitignore file.
-   Env var `GIGNORE_TEMPLATE_DIR` changed to `GIGNORE_TEMPLATE_ROOT`
-   Env var `GIGNORE_LOGLEVEL` may now be used to set the logging level.

# [0.4.0] - 2025-04-05

### Changed

-   `-loglevel` flag is now of type string. It accepts any one of trace, debug, info, warn, error, fatal or panic.
    -   It defaults to warn.

# [0.3.0] - 2025-14-03

### Added

-   CLI may now accept multiple template names, example `gignore go python`. One will be appended after the other.

### Changed

-   Filewriter now opens file in append mode.

# [0.2.0] - 2025-10-03

### Fixed

-   Template .gitignore are now written concurrently.

# [0.1.0] - 2025-09-03

### Added

-   Initial release.