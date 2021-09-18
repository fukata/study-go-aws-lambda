# study-go-aws-lambda

これは勉強用のリポジトリです。

## ビルド

```shell
$ make build
```

## リリースビルド

```shell
$ make
```

## ローカルで動作確認する

```shell
$ sam local start-api
$ curl -XPOST 'http://127.0.0.1:3000/'
```

## 参照

- https://github.com/aws/aws-lambda-go
- https://future-architect.github.io/articles/20210602a/
- https://future-architect.github.io/articles/20200323/