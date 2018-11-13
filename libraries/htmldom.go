package libraries

/*
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	xhtml "golang.org/x/net/html"
)

type Embeded struct {
	Type string `json:"type"`
	HTML string `json:"html"`
}

func MediaManipulation(page string) (string, error) {
	layout, err := xhtml.Parse(strings.NewReader(page))
	if err != nil {
		return "", err
	}

	var newLayout *xhtml.Node
	newLayout, err = HTMLDOM(layout)
	if err != nil {
		return "", err
	}

	newLayoutString := renderNode(newLayout)
	newLayoutString = strings.Replace(newLayoutString, "<html>", "", -1)
	newLayoutString = strings.Replace(newLayoutString, "</html>", "", -1)
	newLayoutString = strings.Replace(newLayoutString, "<head>", "", -1)
	newLayoutString = strings.Replace(newLayoutString, "</head>", "", -1)
	newLayoutString = strings.Replace(newLayoutString, "<body>", "", -1)
	newLayoutString = strings.Replace(newLayoutString, "</body>", "", -1)

	newLayoutString = cleanSpecialCharacter(newLayoutString)

	return newLayoutString, nil
}

func HTMLDOM(node *xhtml.Node) (*xhtml.Node, error) {
	var doc *xhtml.Node

	if node.Type == xhtml.ElementNode && node.Data == "img" {
		for i, img := range node.Attr {
			if img.Key == "src" {
				wp2017 := "https://theshonet.com/wp-content/uploads/2017/"
				wp2018 := "https://theshonet.com/wp-content/uploads/2018/"

				if strings.Contains(img.Val, wp2017) || strings.Contains(img.Val, wp2018) {
					path := strings.Split(img.Val, "/")
					filename := path[cap(path)-1]
					path[2] = "theshonet-assets.s3.ap-southeast-1.amazonaws.com"
					node.Attr[i] = xhtml.Attribute{}
					node.Attr = append(node.Attr, xhtml.Attribute{Key: "src", Val: strings.Join(path, "/")})

					path = path[:cap(path)-1]
					newPath := strings.Join(path, "/")

					node.Attr = append(node.Attr, xhtml.Attribute{Key: "srcset", Val: newPath + "/w320/" + filename + " 320w, " + newPath + "/w480/" + filename + " 480w, " + newPath + "/w800/" + filename + " 800w"})
					node.Attr = append(node.Attr, xhtml.Attribute{Key: "sizes", Val: "(max-width: 320px) 280px, (max-width: 480px) 440px, 800px"})
				}
			}
		}
	}

	if node.Type == xhtml.ElementNode && node.Data == "p" {
		for _, p := range node.Attr {
			if !strings.Contains(p.Val, "text-align:") {
				node.Attr = []xhtml.Attribute{}
			}
		}

		if node.FirstChild != nil && node.FirstChild.Type == xhtml.TextNode {
			if strings.Contains(node.FirstChild.Data, "https://www.instagram.com") {
				embed, err := instagramEmbed(node.FirstChild.Data)
				if err != nil {
					return doc, err
				}

				node.FirstChild.Data = ""
				node.AppendChild(embed)
			}

			if strings.Contains(node.FirstChild.Data, "https://twitter.com") {
				embed, err := twitterEmbed(node.FirstChild.Data)
				if err != nil {
					return doc, err
				}

				node.FirstChild.Data = ""
				node.AppendChild(embed)
			}

			if strings.Contains(node.FirstChild.Data, "https://www.youtube.com") || strings.Contains(node.FirstChild.Data, "https://youtu.be/") {
				embed, err := youtubeEmbed(node.FirstChild.Data)
				if err != nil {
					return doc, err
				}

				node.FirstChild.Data = ""
				node.AppendChild(embed)
			}
		}
	}

	if node.Type == xhtml.ElementNode && node.Data == "div" {
		for _, div := range node.Attr {
			if strings.Contains(div.Key, "style") {
				node.Attr = []xhtml.Attribute{}
			}
		}

		if node.FirstChild != nil && node.FirstChild.Type == xhtml.TextNode {
			if strings.Contains(node.FirstChild.Data, "https://www.instagram.com") {
				embed, err := instagramEmbed(node.FirstChild.Data)
				if err != nil {
					return doc, err
				}

				node.FirstChild.Data = ""
				node.AppendChild(embed)
			}

			if strings.Contains(node.FirstChild.Data, "https://twitter.com") {
				embed, err := twitterEmbed(node.FirstChild.Data)
				if err != nil {
					return doc, err
				}

				node.FirstChild.Data = ""
				node.AppendChild(embed)
			}

			if strings.Contains(node.FirstChild.Data, "https://www.youtube.com") || strings.Contains(node.FirstChild.Data, "https://youtu.be") {
				embed, err := youtubeEmbed(node.FirstChild.Data)
				if err != nil {
					return doc, err
				}

				node.FirstChild.Data = ""
				node.AppendChild(embed)
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		HTMLDOM(child)
	}

	doc = node
	return doc, nil
}

func instagramEmbed(url string) (*xhtml.Node, error) {
	var instaDoc *xhtml.Node
	instaHTML := fmt.Sprintf("https://api.instagram.com/oembed?url=%s", url)

	req, err := http.NewRequest("GET", instaHTML, nil)
	if err != nil {
		return instaDoc, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return instaDoc, err
	}

	defer resp.Body.Close()

	var instagram Embeded

	if err := json.NewDecoder(resp.Body).Decode(&instagram); err != nil {
		return instaDoc, err
	}

	instaDoc, err = xhtml.Parse(strings.NewReader(instagram.HTML))
	if err != nil {
		return instaDoc, err
	}

	return instaDoc, nil
}

func twitterEmbed(url string) (*xhtml.Node, error) {
	var twitterDoc *xhtml.Node
	twitterURL := fmt.Sprintf("https://publish.twitter.com/oembed?url=%s", url)

	req, err := http.NewRequest("GET", twitterURL, nil)
	if err != nil {
		return twitterDoc, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return twitterDoc, err
	}

	defer resp.Body.Close()

	var twitter Embeded

	if err := json.NewDecoder(resp.Body).Decode(&twitter); err != nil {
		return twitterDoc, err
	}

	twitterDoc, err = xhtml.Parse(strings.NewReader(twitter.HTML))
	if err != nil {
		return twitterDoc, err
	}

	return twitterDoc, nil
}

func youtubeEmbed(uri string) (*xhtml.Node, error) {
	var youtubeDoc *xhtml.Node
	var id string

	xURL, err := url.Parse(uri)
	if err != nil {
		return youtubeDoc, err
	}

	if strings.Contains(uri, "https://www.youtube.com") {
		path, err := url.ParseQuery(xURL.RawQuery)
		if err != nil {
			return youtubeDoc, err
		}
		id = path.Get("v")
	}

	if strings.Contains(uri, "https://youtu.be/") {
		path, err := url.Parse(uri)
		if err != nil {
			return youtubeDoc, err
		}
		id = strings.Replace(path.EscapedPath(), "/", "", -1)
	}

	iframe := `<iframe id="ytplayer" type="text/html" width="100%" height="450px" src="https://www.youtube.com/embed/` + id + `?rel=0&showinfo=0&color=white&iv_load-policy=3" frameborder="0" allowfullscreen></iframe>`

	youtubeDoc, err = xhtml.Parse(strings.NewReader(iframe))
	if err != nil {
		return youtubeDoc, err
	}

	return youtubeDoc, nil
}

func renderNode(node *xhtml.Node) string {
	var buf bytes.Buffer

	write := io.Writer(&buf)
	xhtml.Render(write, node)

	return buf.String()
}

func cleanSpecialCharacter(page string) string {
	page = strings.Replace(page, "â&#128;&#152;", "'", -1)
	page = strings.Replace(page, "â&#128;&#153;", "'", -1)
	page = strings.Replace(page, "â&#128;&#156;", "'", -1)
	page = strings.Replace(page, "â&#128;&#157;", "'", -1)
	page = strings.Replace(page, "â&#128;&#147;", "-", -1)
	page = strings.Replace(page, "Â", "", -1)
	page = strings.Replace(page, "â", "", -1)

	page = strings.Replace(page, "&acirc;&#128;&#152;", "'", -1)
	page = strings.Replace(page, "&acirc;&#128;&#153;", "'", -1)
	page = strings.Replace(page, "&acirc;&#128;&#156;", "'", -1)
	page = strings.Replace(page, "&acirc;&#128;&#157;", "'", -1)
	page = strings.Replace(page, "&acirc;&#128;&#147;", "-", -1)
	page = strings.Replace(page, "&Acirc;", "", -1)
	page = strings.Replace(page, "&acirc;", "", -1)

	return page
}
*/
