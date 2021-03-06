The Things Network
==================

[![Build Status](https://travis-ci.org/TheThingsNetwork/ttn.svg?branch=develop)](https://travis-ci.org/TheThingsNetwork/ttn) [![Slack Status](https://slack.thethingsnetwork.org/badge.svg)](https://slack.thethingsnetwork.org/) [![Coverage Status](https://coveralls.io/repos/github/TheThingsNetwork/ttn/badge.svg?branch=develop)](https://coveralls.io/github/TheThingsNetwork/ttn?branch=develop)

![The Things Network](http://thethingsnetwork.org/static/ttn/media/The%20Things%20Uitlijning.svg)

The Things Network is a global open crowdsourced Internet of Things data network.

## Getting Started With The Things Network

When you get started with The Things Network, you'll probably have some questions. Here are some things you can do to find the answer to them:

- Check out our [website](https://www.thethingsnetwork.org/)
- Read the [official documentation](https://www.thethingsnetwork.org/docs/)
- Register on the [forum](https://www.thethingsnetwork.org/forum/) and search around
- Join [Slack](https://slack.thethingsnetwork.org) and ask us what you want to know
- Read background information on the [wiki](https://www.thethingsnetwork.org/wiki/)

## Prepare your Development Environment

1. Make sure you have [Go](https://golang.org) installed.
2. Set up your [Go environment](https://golang.org/doc/code.html#GOPATH)
3. Install the [protobuf compiler (`protoc`)](https://github.com/google/protobuf/releases)
4. Make sure you have [Mosquitto](http://mosquitto.org/download/) and [Redis](http://redis.io/download) **installed** and **running**.  
   If you're on Linux, you probably know how to do that. On a Mac, just run `brew bundle`. Windows users are currently on their own (feel free to contribute a guide for Windows).

## Set up The Things Network's backend for Development

1. Fork this repository
2. Clone your fork: `git clone --branch v2-preview --recursive https://github.com/YOURUSERNAME/ttn.git $GOPATH/src/github.com/TheThingsNetwork/ttn`
3. `cd $GOPATH/src/github.com/TheThingsNetwork/ttn`
4. Install the dependencies for development: `make dev-deps`
5. Run the tests: `make test`

**NOTE:** From now on you should run all commands from the `$GOPATH/src/github.com/TheThingsNetwork/ttn` directory.

## Build, install and run The Things Network's backend locally

Additional to setting up the backend for development, you should do the following to build, install and run the backend locally.

1. Configure your `ttnctl` with the settings in `.env/ttnctl.yml.dev-example` by copying that file to `~/.ttnctl.yml`.
2. Run `make build` to build both `ttn` and `ttnctl` from source. 
2. Run `make link` to link them to `$GOPATH/bin/` (In order to run the commands, you should have `export PATH="$GOPATH/bin:$PATH"` in your profile).
3. Run `forego start` to start all backend services at the same time. Make sure that Redis and Mosquitto **are running** on your machine.
4. First time only (or when Redis is flushed):
  * Run `ttn broker register-prefix 00000000/0 --config ./.env/broker/dev.yml`
  * Restart the backend services

## Build and run The Things Network's backend in Docker

1. Configure your `ttnctl` with the settings in `.env/ttnctl.yml.dev-example` by copying that file to `~/.ttnctl.yml`.
2. Add the following line to your `/etc/hosts` file:
    `127.0.0.1 router handler`
3. Run `make docker` to build the docker image
4. Run `docker-compose up` to start all backend services in Docker. Make sure that Redis and Mosquitto **are not running** on your local machine.
5. First time only (or when Redis is flushed):
  * Run `docker-compose run broker broker register-prefix 00000000/0 --config ./.env/broker/dev.yml`
  * Restart the backend

## Contributing

Source code for The Things Network is MIT licensed. We encourage users to make contributions on [Github](https://github.com/TheThingsNetwork/ttn) and to participate in discussions on [Slack](https://slack.thethingsnetwork.org).

If you encounter any problems, please check [open issues](https://github.com/TheThingsNetwork/ttn/issues) before [creating a new issue](https://github.com/TheThingsNetwork/ttn/issues/new). Please be specific and give a detailed description of the issue. Explain the steps to reproduce the problem. If you're able to fix the issue yourself, please help the community by forking the repository and submitting a pull request with your fix.

For contributing a feature, please open an issue that explains what you're working on. Work in your own fork of the repository and submit a pull request when you're done.

If you want to contribute, but don't know where to start, you could have a look at issues with the label [*help wanted*](https://github.com/TheThingsNetwork/ttn/labels/help%20wanted) or [*difficulty/easy*](https://github.com/TheThingsNetwork/ttn/labels/difficulty%2Feasy).

## License

Source code for The Things Network is released under the MIT License, which can be found in the [LICENSE](LICENSE) file. A list of authors can be found in the [AUTHORS](AUTHORS) file.
