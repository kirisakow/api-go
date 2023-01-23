package main

import (
	"net/http"

	"kirisakow/url_tools/url_cleaner"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// test:
	// curl http://localhost:3000/url_clean?url=https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534
	router.GET("/url_clean", func(c *gin.Context) {
		url_to_clean := c.Query("url")
		clean_url := url_cleaner.Clean_url_from_unwanted_query_params(url_to_clean)
		c.String(http.StatusOK, clean_url+"\n")
	})
	// test:
	// curl http://localhost:3000/url_clean/https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534
	router.GET("/url_clean/*url", func(c *gin.Context) {
		url_to_clean := c.Param("url")[1:]
		clean_url := url_cleaner.Clean_url_from_unwanted_query_params(url_to_clean)
		c.String(http.StatusOK, clean_url+"\n")
	})

	router.Run("127.0.0.1:3000")
}
