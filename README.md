# Alai Server

This is the server which will host the Alai game binaries, backend API that receives game data, and maybe at some point a frontend for all of that.

## Authors

- Christohper Cromer
- Martín Araneda Acuña

## Backend API

The backend is written in *go* and uses gorm to handle its database schema.

## Game

The directory game should contain the binaries for Alai which will be exposed through nginx for downloading.

## Build and run

```
docker-compose up -d --build
```

## Stop the server

```
docker-compose down
```

## License
The server is licensed under the [3-Clause BSD License](LICENSE).