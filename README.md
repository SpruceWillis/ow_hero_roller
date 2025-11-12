# ow_hero_roller
For those times when you don't know which hero to play

## Setup

The project is written for Go 1.21.4. Find the installation instructions [here](https://go.dev/doc/manage-install)

When developing, you will need an instance of the Discord client, which can be installed from [here](https://discord.com/download). Once that is set up, you can enable developer mode following [these instructions](https://beebom.com/how-enable-disable-developer-mode-discord/)

Once the developer settings are enabled, you should enable Application Test Mode from the `Advanced` user settings tab. The application ID can be provided on request.

## Running the application locally

Building and running the binary:
```
go build .
./ow_hero_roller -d hero_data.conf
```

To run properly, the application requires two environment variables:
* `TENOR_API_KEY` is used for retrieving GIFs from the Tenor API. Instructions for signup can be found [here](https://developers.google.com/tenor/guides/quickstart)
* `BOT_TOKEN` is the API key for the application itself. This will need to be requested (or generated if you are forking this project)

The container build supports variables for your development machine:
```
docker build --build-arg OS=${YOUR_OS} --build-arg ARCH=${YOUR_ARCH} .
```

To container will require that the hero data file and SSL certs to be mounted into the container, in addition to the environment variables.
