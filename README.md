## Chinese Checkers

- [Chinese Checkers](#chinese-checkers)
  - [Project Outline](#project-outline)
  - [Quick Start Guide](#quick-start-guide)
  - [Documentation](#documentation)
  - [Websocket Operations Documentation](#websocket-operations-documentation)

### Project Outline

Project outline is available at [Chinese Checkers DrawIO UML](https://drive.google.com/file/d/1iIDCE1dcRjzx1F8HkmPSoind6I9Joa1u/view?usp=sharing)

### Quick Start Guide

Clone the repository and run the server

```bash
git clone https://github.com/Rafisto/chinese-checkers.git
cd chinese-checkers
```

Build and run *the server* using the following commands

```bash
cd server
go build
./chinese-checkers
```

Build and run *the client* using the following commands

```bash
cd client
go build
./chinese-checkers-client
```

### Documentation

Swagger is available for the Server API at [http://localhost:8080/swagger](http://localhost:8080/swagger/)

You can run the documentation locally using `godoc` via

```bash
# install godoc
go install -v golang.org/x/tools/cmd/godoc@latest

cd server
godoc -http=:6060
```

Then navigate to [http://localhost:6060/pkg/chinese-checkers/](http://localhost:6060/pkg/chinese-checkers/)

### Websocket Operations Documentation

Client sends to Server:

1. State of the game

```json
{
  "type": "player",
  "action": "state",
}
```

2. Get Board

```json
{
  "type": "player",
  "action": "board"
}
```

3. Get Pawns

```json
{
  "type": "player",
  "action": "pawns"
}
```

4. Send Move:

```json
{
  "type": "player",
  "action": "move",
  "player_id": 0,
  "start": { // or null if player wants to skip the turn
    "row":0,
    "col":0,
  },
  "end": { // or null -||-
    "row":0,
    "col":0,
  }
}
```
(automatically ends turn)

Server broadcasts to clients:

1. Broadcast New Move (automatically new turn):

```json
{
  "type": "server",
  "action": "move",
  "player_id": 0, // which player's turn is it
  "start": { // or null if previous player skipped his turn
    "row": 0,
    "col": 0,
  },
  "end": { // or null -||-
    "row": 0,
    "col": 0,
  }
}
```

