# go-chi-gorm

## swagger の自動生成方法

[swag](https://github.com/swaggo/swag)を使って，コメントから swagger を自動生成しています．

```bash
~/go/bin/swag init --parseDependency --parseInternal
```

http://localhost:8080/swagger/index.html にアクセス

## シードデータの作成

```bash
go run seed/main.go
```
