package convert

import (
	"context"

	"github.com/IsiahMc/HTML2PDF/internal/data"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// needs to take in params
func Convert(ctx context.Context, html string, options data.Conversion) ([]byte, error) {
	// create ctx
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	var buf []byte

	// create temp file and url var
	var source string
	if options.URL != "" {
		source = options.URL
	} else {
		source = "data:text/html," + html
	}

	// navigate to html/url and convert
	tasks := chromedp.Tasks{
		chromedp.Navigate(source),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			//print page
			printOptions := page.PrintToPDF()
			if options.PageHeight > 0 {
				printOptions.WithPaperHeight(options.PageHeight)
			}
			if options.PageWidth > 0 {
				printOptions.WithPaperWidth(options.PageWidth)
			}
			buf, _, err = printOptions.Do(ctx)
			return err
		}),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		return nil, err
	}

	return buf, nil
}
