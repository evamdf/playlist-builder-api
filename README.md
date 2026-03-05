# Playlist Builder API

This project is a work in progress and currently incomplete :)

## Running

Create `.env` file with variables:

```sh
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=api-project
DB_SSLMODE=disable
```

```bash
go run main.go
```

```bash
curl http://localhost:8080/api/v1/health

curl http://localhost:8080/api/v1/artists # Get all artists
curl http://localhost:8080/api/v1/artists/12 # Get specific artist
curl http://localhost:8080/api/v1/artists/:id/albums # Get albums by specific artist

curl http://localhost:8080/api/v1/playlists # Get the names and ids of all playlists
curl http://localhost:8080/api/v1/playlists/3 # Get all the names and composers of the tracks in playlist ID 3 (or any ID value)

curl http://localhost:8080/api/v1/albums # Get all albums 

```

Uses one of the sample neon postgres databases (chinook)

## TO DO

- Make output of playlist retriever readable
- Ability to add/remove songs from a playlist
- GET for all the other info - eg see details about a song in a playlist
- ...
