# Ability Stone Simulator
This repo holds both a client app and a server app to simulate cutting an Ability Stone in Lost Ark

The client app servers a user facing interface to present the simulation. The server app, which the client app will call to make calculations, will do the simulation calculations.

This architecture was chosen as the server app is planned to be a standalone API that can talk to many clients if necessary. This is important, say, if one were to use the API so they could apply machine learning to find a somewhat optimum way to get a consistently good Ability Stone.

---
## Running the apps
To run the server app:
1. `cd server-app`
2. `go run main.go`

To run the client app:
1. `cd **client`
2. `npm install`
3. `npm run start`

Note that you will need the server-app to run the client-app