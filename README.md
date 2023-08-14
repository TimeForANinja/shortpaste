# Short{Paste}

[![Drone Build Status](https://drone.adyanth.site/api/badges/adyanth/shortpaste/status.svg)](https://drone.adyanth.site/adyanth/shortpaste)
[![Docker pulls](https://img.shields.io/docker/pulls/adyanth/shortpaste.svg)](https://hub.docker.com/r/adyanth/shortpaste)
[![Docker Compose](https://img.shields.io/badge/Compose-docker--compose.yml-blue)](https://git.adyanth.site/DockerComposeApps/shortpaste/src/branch/main/docker-compose.yml)
[![Go Report](https://goreportcard.com/badge/git.adyanth.site/adyanth/shortpaste)](https://goreportcard.com/report/git.adyanth.site/adyanth/shortpaste)
[![Demo on Heroku](https://img.shields.io/badge/Demo-Heroku-7056bf)](https://shortpaste.herokuapp.com/)

`Short{Paste}` is open-source software written in Go and VueJS. It is a minimalistic shortener that combines three things that need shortening: links, files, and text. It is a self-hosted alternative to many popular services like paste bin and using dropbox to send a file to someone quickly.

The Go backend handles saving files, links, and text in DB and filesystem as needed, while the Vue UI provides a pretty view for you to add and review content. Added bonus, it tracks hit counts too!

## Deployment

The whole backend packages to a single binary, and I bundled the app as a docker container based on `alpine` favored for its tiny size. To run this yourself, execute the command below.

```bash
docker run -d \
    -p 8080:8080 \
    -v ${PWD}/shortpaste/:/root/.shortpaste \
    adyanth/shortpaste:latest
```

The command will publish the application on port `8080`, making it available on `http://localhost:8080/` and use the bind-mounted folder called `shortpaste` in your current working directory to save the SQLite DB, the files, and texts published.

If you prefer docker-compose, here is an [example deployment](https://git.adyanth.site/DockerComposeApps/shortpaste).

## Build it yourself

With docker, you can build this yourself. A `Dockerfile` is provided at the root of this repository.

It uses a multi-stage build consisting of three stages:

- Go Build  : Builds a statically linked Go binary containing the backend API server and the static server.
- Vue Build : Builds the VueJS UI to generate a `dist` folder with resources.
- Container : Alpine container where the binary and `dist` are copied and served.

## Environment Variables

You can customize the behavior using environment variables. Here is a list of configurable parameters.

| Environment Variable | Default Value     | Behaviour                                                                                                          |
| -------------------- | ----------------- | ------------------------------------------------------------------------------------------------------------------ |
| `SP_BIND_ADDR`       | `":8080"`         | Sets the bind-address inside the container.                                                                        |
| `PORT`               |                   | If set, it overrides the `SP_BIND_ADDR` to use the given port on all interfaces (support for Heroku deployment)    |
| `SP_STORAGE_PATH`    | `"~/.shortpaste"` | Sets the location for saving data inside the container.                                                            |
| `SP_307_REDIRECT`    |                   | Setting this to anything causes a 307 redirect to be sent instead of showing the landing page for shortened links. |
| `SP_USERNAME`        | `admin`           | Sets the username to login to the UI (only applies to creating resources, links will still work fine)              |
| `SP_PASSWORD`        | `admin`           | Sets the password to login to the UI (only applies to creating resources, links will still work fine)              |
| `SP_NOAUTH`          |                   | Setting this to anything bypasses authentication for creating resources.                                           |

## Screenshots

Here are some screenshots to get a taste of it, see the [demo](https://shortpaste.herokuapp.com/) for more :)

| Type  |                Create                |                  View                   |
| :---: |:------------------------------------:|:---------------------------------------:|
| Links | ![Link Upload](docs/images/link.png) | ![Link View](docs/images/link_link.png) |
| Text  | ![Text Upload](docs/images/text.png) | ![Text View](docs/images/text_link.png) |
| Files | ![File Upload](docs/images/file.png) | ![File View](docs/images/file_link.png) |
