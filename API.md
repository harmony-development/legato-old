# Harmony API
## Overview
All of the API is structured like this :
```json
{
    "type": string,
    "data": interface{}
}
```
This applies to all serverbound and clientbound packets.

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
* **GetServers Event**
	```json
	{
		"type": "getservers",
		"data": {
			"token": string
		}
	}
	```

## Client Bound API
* **Deauth**
	```json
	{
		"type": "Deauth",
		"data": {
			"message": "token is missing or invalid"
		}
	}
	```