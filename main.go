package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	extFlag      = flag.String("ext", "", "Comma-separated file extensions (default: common code files)")
	excludeFlag  = flag.String("exclude", "", "Comma-separated paths to exclude")
	includeFlag  = flag.String("include", "", "Only include files under these paths")
	markdownFlag = flag.Bool("markdown", false, "Wrap files in markdown code blocks")
	lineNumFlag  = flag.Bool("line-numbers", false, "Print line numbers")
	maxSizeFlag  = flag.Int64("max-size", 500, "Max file size in KB")
)

var defaultExtensions = []string{
	"go", "js", "ts", "jsx", "tsx",
	"py", "java", "c", "h", "cpp", "hpp",
	"rs", "rb", "php", "cs",
	"swift", "kt", "scala",
	"sh", "bash", "lua",
	"sql", "yaml", "yml", "json",
	"toml", "ini", "env", "md",
}

func main() {
	flag.Parse()

	extensions := parseList(*extFlag, defaultExtensions)
	excludes := parseList(*excludeFlag, nil)
	includes := parseList(*includeFlag, nil)

	var files []string

	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if !hasExtension(path, extensions) {
			return nil
		}

		if isExcluded(path, excludes) {
			return nil
		}

		if len(includes) > 0 && !isIncluded(path, includes) {
			return nil
		}

		if info.Size() > (*maxSizeFlag * 1024) {
			return nil
		}

		if isBinary(path) {
			return nil
		}

		files = append(files, path)
		return nil
	})

	sort.Strings(files)

	for _, file := range files {
		printFile(file)
	}
}

// ---------------- helpers ----------------

func parseList(val string, fallback []string) []string {
	if val == "" {
		return fallback
	}
	return strings.Split(val, ",")
}

func hasExtension(path string, exts []string) bool {
	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	for _, e := range exts {
		if ext == e {
			return true
		}
	}
	return false
}

func isExcluded(path string, excludes []string) bool {
	for _, ex := range excludes {
		if strings.Contains(path, ex) {
			return true
		}
	}
	return false
}

func isIncluded(path string, includes []string) bool {
	for _, in := range includes {
		if strings.Contains(path, in) {
			return true
		}
	}
	return false
}

func isBinary(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return true
	}
	defer f.Close()

	buf := make([]byte, 800)
	n, _ := f.Read(buf)
	return bytes.IndexByte(buf[:n], 0) != -1
}

func printFile(path string) {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Printf("FILE: %s\n", path)
	fmt.Println("========================================")

	lang := strings.TrimPrefix(filepath.Ext(path), ".")

	if *markdownFlag {
		fmt.Printf("```%s\n", lang)
	}

	f, _ := os.Open(path)
	defer f.Close()

	reader := bufio.NewReader(f)
	line := 1

	for {
		l, err := reader.ReadString('\n')
		if *lineNumFlag {
			fmt.Printf("%4d | %s", line, l)
		} else {
			fmt.Print(l)
		}
		line++
		if err == io.EOF {
			break
		}
	}

	if *markdownFlag {
		fmt.Println("\n```")
	}
}
