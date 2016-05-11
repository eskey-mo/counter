# counter

このアプリはGo言語で作られたカウンターアプリです

コンソール上で動作し以下の機能を持ちます。

* カウントアップ・ダウン・リセット
* カウンターの名前付け・一覧表示

# コマンド

コマンド入力時の動作一覧

``` sh
# +1(default)
$ counter up
1

# -1(default)
$ counter down
0

# name option
$ counter up hoge
1
$ counter up hoge2
1

# name list
$ counter list
default 0
hoge 1
hoge2 1

# reset
$ counter reset hoge2
0

# delete
$ counter delete hoge2
$ counter list
default 0
hoge 1

```

# カウンターのデータ型

カウンターの値はJSON型で保管してあります。
{
    counter: [
        {
          name: "default",
          num: "0"     
        },
        {
          name: "hoge",
          num: "1"     
        },
        {
          name: "hoge2",
          num: "1"
        }

    ]
}

# 注意事項

名前を付けない場合、defaultとなります。  

