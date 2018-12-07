# release-bot

[![Build Status](https://img.shields.io/travis/karriereat/release-bot.svg?style=flat-square)](https://travis-ci.org/karriereat/release-bot)
[![Go Report Card](https://goreportcard.com/badge/github.com/karriereat/release-bot?style=flat-square)](https://goreportcard.com/report/github.com/karriereat/release-bot)
[![license](https://img.shields.io/badge/license-Apache%202.0-brightgreen.svg?style=flat-square)](https://github.com/karriereat/release-bot/blob/master/LICENSE)

A bot that converts GitLab tag pushes to Slack messages


## Installation
- copy the `sample.toml` from `/conf` and edit it to your needs
- run `release-bot -c config.toml` - the `-c` flag can be ommited if there is a `config.toml` in the same directory as the binary
- Point the gitlab webhook to `http://your.domain/hooks/gitlab`


## Features
Currently the release-bot only understands `gitlab tag pushes` and `slack` notifications


## Demo
![Message Example](assets/example-message.png)