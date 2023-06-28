package main

import "fmt"

func main() {
  videos := getVideos()
  
  for i, _ := range videos {
    videos[i].Description = videos[i].Description + " Like and subscribe"
  }
  
  fmt.Println(videos)

  saveVideos(videos)
}
