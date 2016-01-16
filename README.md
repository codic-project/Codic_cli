# Codic_cli
[Codic](https://codic.jp/) is naming tool for programmers.

## install
To install, use `go get:`
```
$ go get github.com/codic-project/Codic_cli
```

## how to use 
### Set AccessToken.
Get AccessToken from your [account page](https://codic.jp/my/api_status).

```
$ codic_tool -token=XXXXXXXXXXX 東京で狩りをする                                                
[ 東京で狩りをする ]=> do_hunting_in_tokyo
```

Token set is only once, the token cashe to `/tmp/token_codic`.

### Set code casing.
You can set casing type as following.

```
$ codic_tool -casing=hyphen こんにちは世界               
[ こんにちは世界 ]=> hello-world
```

| camel | pascal | lower_underscore|upper_underscore|hyphen|
|:-----------|:----------|:----------|:----------|:----------|
| HelloWorld|HelloWorld |hello_world|HELLO_WORLD|hello-world|

casing also cache, so you can run same before.

```
$ codic_tool こんにちは世界               
[ こんにちは世界 ]=> hello-world
```
