package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: go-rm-dist [path]")
    return
  }
  path := os.Args[1]
  fmt.Println("Path:", path)
  deletedFileName := "dist"
  deleteBuild(path, deletedFileName)
}

func deleteBuild(path string, deletedFileName string) {
  err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
    stat, _ := os.Stat(path)
    if stat == nil {
      return nil
    }
    if err != nil {
      fmt.Println("error", err)
    }
    if info.IsDir() {
      folderName := info.Name()

      if folderName == deletedFileName {
        err := os.RemoveAll(path)
        if err != nil {
          fmt.Println("delete faild", path)
        } else {
          fmt.Println("deleted", path)
          return nil
        }
      }
    }
    return nil
  })

  if err != nil {
    fmt.Println("Error:", err)
  }

  fmt.Println("Done")
}
