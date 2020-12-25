package main
import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	var ua string
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://studygolang.com`),
		chromedp.WaitVisible(`#custom-ua-string`),
		chromedp.Text(`#custom-ua-string`, &ua),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user agent: %s", ua)
}

