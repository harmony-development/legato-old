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

## Server Bound API
* **Login Event**
	```json
	{
		"type": "login",
		"data": {
			"email": string,
			"password": string
		}
	}
	```
*  **Register Event**
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
* **GetGuilds Event**
  ```json
	{
		"type": "getguilds",
		"data": {
			"token": string
		}
	}
	```
* **JoinGuild Event**
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
* **RegisterError**
  ```json
  {
    "type": "registererror",
    "data": {
      "message": string
    }
  }
  ```
* **LoginError**
  ```json
  {
    "type": "loginerror",
    "data": {
      "message": string
    }
  }
    ```
* **Deauth**
	```json
	{
		"type": "deauth",
		"data": {
			"message": "token is missing or invalid"
		}
	}
	```
* **Token**
  ```json
  {
    "type": "token",
    "data": {
      "token": string
    }
  }
  ```
* **GetGuilds**
  ```json
  {
    "type": "getguilds",
    "data": {
      "guilds": {
        "guildid": {
          "guildname": string,
          "picture": string
        }
      }
    }
  }
  ```
* **JoinGuild**
  ```json
  {
    "type": "joinguild",
    "data": {
      "guild": string
    }
  }
  ```

  ```
* **Message**
  ```json
  {
    "type": "message",
    "data": {
      "guild": string,
      "userid": string,
      "message": string,
      "messageid": string
    }
  }
  ```
