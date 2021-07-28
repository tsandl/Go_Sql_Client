package main
import (
	"log"
)

func main() {
	cli := NewClient("root", "hadoop", "happy_mall", "tcp(localhost:3306)")
	if cli.Conn() != nil {
		log.Panic("failed to conn to mysql ")
	}
	cli.Run()
}