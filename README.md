# Anthelion Uploader
## Disclaimer

I am in no way offiliated with Anthelion, their staff, their website, or their community.
Any code belonging to this repository is in no way offiliated with Anthelion, their staff, their website, or their community. 

## About

Anthelion Uploader is a small program written in Go that interacts with Anthelion's official API to upload torrents. The intent of this program is to help uploaders by speeding up and automating some of the process.

## System Support
Being written in Go, binaries can compile for any major system and architecture. Unfortunately, in it's current state, the program is reliant on external bash script execution, meaning that you may run into issues running this under a Windows environment.

## Setup
#### Install Dependencies
You will need the following programs installed and accessible from the commandline to use this program:

- [bash](https://www.gnu.org/software/bash/)
- [mktorrent](https://github.com/pobrn/mktorrent) 
- [mediainfo](https://github.com/MediaArea/MediaInfo)

#### Set environment variables
The following environment variables should be set:

```sh
export ANTHELION_API_KEY=<your privae api key>
export ANTHELION_ANNOUNCE_URL=<your private announce url>
export ANTHELION_API_URL=<anthelion's api url>
```

## Building

There are 6 Makefile bootstrapping directives for building under specific architectures. The example below is for building for 64bit MacOS, but if your system is different please consult the Makefile.
```sh
cd antherion_uploader
make build_mac 
```
The newly build binary can be found in the bin/ directory.

## Running
#### Flags
- -p <File Path> # This is the path to video file you intend to use
- -t <tmdbID> # This is the id from The Movied Database
- -u <announce url> # This is your personal announce URL from Anthelion
- -k <api key> # This is the API key you generated for Anthelion

The announce url and the api key can be ignored if they are set as environment variables (recommended).
#### Example run
```sh
./anthelion_uploader -p ~/path/to/cool/movie/Movie.x264.Bluray.mkv -t 1337
```

If it takes a while to run, be patient. Hashing the .torrent file of large files can take awhile. 

If it ran successfully, it will have downloaded the new torrent from Anthelion automatically, and it should be in the same directory as the binary.
