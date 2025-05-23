## [0.1.9](https://github.com/paralus/relay/compare/v0.1.8...v0.1.9) (2025-05-23)

### Added
- restricted the port-forwording for read-only user ([1053483](https://github.com/paralus/relay/commit/1053483867883b10a0f4c47a4d4bfea12e7e109f))

## [0.1.8](https://github.com/paralus/relay/compare/v0.1.7...v0.1.8) (2024-06-14)

### Added
- updated go version to fix reported vulnerabilities ([91e0dae](https://github.com/paralus/relay/commit/91e0daea955ba8b4a73e3e99e68b0b92c7cef799))
- upgraded xnet library version to v0.23 to fix vulnerabilities ([75b6977](https://github.com/paralus/relay/commit/75b6977ea9dcbfa3df02b3afc6ca285f679f4c29))

## [0.1.7](https://github.com/paralus/relay/compare/v0.1.6...v0.1.7) (2024-02-28)

### Features

* update cluster health based on the dialin connection ([#108](https://github.com/paralus/relay/issues/108)) ([47e3f14](https://github.com/paralus/relay/commit/47e3f141d07109942756960069891c207bc2d52e))

## [0.1.6](https://github.com/paralus/relay/compare/v0.1.5...v0.1.6) (2023-09-25)

### Features

* changes to add project information with cluster lookup ([#225](https://github.com/paralus/paralus/issues/225)) ([e037e71](https://github.com/paralus/relay/commit/e037e718521dff8159cbef36ad2487127f549bb9))


### Bug Fixes

* use make check to clean up code formatting problems ([#78](https://github.com/paralus/relay/issues/78)) ([3c488bb](https://github.com/paralus/relay/commit/3c488bb8acf2e159ab536d838158150cd7e612ea))

## [0.1.5](https://github.com/paralus/relay/compare/v0.1.4...v0.1.5) (2023-08-11)

### Bug Fixes

* fix for readonly users should not access secrets ([#64](https://github.com/paralus/relay/issues/64)) ([a429d46](https://github.com/paralus/relay/commit/a429d4656261454da80b34f1bbc6a31812c6e92a))

### [0.1.4](https://github.com/paralus/relay/compare/v0.1.3...v0.1.4) (2023-03-31)

### Bug Fixes

* intermittent relay connection failures ([#48](https://github.com/paralus/relay/issues/48)) ([5fde31a](https://github.com/paralus/relay/commit/5fde31aa17545ba3d2e917a5668ac2615ccac997))
* same bootstrap register requests from diff target clusters ([#47](https://github.com/paralus/relay/issues/47)) ([cdf6f39](https://github.com/paralus/relay/commit/cdf6f39fa7ffed06bf84b6702f168c78a537cf70))

## [0.1.3] - 2023-02-24
### Added
-  Configure the SA account lifetime from [mabhi](https://github.com/mabhi)

## [0.1.2] - 2022-09-30
### Added
- Add arm builds from [meain](https://github.com/meain)

## [0.1.1] - 2022-08-26
### Changed
- Removed building images to registry on pull requests [niravparikh05](https://github.com/niravparikh05)
- Changes after fixing lint issue due to buf in paralus core [vivekhiwarkar](https://github.com/vivekhiwarkar)

### Fixed

## [0.1.0] - 2022-06-22
### Added
- Initial release

[Unreleased]: https://github.com/paralus/relay/compare/v0.1.7...HEAD
[0.1.7]: https://github.com/paralus/relay/compare/v0.1.6...v0.1.7
[0.1.6]: https://github.com/paralus/relay/compare/v0.1.5...v0.1.6
[0.1.5]: https://github.com/paralus/relay/compare/v0.1.4...v0.1.5
[0.1.4]: https://github.com/paralus/relay/compare/v0.1.3...v0.1.4
[0.1.3]: https://github.com/paralus/relay/compare/v0.1.2...v0.1.3
[0.1.2]: https://github.com/paralus/relay/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/paralus/relay/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/paralus/relay/releases/tag/v0.1.0
