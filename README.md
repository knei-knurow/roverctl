# roverctl

Program to control the Knurów rover.

Use `--help` to learn more about it.

## Frame formats

We use our own [package frames](https://github.com/knei-knurow/lidar-tools/tree/master/frames) to construct frames we send through <interface> (USART? SPI?).

Example frames below.

## Go commands

`MT+G${S}#${CRC}`

where:

S = speed, signed byte (from -128 to 127)

- CRC - calculated simple 8-bit CRC checksum of frame's all bytes (except the last byte, which is the CRC itself)

`G` indicates that this is a `GO` command

## Turn commands

`MT+T${D}#${CRC}`

where:

- D - degrees, signed byte (from -128 to 127)

- CRC - calculated simple 8-bit CRC checksum of frame's all bytes (except the last byte, which is the CRC itself)

`T` indicates that this is a `TURN` command

## Future

In the future, it might just serve as a client to [roverd](https://github.com/knei-knurow/roverd).
[Similar architecture is used by Docker](https://docs.docker.com/get-started/overview/#docker-architecture)
