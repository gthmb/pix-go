Learnin' me up some Go with Pixability's `golang-chapter`

# Week 2: Card Dealer Server

## Running

To run 'normally':

```
go run main.go
```

To force a `time.Sleep` of 2 seconds during the `deck.DrawCards` method (For testing concurrency with a forced latency):

```
go run main.go -slowly
```

## Testing

```
go test ./...
```

# API

There is a basic API detailed below. There is no database, but each module writes its own data to disk, and will load it's respective data file within its `init` function to persist data across restarts of the program.

There are 4 resources available over `http`:

- `decks`
- `games`
- `players`
- `playergames`

### Decks

> A `deck` contains a list of cards and an Index that signifies the current 'top'of the deck.

Get the details for a `deck` resource:

```
GET /decks/:deckId
```

### Games

> A `game` contains a `DeckID`, a `Started` flag, and an array of `PlayerGames` instances that track which players are in the game.

Get a list of all the `game` resources:

```
GET /games
```

Create a new `game` resource:

```
POST /games
```

Get the details for a `game` resource:

```
GET /games/:gameId
```

Attempt to start a `game`:

```
POST /games/:gameId/start
```

### Players

> A `player` just has an `ID`.

Gets a list of all the `player` resources:

```
GET /players
```

Creates a new `player` resource:

```
POST /players
```

Gets the details for a `player` resource:

```
GET /players/:playerId
```

### PlayerGames

> A `playergame` has a `PlayerID` and a `GameID` and joins a `player` instance to a `game` instance

Get a list of all the `playergame` resources:

```
GET /playergames
```

Create a new `playergame` resource:

```
POST /playergames

{
    "PlayerID": string,
    "GameID": string,
}
```

Get the details for a `playergame` resource:

```
GET /playergames/:playergameId
```

Attempt to draw the `NumberOfCards` from the `game` decks into the `player` deck:

```
GET /playergames/:playergameId/draw

{
    "NumberOfCards": int
}
```
