package main

import (
  "io/ioutil"
  "log"
  "encoding/json"
)

type video struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  ImageUrl string `json:"imageurl"`
  Url string `json:"url"`
}

func getVideos() (videos []video) {
  byteContent, err := ioutil.ReadFile("videos.json")

  if err != nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(byteContent, &videos)

  if err != nil {
    log.Fatal(err)
  }

  return videos
}

func saveVideos(videos []video) {
  byteContent, err := json.Marshal(videos)

  if err != nil {
    log.Fatal(err)
  }
  
  err = ioutil.WriteFile("updated-videos.json", byteContent, 0644)

  if err != nil {
    log.Fatal(err)
  }
}
