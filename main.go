package main

import (
  "fmt"
  "os"
  "time"

  "github.com/urfave/cli"
  "github.com/kyokomi/emoji"
)

func main() {
  app := cli.NewApp()

  app.Name = "go-dinner"
  app.Usage = "Deliver dinner"
  app.Version = "0.0.1"

  app.Action = func (c *cli.Context) error {
    ch := make(chan int, 20)

    sushi := emoji.Sprint(":sushi:")
    pizza := emoji.Sprint(":pizza:")
    hamburger := emoji.Sprint(":hamburger:")

    go receive(sushi, ch)
    go receive(pizza, ch)
    go receive(hamburger, ch)

    i := 0
    for i < 10 {
      ch <- i
      i++
    }
    close(ch)

    time.Sleep(2 * time.Second)
    return nil
  }

  app.Run(os.Args)
}

func receive(name string, ch <-chan int) {
  for {
    _, ok := <-ch
    if ok == false {
      break
    }
    fmt.Println(name)
  }
}
