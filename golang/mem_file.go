package main

import "fmt"
import "strings"

type Directory struct {
    parent *Directory
    dirChild map[string]*Directory
    fileChild map[string]*File
    child []string
    name string
}

type File struct {
    parent *Directory
    name string
    data string
}

const (
    path_sep = "/"
)

var root_directory *Directory

func createDirectory(parentDirectory *Directory, dirName string) *Directory {
    newDirectory := new(Directory)
    if newDirectory == nil {
        return nil
    }
    newDirectory.parent = parentDirectory
    newDirectory.dirChild = make(map[string]*Directory)
    newDirectory.fileChild = make(map[string]*File)
    newDirectory.child = []string{}
    newDirectory.name =  dirName
    return newDirectory
}

func parent(path []string) *Directory {
    var childDir *Directory
    var ok bool
    
    parent_dir := root_directory
    
    if len(path) == 1 {
        return parent_dir
    }

    // We start with the first child    
    path = path[1:]
    for _, dir := range path {
        if childDir, ok = parent_dir.dirChild[dir]; !ok {
            return nil
        }
        parent_dir = childDir;
    }
    return parent_dir
}

func mkdir(path string) {
    if path == path_sep {
      fmt.Printf("Root directory exists\n");
      return;
    }

    pathArray := strings.Split(path, path_sep)
    parentPath := pathArray[:len(pathArray)-1]
    dirName := pathArray[len(pathArray)-1]
    
    parentDirectory := parent(parentPath)
    if (parentDirectory == nil) {
        fmt.Printf("Parent directory does not exist\n");
        return
    }
    
    if _, ok := parentDirectory.fileChild[dirName]; ok {
        fmt.Printf("Directory or File with path %v exists", path)
        return
    }

    if _, ok := parentDirectory.dirChild[dirName]; ok {
        fmt.Printf("Directory or File with path %v exists", path)
        return
    }
    
    newDir := createDirectory(parentDirectory, dirName)
    parentDirectory.dirChild[dirName] = newDir
    fmt.Printf("Directory created successfully\n")
}

func write_file(path string, data string) {
    pathArray := strings.Split(path, path_sep)
    parentPath := pathArray[:len(pathArray)-1]
    fileName := pathArray[len(pathArray)-1]
    
    parentDirectory := parent(parentPath)
    if (parentDirectory == nil) {
        fmt.Printf("Parent directory does not exist");
        return
    }

    if _, ok := parentDirectory.fileChild[fileName]; ok {
        fmt.Printf("Directory or File with path %v exists\n", path)
        return
    }
    if _, ok := parentDirectory.dirChild[fileName]; ok {
        fmt.Printf("Directory or File with path %v exists\n", path)
        return
    }
    
    newFile := new(File)
    newFile.parent = parentDirectory
    newFile.name = fileName
    newFile.data = data
    
    parentDirectory.fileChild[fileName] = newFile
    parentDirectory.child = append(parentDirectory.child, fileName)
    fmt.Printf("File %v written successfully\n", path)
}

func read_file(path string) string {
    pathArray := strings.Split(path, path_sep)
    parentPath := pathArray[:len(pathArray)-1]
    fileName := pathArray[len(pathArray)-1]
    
    parentDirectory := parent(parentPath)
    if (parentDirectory == nil) {
        fmt.Printf("Parent directory does not exist");
        return ""
    }
    
    var file *File
    var ok bool

    if file, ok = parentDirectory.fileChild[fileName]; !ok {
        fmt.Printf("File does not exist\n");
        return ""
    }
    
    return file.data
}


func main() {
    root_directory = createDirectory(nil, "")
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    
    fmt.Println("Creating directory /");
    mkdir("/")
    fmt.Println()
    
    
    fmt.Println("Creating directory /foo");
    mkdir("/foo")
    fmt.Println()

    fmt.Println("Creating directory /foo/bar");
    mkdir("/foo/bar")
    fmt.Println()

    fmt.Println("Creating directory /foo/bar/baz");
    mkdir("/foo/bar/baz")
    fmt.Println()

    fmt.Println("Creating directory /foo/baz/bar");
    mkdir("/foo/baz/bar")  // This will fail
    fmt.Println()

    fmt.Println("Creating directory /foo/baz");
    mkdir("/foo/baz")
    fmt.Println()

    fmt.Println("Writing file /foo");
    write_file("/foo", "Hello World!"); // This will fail
    fmt.Println()

    fmt.Println("Writing file /bar");
    write_file("/bar", "Hello World!");
    fmt.Println()

    fmt.Printf("Reading file /bar: %v\n", read_file("/bar"));
    fmt.Println()

    fmt.Println("Writing file /foo/baz/bar");
    write_file("/foo/baz/bar", "Foo baz bar!");
    fmt.Println()

    fmt.Printf("Reading file /foo/baz/bar: %v\n", read_file("/foo/baz/bar"));
    fmt.Println()
}
