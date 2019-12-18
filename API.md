
# Harmony API
## Overview
All of the API is structured like this :
```json
{
    "type": string,
    "data": interface{}
}
```
This applies to all server-bound and client-bound packets.

# Table Of Contents
- [<h2>Server Bound API</h2>](#server-bound-api)
  - [**Login Event**](#login-event)
  - [**Register Event**](#register-event)
  - [**GetGuilds Event**](#getguilds-event)
  - [**JoinGuild Event**](#joinguild-event)
- [<h2>Client Bound API</h2>](#client-bound-api)
  - [**Token Event**](#token-event)
  - [**GetGuilds Event**](#getguilds-event)
  - [**GetChannels Event**](#getchannels-event)
  - [**Deauth Event**](#deauth-event)
  - [**Message Event**](#message-event)
  - [**GetMessages Event**](#getmessages-event)

## Server Bound API
* ### Login Event
	```json
	{
		"type": "login",
		"data": {
			"email": string,
			"password": string
		}
	}
	```
* ### Register Event
	```json
	{
		"type": "register",
		"data": {
			"email": string,
			"username": string,
			"password": string
		}
	}
	```
* ### GetGuilds Event
  ```json
	{
		"type": "getguilds",
		"data": {
			"token": string
		}
	}
	```
* ### JoinGuild Event
  ```json
  {
    "type": "joinguild",
    "data": {
      "token": string,
      "invitecode": string
    }
  }
  ```

## Client Bound API
* ### Token Event
  > Returns a token and userid, to use for subsequent requests to the server.
  ```json
  "data": {
    "token": string,
    "userid": string
  }
  ```
* ### GetGuilds Event
  > Returns a guildid-guild pair, containing the guild's name, picture, and whether the user is the owner or not
  ```json
  "data": {
    "guilds": {
      "guildid": {
        "guildname": string,
        "picture": string,
        "owner": string
      }
    }
  }
  ```
* ### GetChannels Event
  > Returns an channelid-channelname pair
  ```json
  "data": {
    "channels": {
      "guildid": string
    }
  }
  ```
* ### Deauth Event
  > Returns a request to deauthenticate the client and send em back to the login screen ( comes with a bonus message )
  ```json
  "data": {
    "message": string
  }
  ```
* ### Error Event
	> Returns a general-purpose error message to be displayed to the client
	```json
	"data": {
		"message": string
	}
  ```
* ### Message Event
  > Returns the data for a new message being received from the server
  ```json
  "data": {
    "messageid": string,
    "createdat": number,
    "guild": string,
    "channel": string,
    "userid": string,
    "message": string
  }
  ```

* ### GetMessages Event
  > Returns an array of the 30 messages in a specific channel
  ```json
  "data": {
    "messages": [
      {
        "messageid": string,
        "createdat": number,
        "guild": string,
        "channel": string,
        "userid": string,
        "message": string
      }
    ]
  }
  ```
