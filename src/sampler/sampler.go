package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "os/user"
    "path/filepath"
    . "github.com/kkdai/youtube"
)

func main() {
    flag.Parse()

    if len(os.Args) == 1 {
        fmt.Println("input argument missing.\n")
        os.Exit(3) 
    }

    log.Println(flag.Args())
    
    usr, _     := user.Current()
    currentDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)

    log.Println("download to dir=", currentDir)
    y          := NewYoutube(true)
    arg        := flag.Arg(0)

    if err := y.DecodeURL(arg); err != nil { fmt.Println("err:", err) }
    if err := y.StartDownload(filepath.Join(currentDir, "dl.mp4")); err != nil {
        fmt.Println("err:", err)
    }
}
