package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main()  {
	//创建 chrome实例
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	//创建超时时间
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	//navigate to page, wait for an element ,click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://golang.org/pkg/time/"),
		//等待footer element is visible
		chromedp.WaitVisible("body>footer"),
		chromedp.Click(`#pkg-examples > div`, chromedp.NodeVisible),
		chromedp.Value(`#example_After .play .input textarea`, &example),


	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example: \n%s",example)

}
