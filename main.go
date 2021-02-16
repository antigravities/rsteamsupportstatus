package main

import (
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	output := flag.String("output", "status.html", "Specify the output file")
	flag.Parse()

	res, err := http.Get("https://store.steampowered.com/stats/support")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	cats := make(map[string]string)

	doc.Find(".player_count_row").Each(func(idx int, sel *goquery.Selection) {
		cats[sel.Find(".player_count_row > * > .supportDetail.strong").First().Text()] = sel.Find("td:not([align='right']) > .supportDetail").First().Text()
	})

	details := []string{}

	for key, value := range cats {
		details = append(details, key+": "+value)
	}

	tmpl := `
<html>
	<head>
		<meta http-equiv="refresh" content="0; url=https://store.steampowered.com/stats/support">
		<script>window.location="https://store.steampowered.com/stats/support"</script>

		<meta property="og:title" content="Steam Support Status">
		<meta property="og:description" content="[[detail]]">
		<meta property="description" content="[[detail]]">
		<meta property="og:pubdate" content="[[date]]">
		<meta property="og:site_name" content="Steam">
		<meta name="theme-color" content="#33435c">

		<title>
			Redirecting...
		</title>
	</head>
	<body>
		Click <a href="https://store.steampowered.com/stats/support">here</a> if you aren't automatically redirected.
	</body>
</html>
`

	wr := strings.ReplaceAll(tmpl, "[[detail]]", strings.Join(details, "\n"))
	wr = strings.ReplaceAll(wr, "[[date]]", time.Now().Format(time.RFC3339))

	file, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write([]byte(wr))
}
