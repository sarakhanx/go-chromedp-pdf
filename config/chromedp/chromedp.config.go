package chromedp

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func PdfInit() {
	ctx, cencel := chromedp.NewContext(context.Background())
	defer cencel()

	//capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPdf(`https://crm.need.co.th/pdfdestination/31`, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("hifolks.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The PDF was created !")
}

func printToPdf(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(1 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(false).
				WithPaperWidth(8.27).
				WithPaperHeight(11.69).
				WithMarginBottom(0.2).
				WithMarginTop(0.2).
				WithMarginLeft(0.2).
				WithMarginRight(0.2).
				Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
