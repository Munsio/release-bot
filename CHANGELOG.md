# Changelog

All notable changes to this project will be documented in this file.

## [v0.2.1] - 2019-02-11

### Fixed
- check if gitlab sends double tag_push event - https://gitlab.com/gitlab-org/gitlab-ce/issues/52560

## [v0.2.0] - 2019-01-21

### Added
- CHANGELOG.md
- New option to skip sslVerification for Gitlab #1
- Use `:releasegopher:` emoji for slack message avatar
- Check for `%SKIP-NOTIFY%` in release message to skip notification

### Changed
- Show versions without release message #4
- Shorter title text
- Use namespace+repo instead of repo name in title #3