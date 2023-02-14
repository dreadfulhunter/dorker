# Search Engine Scraper
This is a simple command line tool to scrape subdomains from popular search engines like Google, Bing, DuckDuckGo, Yandex, Yahoo, and Baidu.

## Usage
bash
```
./search_engine_scraper SEARCH_ENGINE QUERY
SEARCH_ENGINE should be one of the following:
```

google
bing
duck
yandex
yahoo
baidu
QUERY is the string to search for.

Example
bash
```
./search_engine_scraper google "example query"
```
Available Search Engines
Google
Bing
DuckDuckGo
Yandex
Yahoo
Baidu
Output
The program will output a list of matched subdomains in the format:

```javascript
Copy code
Response from https://www.google.com/search?q=example+query :
Matched subdomains:
subdomain1
subdomain2
...
