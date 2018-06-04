## FSMODIFY

A Golang library to watch file changes.

### Motivation

The <a href="https://github.com/fsnotify/fsnotify">fsnotify</a> library it is amazing for watch files changes, but when using the program inside a docker container that maps a volume on the host, the files changes was not fired.

The event happens in docker host machine, but not in container.

0FSmodify was created to solve this issue, it will watch for modification time in file, this way works in docker containers using volumes mapped.
## Installation

go get github.com/rgobbo/fsmodify

## Usage

```go
import "github.com/rgobbo/fsmodify"

func main () {
    var i = 3  // 3 seconds - Time to refresh the file infos
	fmt.Println("Watching files")
	fsmodify.NewWatcher("./test",".txt", i, func(fileName string) {
		fmt.Println("Changed file->", fileName)
	})
}

```