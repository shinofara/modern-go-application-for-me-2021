# 2021年版 新しくWEBサービスを開発・運用するとしたらGoのリポジトリをどうするか考えてみた

当リポジトリは個人的に今0からGoで開発するとしたら、どんな構成になるだろうかとふと思って開発した物になります。
もしかしたら数日後には違うなと思うかもしれないですが、一旦整理していきます。

## 例題

よくあるシンプルなTODO管理APIをOpen APIという形で提供します。

## 前提となるキーワード

- スキーマ駆動でジェネレート
- DBスキーマとコードにずれを起こさせない。
- コマンドクエリの関心分離

## 設計

### Open API

REST APIの開発には、[OpenAPI 3.0 ](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md)に準拠する形で開発を進めるため、[deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)を利用しました。
[openapi.yaml](./openapi.yaml)でスキーマ管理を進めて、必要なコードをジェネレートしています。

ジェネレートされたインターフェースや、リクエストオブジェクト等は、[http/oapi](http/oapi)に書き出されます。

### HTTP

[deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)では、[labstack/echo](https://github.com/labstack/echo)というフレームワーク用にジェネレートされます。
今回の例では、よりGoらしくしていく為に、フルスタックフレームワークと呼ばれる物は利用せず必要に応じて拡張していきたい為、[go-chi/chi](https://github.com/go-chi/chi)という`net/http` インターフェースに準拠した物を採用しています。
それに合わせて、`oapi-codegen` 実行オプションで `-generate chi-server` を指定して、`chi` 用のインターフェースをジェネレートしました。

そのため、[http/handler](http/handler)には、oapi-codegenで作成したインターフェースの実装を、`net/http` 用のハンドラーとして作成してます。
またhandlerの実装方針としては、ormからシンプルに取得してくるだけなどで、あればその他に依存させる事なく、直接ormを利用しています。
つまりhandlerはusecaseにもrepositoryにも依存しないで、DBからデータを取得できる状態です。

### UseCase

[UseCase](./usecase)では特定の処理を行う際に複数のdomain操作や、インフラ操作が発生するときに利用
例えばユーザ登録というケースで、認証情報登録、ユーザ情報登録、登録完了メール送信とDB操作とメール送信など一つのケースに複数の処理が含まれる場合に利用しています。

### Repository

### Infrastructure

### Config

## スキーマを追加

```
go run entgo.io/ent/cmd/ent init User
```

## モデルを作成

```
make model
```

## マイグレーション

```
make migrate
```

## マイグレーションの流れ

```
// env/schame/user.go
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
+		field.String("aaaaaa").
+			Default("unknown"),
	}
}
```

```
make model
```

```
make migrate
```

```
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `age` bigint(20) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown',
  `aaaaaa` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```


```
// env/schame/user.go
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
-		field.String("aaaaaa").
-			Default("unknown"),
	}
}
```

```
make model
```

```
make migrate
```

```
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `age` bigint(20) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```