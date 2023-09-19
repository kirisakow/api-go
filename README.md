**WARNING** This library is still under development and intended for experimental purposes only.

### This is a REST API implementation of a Go function which cleans any given URL of junk query parameters

* Core dependency: [`"kirisakow/url_tools/url_cleaner"`](https://github.com/kirisakow/url_tools#go-function-url_clean)
* Framework: [gin](https://github.com/gin-gonic/)

### Examples (live)

* Request: send any URL full of junk query parameters: `https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534`
  * either as a query param: [`https://crac.ovh/url_clean?url=`https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534](https://crac.ovh/url_clean?url=https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534)
  * or as a path param: [`https://crac.ovh/url_clean/`https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534](https://crac.ovh/url_clean/https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else?utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=20230121&utm_content=ed-picks-article-link-6&etear=nl_special_6&utm_campaign=a.coronavirus-special-edition&utm_medium=email.internal-newsletter.np&utm_source=salesforce-marketing-cloud&utm_term=1/21/2023&utm_id=1456534)
* Response: a clean URL:
  * https://www.economist.com/christmas-specials/2020/12/19/awesome-weird-and-everything-else

### Installation on a latest Ubuntu Linux server

#### 1. Download, compile and run

```sh
git clone ...

cd api-go

go build .

# run as a transient systemd service
sudo systemd-run --uid=myuser --gid=myuser --same-dir --unit=api-go_$(date +%Y%m%d_%H%M%S)_myuser /home/myuser/api-go/url_clean --port $PORT_NUM_AS_ENV_VAR
```

#### 2. Configure Nginx as reverse proxy, typically:

```sh
server {
   server_name ${DOMAIN_NAME_AS_ENV_VAR};

   location ~ ^/(endpoint|another_endpoint|yet_another) {
       proxy_pass http://localhost:${PORT_NUM_AS_ENV_VAR};
   }

   #
   # ...TLS / SSL / HTTPS / LetsEncrypt / certbot stuff...
   #

}
server {
   #
   # ...TLS / SSL / HTTPS / LetsEncrypt / certbot stuff...
   #
}
```

#### 3. Test Nginx configuration and, if valid, restart Nginx service:

```sh
sudo nginx -t && sudo systemctl restart nginx.service
```
