[Goで学ぶGraphQLサーバーサイド入門](https://zenn.dev/hsaki/books/golang-graphql)

## 利用ライブラリ
- HTTP Server: net/http
- GraphQL Server: gqlgen
- Test: testing

## ディレクトリ構造
- graph/
  - model/: エンティティの型定義
  - services/: ビジネスロジック
    - *_test.go: テストコード
  - schema.resolvers.go: リゾルバ
  - schema.resolvers_test.go: リゾルバのテストコード
  - generated.go: 自動生成。クエリとリゾルバのマッピング
