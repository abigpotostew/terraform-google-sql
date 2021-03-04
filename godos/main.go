package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

const (
	appUrl = "https://myapp-dot-prysm-glzh-dc71.uk.r.appspot.com/"
	numReq = 8000
	repeat = 100
)
func main() {
	client := http.DefaultClient
	//urlParse,_ := url.Parse(appUrl)

	for i := 0; i< repeat; i++{
		i:=i // pin
		errWg  := &errgroup.Group{}
		for n := 0; n<numReq; n++{
			n := n // pin
			errWg.Go(func() error {
				res, err := client.Get(appUrl)
				if err!=nil {
					return fmt.Errorf("%d-%d failed: %v", i, n, err)
				}
				log.Printf("%d-%d returned %d", i, n, res.StatusCode)
				return nil
			})
		}
		if err:= errWg.Wait(); err!=nil{
			panic(err)
		}
	}
}
