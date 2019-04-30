package main

import (
 //  "fmt"
   "github.com/zserge/webview"
)

func main() {

   webView := webview.New(webview.Settings{
      Title:     "Wazaterm",
      URL:       "https://www.wazaterm.com",
      Width:     1000,
      Height:    800,
      Resizable: true,
      Debug:     true,
      ExternalInvokeCallback: nil,
   })

   webView.SetFullscreen(true)
   webView.Run()

 }
