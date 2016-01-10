# Codic_cli
[Codic](https://codic.jp/)のAPIをCLIから叩くツールです。

## install
落としてきてgo installすればいいと思います

## how to use
[ここ](https://codic.jp/my/api_status)から自分のAccessTokenを落としてきて使ってください
もうちょっとで完成します。

```
$ codic_tool -token=XXXXXXXXXXX 東京で狩りをする                                                
[ 東京で狩りをする ]=> do hunting in tokyo
```

-token=のオプションは一回入れればいいです。

またcasingも指定できます

```
$ codec_cli --help
  -casing string
    	[camel, pascal, lower_underscore, upper_underscore, hyphen] (default "camel")
  -token string
    	initial setting token. (default "default")
$ codec_cli --casing=lower_underscore こんにちは世界
[ こんにちは世界 ]=> HELLO_WORLD
```
