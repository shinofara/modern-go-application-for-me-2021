# 2021年版 新しくWEBサービスを開発・運用するとしたらGoのリポジトリをどうするか考えてみた

当リポジトリは個人的に今0からGoで開発するとしたら、どんな構成になるだろうかとふと思って開発した物になります。
もしかしたら数日後には違うなと思うかもしれないですが、一旦整理していきます。

## 例題

よくあるシンプルなTODO管理APIをOpen APIという形で提供します。

## 前提となるキーワード

- スキーマ駆動でジェネレート
- DBスキーマとコードにずれを起こさせない。
- コマンドクエリの関心分離
  - [参考記事1](https://little-hands.hatenablog.com/entry/2019/12/02/cqrs#DDD%E3%81%AE%E5%8F%82%E7%85%A7%E7%B3%BB%E5%87%A6%E7%90%86%E3%81%A7%E7%99%BA%E7%94%9F%E3%81%99%E3%82%8B%E8%AA%B2%E9%A1%8C)

## 設計

### レイヤーの考え方

当リポジトリでのレイヤー設計は大きく分けると下記の通りとなります。
DDDとかClean Architectureで考えられているArchitectureを参考にして、最低限これくらいかなという分け方にしてます。

| レイヤ | 責務 | 該当ディレクトリ |
| ---------------- |----------------------------------------------------------------- | ------------------------------------ |
| Application      | ユーザとの接点やアプリケーションとして行いたい振る舞いに関心を持つ          | - http<br> - usecase<br>  - openapi  |
| Domain           | Applicationが何であるかを意思する事無く、ドメイン処理にただただ関心を持つ  | - repository<br> - ent               | 
| Infrastructure   | DBや外部サービスなどApplicationやDomain処理が関係する世界との関係を作る   | - infrastructure                     |
| Config           | 当リポジトリ内で環境毎に変化させたい設定値などを管理                      | - config<br> - environment           |

### 各ディレクトリの説明

| ディレクトリ      | 責務                                                                                                                                                 | 依存                        |
| --------------- |---------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------- |
| openapi/src     | [OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md)に準拠するREST APIのスキーマ管理                                 | -                          |
| openapi         | [openapi.yml](openapi/src/openapi.yaml)を元に[deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)で生成された `*.gen.go` ファイルが置かれる場所 | -                          |
| http            | HTTP Applicationに関する関心事を管理                                                                                                                    |                            | 
| http/handler    | HTTP Applicationのroute毎に呼び出されるhandler処理を管理。handlerはrequestを受け取り、そしてresponseを返すことが仕事。                                          | usecase, repository        |
| usecase         | usecaseにはアプリケーションのケース単位で処理を集約します                                                                                                    | repository, infrastructure |
| repository      | 当リポジトリ内で環境毎に変化させたい設定値などを管理                                                                                                         | infrastructure              |
| ent             | DBスキーマに関する管理と操作、そしてその結果ORMを提供します                                                                                                  |
| infrastructure  | 外部とのコミュニケーションを行う為に利用。可能な限りinterface化して、テストなどではmockを利用できるように                                                           |
| config          | configはApplication/Infrastructureレイヤで発生する環境毎の値や秘密情報などを管理                                                                             |
| environment     | 当リポジトリ内で環境毎に変化させたい設定値などを管理。environmentで使用できる値は、configで設定してるものに限る                                                     |

#### Open API

REST APIの開発には、[OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md)に準拠する形で開発を進めるため、[deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)を利用しました。
[openapi.yaml](./openapi/src/openapi.yaml)でスキーマ管理を進めて、必要なコードをジェネレートしています。

ジェネレートされたインターフェースや、リクエストオブジェクト等は、[openapi](./openapi)に書き出されます。

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

シンプルに外側のclient実装ですね。DBとかTraceとかLoggerとか

### Config

設定値などは、Envに持たせるパターンもあると思いますが、今回は[config.yml](environment/development/config.yml)に必要な設定をもたせる形にしました。
理由は下記の通り

1. Envの場合はどこからでも参照できてしまうリスクがある。
   1. もちろん運用でカバーは可能
   2. もちろんそれがOKという考えもあるが、ここでは必要な設定が無いとむしろpanicさせたいので、起動時に評価したい考え
2. 全てがStringになってしまう。もちろん `os.Getenv` で取得後キャストするのもあり

今回は秘密情報を管理はしていないですが、秘密情報もsecret.ymlなどyaml形式で渡せるようになります。
GCPであればSecret Managerで管理して、起動時にマウントするなどします。

## Usage

### 環境構築

```shell
$ make init
```

[Makefile](./Makefile)のinitに起動時に行う処理を全て書いてます。

### モデル関連

[facebook/ent](https://github.com/ent/ent)を利用していますので、マイグレーションの為に別のファイルが存在するなどは無いです。
entの使い方は別途どこかのドキュメントリンクを記述します。

#### モデルの準備

モデルの追加は下記のコマンドを実行する事で作成できます。

```
go run entgo.io/ent/cmd/ent@latest init User
```

#### モデルを元にORM作成とマイグレーションファイルの更新

```
make model
```

### マイグレーション実行

```
make migrate
```

### マイグレーションの流れ

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