This project is a work in progress and currently incomplete :) 

#### Running

```bash
go run main.go
```

```bash
curl http://localhost:8080/api/v1/health
curl http://localhost:8080/api/v1/artists # to get all artists
curl http://localhost:8080/api/v1/playlists # To get the names and ids of all playlists
curl http://localhost:8080/api/v1/playlists/3 # To get all the names and composers of the tracks in playlist ID 3 (or any ID value)
```

Uses one of the sample neon postgres databases (chinook)
