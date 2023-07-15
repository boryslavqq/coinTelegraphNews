package entity

import "time"

type Response struct {
	Data struct {
		Locale struct {
			Category struct {
				CacheKey string `json:"cacheKey"`
				Id       string `json:"id"`
				Posts    struct {
					Data []struct {
						CacheKey      string `json:"cacheKey"`
						Id            string `json:"id"`
						Slug          string `json:"slug"`
						Views         int    `json:"views"`
						PostTranslate struct {
							CacheKey             string    `json:"cacheKey"`
							Id                   string    `json:"id"`
							Title                string    `json:"title"`
							Avatar               string    `json:"avatar"`
							Published            time.Time `json:"published"`
							PublishedHumanFormat string    `json:"publishedHumanFormat"`
							LeadText             string    `json:"leadText"`
							Author               struct {
								CacheKey            string      `json:"cacheKey"`
								Id                  string      `json:"id"`
								Slug                string      `json:"slug"`
								InnovationCircleUrl interface{} `json:"innovationCircleUrl"`
								AuthorTranslates    []struct {
									CacheKey string `json:"cacheKey"`
									Id       string `json:"id"`
									Name     string `json:"name"`
									Typename string `json:"__typename"`
								} `json:"authorTranslates"`
								Typename string `json:"__typename"`
							} `json:"author"`
							Typename string `json:"__typename"`
						} `json:"postTranslate"`
						Category struct {
							CacheKey           string `json:"cacheKey"`
							Id                 string `json:"id"`
							Slug               string `json:"slug"`
							CategoryTranslates []struct {
								CacheKey string `json:"cacheKey"`
								Id       string `json:"id"`
								Title    string `json:"title"`
								Typename string `json:"__typename"`
							} `json:"categoryTranslates"`
							Typename string `json:"__typename"`
						} `json:"category"`
						Author struct {
							CacheKey         string `json:"cacheKey"`
							Id               string `json:"id"`
							Slug             string `json:"slug"`
							AuthorTranslates []struct {
								CacheKey string `json:"cacheKey"`
								Id       string `json:"id"`
								Name     string `json:"name"`
								Typename string `json:"__typename"`
							} `json:"authorTranslates"`
							Typename string `json:"__typename"`
						} `json:"author"`
						PostBadge struct {
							CacheKey            string `json:"cacheKey"`
							Id                  string `json:"id"`
							Label               string `json:"label"`
							PostBadgeTranslates []struct {
								CacheKey string `json:"cacheKey"`
								Id       string `json:"id"`
								Title    string `json:"title"`
								Typename string `json:"__typename"`
							} `json:"postBadgeTranslates"`
							Typename string `json:"__typename"`
						} `json:"postBadge"`
						ShowShares bool   `json:"showShares"`
						ShowStats  bool   `json:"showStats"`
						Typename   string `json:"__typename"`
					} `json:"data"`
					PostsCount int    `json:"postsCount"`
					Typename   string `json:"__typename"`
				} `json:"posts"`
				Typename string `json:"__typename"`
			} `json:"category"`
			Typename string `json:"__typename"`
		} `json:"locale"`
	} `json:"data"`
	Extensions struct {
		State string `json:"state"`
	} `json:"extensions"`
}
