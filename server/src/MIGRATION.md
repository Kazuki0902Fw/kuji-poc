# マイグレーションガイド

このプロジェクトでは `sql-migrate` を使用してデータベースマイグレーションを管理しています。

## セットアップ

依存関係は既に `go.mod` に含まれています。初回セットアップ時は以下を実行してください：

```bash
cd server/app
go mod download
```

## マイグレーションファイルの作成

マイグレーションファイルは `migrations/` ディレクトリに配置します。

ファイル名の形式: `{番号}_{説明}.{up|down}.sql`

例:
- `000001_initial_schema.up.sql` - マイグレーション適用時
- `000001_initial_schema.down.sql` - マイグレーションロールバック時

## マイグレーションの実行

### Dockerコンテナ内で実行（推奨）

コンテナ内でsql-migrate CLIツールを直接使用できます：

#### 方法1: シェルスクリプトを使用（環境変数対応）

```bash
# マイグレーションを適用
docker-compose exec app sh /app/migrate.sh production up

# マイグレーションをロールバック
docker-compose exec app sh /app/migrate.sh production down

# マイグレーション状態を確認
docker-compose exec app sh /app/migrate.sh production status
```

#### 方法2: sql-migrateコマンドを直接使用

環境変数が設定されている場合：

```bash
# マイグレーションを適用
docker-compose exec app sql-migrate up -config=/app/dbconfig.yml -env=production

# マイグレーションをロールバック
docker-compose exec app sql-migrate down -config=/app/dbconfig.yml -env=production

# マイグレーション状態を確認
docker-compose exec app sql-migrate status -config=/app/dbconfig.yml -env=production
```

### ローカル環境（development）

```bash
cd server/app
sql-migrate up -env=development
```

### 本番環境（production）

```bash
cd server/app
sql-migrate up -env=production
```

### その他のコマンド

```bash
# 特定の数のマイグレーションを適用
sql-migrate up -env=production -limit=1

# 特定の数のマイグレーションをロールバック
sql-migrate down -env=production -limit=1

# マイグレーション状態を確認
sql-migrate status -env=production
```

## マイグレーション状態の確認

```bash
# Dockerコンテナ内で実行
docker-compose exec app sql-migrate status -env=production

# ローカルで実行
cd server/app
sql-migrate status -env=development
```

## 環境変数

マイグレーション実行時に以下の環境変数が必要です：

- `DB_USER`: データベースユーザー名（デフォルト: root）
- `DB_PASSWORD`: データベースパスワード
- `DB_NAME`: データベース名（デフォルト: kujicole）
- `DB_HOST`: データベースホスト（development: localhost, production: db）

環境変数は `env/app.env` ファイルから読み込まれます。

## マイグレーションファイルの例

### Up マイグレーション（適用）

```sql
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### Down マイグレーション（ロールバック）

```sql
-- +migrate Down
DROP TABLE IF EXISTS users;
```

## 注意事項

- マイグレーションファイルの先頭には `-- +migrate Up` または `-- +migrate Down` を必ず記述してください
- ファイル名の番号は連番で、重複しないようにしてください
- 本番環境でマイグレーションを実行する前に、必ずバックアップを取得してください

