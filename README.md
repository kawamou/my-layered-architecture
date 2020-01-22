# go-cleanarchitecture-restapi
Example of  REST API implemented in Go.
Go言語で実装したクリーンアーキテクチャのサンプルです。
# 使用法
実行
```
go run main.go
```
テスト
```
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"id": "1", "title": "hoge", "text":"hoge","date":"hoge"}' localhost:8888/create
```
# 参考
- https://github.com/so-heee/golang-clean-architecture-example
- https://qiita.com/so_heee_/items/0cca93008eae635c642a