# Video Filter

This project is a Go application for applying filters to video files with the help of FFmpeg.

## Example

```sh
go run main.go -f 30 -p 3 -i input.mp4 -o output.mp4
```

[![input.mp4](/example/input.png)](/example/input.mp4)
[![output.mp4](/example/output.png)](/example/output.mp4)

## Why does it exists?

-   The goal of this project is for me to practice Go.

## What is missing?

-   There is no input validation.
-   All the filters are simple in nature, as the goal of this project was not to study computer graphics.

## Requirements

-   Go 1.23.4 or higher
-   FFmpeg

## Usage

To apply a filter to a video file, use the following command:

```sh
go run main.go -f <output fps> -p <filter profile> -i <input file name> -o <output file name>
```

To see all flags:

```sh
go run main.go -h
```
