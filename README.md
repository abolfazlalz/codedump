# codedump

`codedump` is a simple CLI tool written in Go that scans a project directory and prints the contents of source files in a structured format.

It’s useful when you want to quickly inspect, share, or export your project’s code.

## Features

- Recursively scans the current directory  
- Filters files by extension  
- Skips binary files automatically  
- Optional line numbers  
- Optional Markdown code blocks  
- File size limit to avoid large files  
- Include or exclude specific paths  
- Sorted output for consistent results  

## Installation

```bash
go install github.com/abolfazlalz/codedump@latest
```

## Usage

Run inside your project directory:

```bash
codedump
```

### With line numbers

```bash
codedump --line-numbers
```

### With Markdown formatting

```bash
codedump --markdown
```

### Include only specific directories

```bash
codedump --include=cmd,internal
```

### Exclude specific directories

```bash
codedump --exclude=node_modules,vendor
```

## Flags

```
-ext string          Comma-separated file extensions
-exclude string      Comma-separated paths to exclude
-include string      Only include files under these paths
-markdown            Wrap output in Markdown code blocks
-line-numbers        Print line numbers
-max-size int        Max file size in KB
```

## License

MIT
