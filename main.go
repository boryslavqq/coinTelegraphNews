package main

import (
	"cointelegraphLatestNews/entity"
	"encoding/json"
	"fmt"
	"github.com/subosito/gotenv"
	"github.com/valyala/fasthttp"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"strings"
	"time"
)

const breakingTag = "Breaking news"

const (
	url     = "https://conpletus.cointelegraph.com/v1/"
	slugUrl = "https://cointelegraph.com/news/"
)

func newBot(token string) *tele.Bot {
	pref := tele.Settings{
		Token:     token,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeHTML,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	return b

}

func generateRequest() *fasthttp.Request {
	req := fasthttp.AcquireRequest()

	body := map[string]interface{}{
		"query":         "query CategoryPagePostsQuery($short: String, $slug: String!, $offset: Int = 0, $length: Int = 10, $hideFromMainPage: Boolean = null) {\n  locale(short: $short) {\n    category(slug: $slug) {\n      cacheKey\n      id\n      posts(\n        order: \n\"postPublishedTime\"\n        offset: $offset\n        length: $length\n        hideFromMainPage: $hideFromMainPage\n      ) {\n        data {\n          cacheKey\n          id\n          slug\n          views\n          postTranslate {\n            cacheKey\n            id\n            title\n            avatar\n            published\n            publishedHumanFormat\n            leadText\n            author {\n              cacheKey\n              id\n              slug\n              innovationCircleUrl\n              authorTranslates {\n                cacheKey\n                id\n                name\n                __typename\n              }\n              __typename\n            }\n            __typename\n          }\n          category {\n            cacheKey\n            id\n            slug\n            categoryTranslates {\n              cacheKey\n              id\n              title\n              __typename\n            }\n            __typename\n          }\n          author {\n            cacheKey\n            id\n            slug\n            authorTranslates {\n              cacheKey\n              id\n              name\n              __typename\n            }\n            __typename\n          }\n          postBadge {\n            cacheKey\n            id\n            label\n            postBadgeTranslates {\n              cacheKey\n              id\n              title\n              __typename\n            }\n            __typename\n          }\n          showShares\n          showStats\n          __typename\n        }\n        postsCount\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}",
		"operationName": "CategoryPagePostsQuery",
		"variables": map[string]interface{}{
			"slug":             "latest-news",
			"offset":           0,
			"length":           1,
			"hideFromMainPage": false,
			"short":            "en",
			"cacheTimeInMS":    300000,
		},
	}

	reqBody, _ := json.Marshal(body)

	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBody(reqBody)
	req.Header.Set("Accept", "application/graphql+json, application/json")
	req.Header.SetContentType("application/json")

	req.Header.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	return req

}

func main() {

	if err := gotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var channelID int64 = -1001894461736
	_unique := make(map[string]bool)
	client := fasthttp.Client{}
	bot := newBot(os.Getenv("TOKEN"))

	req := generateRequest()

	options := &tele.SendOptions{ParseMode: tele.ModeHTML}

	for {
		time.Sleep(time.Second * 3)
		res := fasthttp.AcquireResponse()
		breaking := false

		var body entity.Response

		err := client.Do(req, res)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(res.Body(), &body)
		if err != nil {
			log.Fatal(err)
		}

		post := body.Data.Locale.Category.Posts.Data[0]

		if _unique[post.Slug] {
			continue
		}

		alreadyViews := post.Views
		avatar := post.PostTranslate.Avatar
		title := post.PostTranslate.Title
		leadText := post.PostTranslate.LeadText

		tags := ""

		for _, category := range post.PostBadge.PostBadgeTranslates {
			tags += "#" + category.Title + " "
			if strings.Contains(category.Title, breakingTag) {
				breaking = true
			}
		}

		newsUrl := slugUrl + post.Slug

		linkToPost := fmt.Sprintf(`<a href="%s">ORIGINAL POST</a>`, newsUrl)

		text := fmt.Sprintf("<b>%s</b>\n\n%s\n\n%s%s%v views 👀\n\n%s", title, leadText, linkToPost, strings.Repeat(" ", 65), alreadyViews, tags)
		if breaking {
			text = "🚨🚨🚨🚨🚨🚨\n\n" + text
		}

		p := &tele.Photo{File: tele.FromURL(avatar), Caption: text}

		_, err = bot.SendAlbum(&tele.User{ID: channelID}, tele.Album{p}, options)
		if err != nil {
			log.Fatal(err)
			return
		}
		_unique[post.Slug] = true

	}

}
