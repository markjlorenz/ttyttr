## Getting a `Bearer` token:

```
POST /oauth2/token HTTP/1.1
Host: api.twitter.com
User-Agent: My Twitter App v1.0.23
Authorization: Basic <base64encode(consumer_key:consumer_secret)>
Content-Type: application/x-www-form-urlencoded;charset=UTF-8
Content-Length: 29

grant_type=client_credentials
```
