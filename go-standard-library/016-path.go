package main 

import (
    "path"
    "path/filepath"
    "fmt"
)

func main() {
    fmt.Println(path.Dir("golang/main/module.zig"))
    fmt.Println(path.Base("golang/main/module.zig"))
    fmt.Println(path.Ext("golang/main/module.zig"))
    fmt.Println(path.Join("golang", "main", "module.zig"))
    // FILE PATH
    fmt.Println(filepath.Dir("golang/main/module.zig"))
    fmt.Println(filepath.Base("golang/main/module.zig"))
    fmt.Println(filepath.Ext("golang/main/module.zig"))
    fmt.Println(filepath.IsAbs("golang/main/module.zig"))
    fmt.Println(filepath.IsLocal("golang/main/module.zig"))
    fmt.Println(filepath.Join("golang", "main", "module.go"))
}