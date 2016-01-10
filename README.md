# Codic_cli
[Codic](https://codic.jp/)のAPIをCLIから叩くツールです。

## Codicとは？
いい感じに日本語から関数名とか変数名を作ってくれるサービスです。
https://codic.jp/
APIも公開してます。
https://codic.jp/docs/api

## install
落としてきてgo installすればいいと思います。

## 使い方
[ここ](https://codic.jp/my/api_status)から自分のAccessTokenを落としてきて使ってください
AccessTokenがコピペできないので泣きの目Copyでお願いします。

```
$ codic_tool -token=XXXXXXXXXXX 東京で狩りをする                                                
[ 東京で狩りをする ]=> do hunting in tokyo
```

-token=のオプションは一回入れればいいです。

またcasingも指定できます

```
$ codic_tool --help
  -casing string
    	[camel, pascal, lower_underscore, upper_underscore, hyphen] (default "camel")
  -token string
    	initial setting token. (default "default")
$ codic_tool --casing=lower_underscore こんにちは世界
[ こんにちは世界 ]=> HELLO_WORLD
```
