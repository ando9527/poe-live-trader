# POE Live Trader
This is the tool similar to live search in web.
We will do this in game.
 
##  How To
###  Setup cookie
1. Brows https://www.pathofexile.com/trade/search/
2. Press F12
3. Get the cookie, value of POESESSID
4. Fill in .env

### Setup Filter
1. Search the item in 
2. Get the filter string "8GvyvVFV" in https://www.pathofexile.com/trade/search/Legion/8GvyvVFV
3. Fill in .env

### Graphql Query
```graphql
query  {
  ssid {
    Content
  }
}


mutation {
  createOrUpdateSSID(input: { Content: "123456789" }) {
    Content
  }
}
```