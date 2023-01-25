# Alai Server

This is the server which will host the Alai site, API, and game downloads.

## Authors

- Christohper Cromer
- Martín Araneda Acuña

## License
The server is licensed under the [3-Clause BSD License](LICENSE).

## Backend API

The backend is written in *go* and uses gorm to handle its database schema.

## Game

The directory download should contain the binaries for Alai which will be exposed through nginx for downloading.

## Configure

Before the server can be built and run, a *.env* file should be created based on *.env.example*. If the server is going to be public facing, make sure to change all the secrets and passwords to make it more secure.

## Build and run

```
docker-compose up -d --build
```

## Stop the server

```
docker-compose down
```
