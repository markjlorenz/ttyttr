## Usage
```
markjlorenz@BabyGoat$ ./bin/ttyttr 434135040256008192
```

```
⌈--------------------------------------------------⌉
|Gary Bernhardt                              14-Feb|
|@garybernhardt                                    |
|                                                  |
|Reminder to never use Array#[] or Hash#[] in Ruby.|
|  I mean this literally: never. If you want nil,  |
|       explicitly say `x.fetch(i) { nil }`.       |
|                                                  |
|/garybernhardt/status/434135040256008192          |
⌊--------------------------------------------------⌋
```
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

## Tests

`make test`

Should output:
```
# LONG TWEET
go run ttyttr.go 444239051965882368
.--------------------------------------------------.
|static recycling                            13-Mar|
|@jennschiffer                                     |
|                                                  |
|  i regularly search "cool-code-pal" here to see  |
|   people responding to css perverts posts. well  |
|  apparently I've angered a bunch of germans lol  |
|                                                  |
|/jennschiffer/status/444239051965882368           |
*--------------------------------------------------*
# TWEET WITH NEW LINES
go run ttyttr.go 442781697596092418
.--------------------------------------------------.
|Dorian S. Nakamoto                          09-Mar|
|@DorianSatoshi                                    |
|                                                  |
|         1. Invent Bitcoin under my name          |
|  2. Take very careful steps to remain anonymous  |
|                      3. ???                      |
|                  4. FREE LUNCH                   |
|                                                  |
|/DorianSatoshi/status/442781697596092418          |
*--------------------------------------------------*
# TWEET WITH TOO-LONG WORD
go run ttyttr.go 415804923591147520
.--------------------------------------------------.
|Akshar                                      25-Dec|
|@AksharPathak                                     |
|                                                  |
|  The longest word in the English dictionary is   |
| mutualfundsaresubjecttomarketriskpleasereadtheoff|
|       erdocumentcarefullybeforeinvesting.        |
|                                                  |
|/AksharPathak/status/415804923591147520           |
*--------------------------------------------------*
# I MADE THIS APP SO I COULD PUT THIS TWEET IN A COMMIT MESSAGE:
go run ttyttr.go 434135040256008192
.--------------------------------------------------.
|Gary Bernhardt                              14-Feb|
|@garybernhardt                                    |
|                                                  |
|Reminder to never use Array#[] or Hash#[] in Ruby.|
|  I mean this literally: never. If you want nil,  |
|       explicitly say `x.fetch(i) { nil }`.       |
|                                                  |
|/garybernhardt/status/434135040256008192          |
*--------------------------------------------------*
```
