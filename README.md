### This is a REST API implementation of a Go function which cleans any given URL of junk query parameters

* Core dependency: [`"kirisakow/url_tools/url_cleaner"`](https://github.com/kirisakow/url_tools#go-function-url_clean)
* Framework: [gin](https://github.com/gin-gonic/)

## Examples

* Request: send any URL full of junk query parameters:
  * either as a query param: https://crac.ovh/url_clean?url=https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534
  * or as a path param: https://crac.ovh/url_clean/https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534
* Response: a clean URL:
  * https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else