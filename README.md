<h1 align=center>
 <a href="https://github.com/harmony-development"><img align=center width="15%" src="https://i.imgur.com/lx2RCZj.png" /></a> Harmony
</h1>
A free and open source communications platform.
<br>Designed as an open source Discord replacement; with messaging, guilds, roles, voice chat and rich presence. 
<br>Join our <a href="https://discord.gg/jypXPA4">project chat</a> for announcements, support and contribution.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/harmony-development/legato/Harmony%20Build?style=for-the-badge)
![Codecov](https://img.shields.io/codecov/c/gh/harmony-development/legato?style=for-the-badge)
![Mom Made Pizza Rolls](https://img.shields.io/badge/mom%20made-pizza%20rolls-green?style=for-the-badge)

## Building
Required dependecies:
- go
- libvips
- libvips-dev

Then run:
```
go get -v -t -d ./...
go build
```

## Usage
To run server make sure to install:
- go
- postgres

Make the following preparations:
- Run command ```./legato -genkey``` to generate new server config
- setup postgres user
- update json config

After all preparations run server:  
```./legato```

## Docker-compose
Broke, WIP
