# Legato

The reference implementation for Harmony, a federated and decentralised chat service, combining the amenities of mainstream messengers with an extensible and decentralised core.

Legato is designed for public-facing deployments. For personal deployments, check out [Scherzo](https://github.com/harmony-development/scherzo/), which requires zero setup to get an instance running.

<br>Join our <a href="https://discord.gg/Abr3upBjHy">project chat</a> for announcements, support and contribution.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/harmony-development/legato/Harmony%20Build?style=for-the-badge)
![Codecov](https://img.shields.io/codecov/c/gh/harmony-development/legato?style=for-the-badge)
![Mom Made Pizza Rolls](https://img.shields.io/badge/mom%20made-pizza%20rolls-green?style=for-the-badge)

## Building
Required dependecies:
- go
- libvips
- libvips-dev

Then run:
```sh
go get -v -t -d ./...
go build
```

### Nix
If you have `nix` installed:
- with Flakes: `nix build` (or you can install it with `nix profile install github:harmony-development/legato`)
- with legacy (without flakes): `nix-build nix/default.nix`

## Usage
Make sure to install `postgres` database

Make the following preparations:
- Run command ```./legato``` for the first time to generate new json config
- Run command ```./legato -genkey``` to generate new server key
- setup `postgres` user
- edit `DB` section inside json config to connect `postgres` database

After all preparations run ```./legato``` again to start server

## Docker-compose
See the [harmony-development/orchestration](https://github.com/harmony-development/orchestration) repo for more details on a docker-compose setup.
