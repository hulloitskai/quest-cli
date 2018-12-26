# quest-cli

_A command-line client for the UWaterloo Quest Information System._

[![Release][release-img]][release]
[![Go Report Card][grp-img]][grp]
[![Travis: Build][travis-img]][travis]

Check your grades in style 😎.

<img src="./docs/demo.gif" width="725px" />

### Why?

lol idk tbh

uhhhhh quest web interface bad 😭 💯

## Features

- [x] Login and authentication
- [x] Check grades
- [x] Windows compatibility
- [ ] Check schedule?
- [ ] Check course availability?
- [ ] ??? other stuff ???

[Open an issue](https://github.com/stevenxie/quest-cli/issues/new) to request
a new feature that you want to see added to `quest-cli`!

<br />

## FAQ

#### Will this steal my Quest login?

Yes.

Okay, not really, but don't take that from me; check the source code for
yourself:

- [`quest-cli/internal/interact/quest.go`](https://github.com/stevenxie/quest-cli/blob/master/internal/interact/quest.go)
- [`uwquest/login.go`](https://github.com/stevenxie/uwquest/blob/master/login.go) (underlying Quest accessor library)

Your credentials are sent to the UWaterloo IDP (ID portal) as if you're
logging into the Quest website itself.

[grp]: https://goreportcard.com/report/github.com/stevenxie/quest-cli
[grp-img]: https://goreportcard.com/badge/github.com/stevenxie/quest-cli
[release]: https://github.com/stevenxie/quest-cli/releases
[release-img]: https://img.shields.io/github/release/stevenxie/quest-cli.svg
[travis]: https://travis-ci.com/stevenxie/quest-cli
[travis-img]: https://travis-ci.com/stevenxie/quest-cli.svg?branch=master
