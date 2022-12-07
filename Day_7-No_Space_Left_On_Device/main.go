package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"golang.org/x/exp/slices"
	"math"
)

type fileNode struct {
	name string
	size int64
}

type directoryNode struct {
	name string
	files []fileNode
	subDirectories []directoryNode
	parentDirectory *directoryNode

	//size int64 // cached size of this directory + subdirectories
}

func (dir *directoryNode) FindSubdirectory(name string) int {
	return slices.IndexFunc(dir.subDirectories, func(dir directoryNode) bool { return dir.name == name })
}

func (dir *directoryNode) FindFile(name string) int {
	return slices.IndexFunc(dir.files, func(file fileNode) bool { return file.name == name })
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var root directoryNode
	root.name = "/"
	currentDirectory := &root

	for scanner.Scan() {
    	line := scanner.Text()

		if line[0] == '$' {
			currentDirectory = processNewCommand(line, currentDirectory, &root)
		} else {
			updateDirectory(line, currentDirectory)
		}
	}

	//printDirectoryTree(root, 0)
	//_, smallDirSize := processDirectoryTree(root, 100000)

	rootSize, subDirSizes := calculateDirectorySize(root)

	unusedSpace := 70000000 - rootSize
	neededSpace := 30000000 - unusedSpace
	directoryDeleteSize := int64(math.MaxInt64)

	for _, size := range subDirSizes {
		if size >= neededSpace {
			if size < directoryDeleteSize {
				directoryDeleteSize = size
			}
		}
	}
	
	fmt.Println(directoryDeleteSize)
}

func processNewCommand(newCommand string, currentDirectory *directoryNode, rootDirectory *directoryNode) *directoryNode {
	if strings.HasPrefix(newCommand, "$ cd ") {
		directoryName := strings.TrimPrefix(newCommand, "$ cd ")

		if directoryName == "/" {
			currentDirectory = rootDirectory
			return currentDirectory
		} 

		if directoryName == ".." {
			// Note: this will crash on root
			currentDirectory = currentDirectory.parentDirectory
			return currentDirectory
		}

		idx := currentDirectory.FindSubdirectory(directoryName)

		if idx == -1 {
			idx = addSubdirectory(directoryName, currentDirectory)
		}

		currentDirectory = &currentDirectory.subDirectories[idx]
	} else if strings.HasPrefix(newCommand, "$ ls") {
		// We technically don't need to do anything right now, just keep reading
	}

	return currentDirectory
}

func updateDirectory(newNode string, currentDirectory *directoryNode) {
	if strings.HasPrefix(newNode, "dir ") {
		directoryName := strings.TrimPrefix(newNode, "dir ")
		idx := currentDirectory.FindSubdirectory(directoryName)
		if idx == -1 {
			addSubdirectory(directoryName, currentDirectory)
		}
	} else {
		split := strings.Split(newNode, " ")
		idx := currentDirectory.FindFile(split[1])
		if idx == -1 {
			size, _ := strconv.ParseInt(split[0], 10, 0)
			addFile(size, split[1], currentDirectory)
		}
	}
}

func addSubdirectory(name string, currentDirectory *directoryNode) int {
	var newDirectory directoryNode
	newDirectory.name = name 
	newDirectory.parentDirectory = currentDirectory
	currentDirectory.subDirectories = append(currentDirectory.subDirectories, newDirectory)
	return len(currentDirectory.subDirectories)-1
}

func addFile(size int64, name string, currentDirectory *directoryNode) int {
	var newFile fileNode
	newFile.name = name
	newFile.size = size
	currentDirectory.files = append(currentDirectory.files, newFile)
	return len(currentDirectory.files)-1
}

// {directorySize, AccumulatedSmallDirectorySize}
func processDirectoryTree(node directoryNode, smallDirectorySize int64) (int64, int64) {
	directorySize := int64(0)
	accumulatedSmallDirSize := int64(0)

	for _, dir := range node.subDirectories {
		subDirSize, subDirSmallSize := processDirectoryTree(dir, smallDirectorySize) 
		directorySize += subDirSize
		accumulatedSmallDirSize += subDirSmallSize
	}

	for _, file := range node.files {
		directorySize += file.size
	}

	if directorySize < smallDirectorySize {
		accumulatedSmallDirSize += directorySize
	}

	return directorySize, accumulatedSmallDirSize
}

// DirectorySize, list of the size of all sub directories
func calculateDirectorySize(node directoryNode) (int64, []int64) {
	directorySize := int64(0)
	subDirectoriesSize := []int64{}

	for _, dir := range node.subDirectories {
		subDirSize, subDirectories := calculateDirectorySize(dir) 
		directorySize += subDirSize
		subDirectoriesSize = append(subDirectoriesSize, subDirSize)
		subDirectoriesSize = append(subDirectoriesSize, subDirectories...)
	}

	for _, file := range node.files {
		directorySize += file.size
	}

	return directorySize, subDirectoriesSize
}

func printDirectoryTree(node directoryNode, depth int) {
	space := ""
	for i := 0; i < depth; i++ {
		space += " "
	}

	space += " "
	for _, file := range node.files {
		fmt.Println(space+"* "+file.name+"("+strconv.FormatInt(file.size,10)+")")
	}

	for _, dir := range node.subDirectories {
		printDirectoryTree(dir, depth+1)
	}
}