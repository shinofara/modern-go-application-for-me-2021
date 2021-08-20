
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