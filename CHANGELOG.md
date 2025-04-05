# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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