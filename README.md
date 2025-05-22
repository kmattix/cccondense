# cccondense

`cccondense` is a simple command-line utility for condensing `.srt` (SubRip Subtitle) files by merging consecutive blocks with identical timestamps. This is useful for cleaning up subtitles that are fragmented or redundant.

## Features

* Parses `.srt` files into structured blocks
* Merges subtitle entries that share the same timestamp
* Outputs a clean, renumbered `.srt` file

## Usage

```bash
cccondense <input.srt> <output.srt>
```

### Example

```bash
cccondense subtitles/input.srt subtitles/output.srt
```

This will read `input.srt`, condense entries with the same timestamp, and write the result to `output.srt`.

## Build Instructions

Precompiled binaries for **Windows**, **macOS**, and **Linux** are available in the `build/` directory.

### From Source

To build from source, ensure you have [Go installed](https://golang.org/dl/) (version 1.24.3 or later), clone the repository, then run:

```bash
go build -o cccondense cccondense.go srtcondenser.go
```
or run this without needing to clone:
```bash
go install github.com/kmattix/cccondense
```

## Notes

* This tool expects well-formed `.srt` files.
* Malformed entries may cause the tool to panic with a descriptive error.

## License

MIT License
