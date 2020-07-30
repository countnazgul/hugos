package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

func basic() string {
  return "This is A Message"
}

func main() {

  js := mewn.String("./frontend/public/bundle.js")
  css := mewn.String("./frontend/public/utils.css")

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "Hugos CMS - Desktop App",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(basic)
  app.Run()
}
