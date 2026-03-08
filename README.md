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

Note: `| jq` is added for nice json formatting, this is optional

```bash
curl http://localhost:8080/api/v1/health | jq

curl http://localhost:8080/api/v1/artists | jq # Get all artists
curl http://localhost:8080/api/v1/artists/12 | jq # Get specific artist
curl http://localhost:8080/api/v1/artists/12/albums | jq # Get albums by specific artist

curl curl http://localhost:8080/api/v1/tracks/7 | jq 

curl http://localhost:8080/api/v1/playlists | jq # Get the names and ids of all playlists
curl http://localhost:8080/api/v1/playlists/3 | jq # Get all the names and composers of the tracks in playlist ID 3 (or any ID value)
curl http://localhost:8080/api/v1/playlists/3/tracks | jq

curl http://localhost:8080/api/v1/albums | jq # Get all albums 
curl http://localhost:8080/api/v1/albums/9 | jq
```

Uses one of the sample neon postgres databases (chinook)

## TO DO

- Make output of playlist retriever readable
- Ability to add/remove songs from a playlist
- GET for all the other info - eg see details about a song in a playlist
- ...

To do

Models for:
[x] album
    [x] singular response function
    [x] plural response function
[x] artist
    [x] singular response function
    [x] plural response function
[x] playlist
    [x] singular response function
    [x] plural response function
[x] track
    [x] detailed singular response function
    [x] singular response func
    [x] plural response function

Standard Response types for:
[x] album
[x] artist
[x] playlist
[x] track simple and full

handlers:
[x] all albums
[x] album by id
[x] all albums by an artist id

[x] all artists
[x] artist by id

[x] all playlists
[x] playlist by id

[x] track by id
[ ] tracks in an album
[x] tracks in a playlist
[ ] tracks by an artist
