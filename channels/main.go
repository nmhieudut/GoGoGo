package main
import (
	"net/http"
	// "io"
	// "os"
	"time"
	"fmt"
)
func main() {
	links := []string{
		"http://youtube.com",
		"http://google.com",
		"http://facebook.com",
		"http://pornhub.com",
		"http://tutorsfinder.surge.sh",
	} 
	c := make(chan string)
	
	for _,link := range links {
		go checkLink(link,c)
	}
	for l := range c {
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}
	fmt.Println("Exit program")
}
func checkLink(link string, c chan string) {
	_,err := http.Get(link)
	if err != nil {
		// fmt.Println(link, "might me down !")
		c <- link
		return 
	}

	// fmt.Println(link,"is up !")
	c <- link 
}