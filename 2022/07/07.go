package main

import (
	"bufio"
	"io"
	"path"
	"strconv"
	"strings"
	"unicode"
)

type dir struct {
	path  string
	files []*file
	dirs  map[string]*dir
	size  int
}

type file struct {
	size int
	name string
}

func walk(n *dir) {
	for _, d := range n.dirs {
		walk(d)
	}

	s := 0
	for _, f := range n.files {
		s += f.size
	}
	for _, d := range n.dirs {
		s += d.size
	}
	n.size = s
}

func run(input io.Reader) (int, int) {

	scanner := bufio.NewScanner(input)

	root := &dir{
		path:  "/",
		files: []*file{},
		dirs:  make(map[string]*dir),
		size:  0,
	}
	allDirs := make(map[string]*dir)
	allDirs[root.path] = root

	curDir := root
	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ cd /" {
			continue
		}
		split := strings.Split(line, " ")

		if split[0] == "$" {
			switch split[1] {
			case "cd":

				if split[2] == ".." {
					upPath := allDirs[path.Join(curDir.path, "..")]
					curDir = upPath
					continue
				}

				path := path.Join(curDir.path, split[2])
				exDir, ok := curDir.dirs[path]
				if !ok {
					newDir := dir{
						path:  path,
						files: []*file{},
						dirs:  make(map[string]*dir),
						size:  0,
					}
					allDirs[newDir.path] = &newDir
					curDir.dirs[newDir.path] = &newDir
					curDir = &newDir
				} else {
					curDir = exDir
				}
			case "ls":
				continue
			}

		} else if unicode.IsDigit(rune(split[0][0])) {
			size, _ := strconv.Atoi(split[0])
			newFile := file{
				size: size,
				name: split[1],
			}
			curDir.files = append(curDir.files, &newFile)

		} else if split[0] == "dir" {
			if _, ok := curDir.dirs[split[1]]; !ok {
				newDir := dir{
					path:  path.Join(curDir.path, split[1]),
					files: []*file{},
					dirs:  make(map[string]*dir),
					size:  0,
				}
				allDirs[newDir.path] = &newDir
				curDir.dirs[newDir.path] = &newDir
			}
		}
	}

	walk(root)

	s := 0
	requiredSpace := 30000000 - (70000000 - root.size)
	candidateDir := root
	for _, dir := range allDirs {
		if dir.size < 100000 {
			s += dir.size
		}

		if dir.size > requiredSpace && dir.size < candidateDir.size {
			candidateDir = dir
		}
	}

	return s, candidateDir.size
}
