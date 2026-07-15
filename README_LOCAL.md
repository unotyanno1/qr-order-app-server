# ローカル環境での起動手順（Docker不使用）

このドキュメントでは、Dockerを使わずにローカル環境でGoサーバーとMySQLを起動する方法を説明します。

## 前提条件

- Go 1.24.0以上がインストールされていること
- MySQL 8.0以上がインストールされていること
- MySQLがローカルで起動していること

## セットアップ手順

### 1. MySQLのインストールと起動

#### Windowsの場合
1. [MySQL公式サイト](https://dev.mysql.com/downloads/mysql/)からMySQLをダウンロードしてインストール
2. MySQLサービスが起動していることを確認

**方法1: サービス管理ツールを使用（GUI）**
   - Windowsのスタートメニューから「サービス」を検索して開く
   - または、コマンドプロンプト（cmd.exe）から：
     ```cmd
     services.msc
     ```
   - サービス一覧で「MySQL」または「MySQL80」「MySQL57」などのサービス名を探し、状態が「実行中」になっているか確認してください。
   - **注意**: `services.msc`はbashシェル（Git Bashなど）からは実行できません。コマンドプロンプトまたはPowerShellから実行してください。

**方法2: コマンドラインで確認（推奨）**

   **コマンドプロンプト（cmd.exe）の場合：**
   ```cmd
   # MySQLサービスを検索
   sc query type= service | findstr /I "mysql"
   ```

   **PowerShellの場合：**
   ```powershell
   # MySQLサービスを検索
   Get-Service | Where-Object{$_.Name -match "mysql"}
   ```

   **bashシェル（Git Bashなど）の場合：**
   ```bash
   # Windowsコマンドを実行
   cmd.exe /c "sc query type= service | findstr /I mysql"
   ```
   または
   ```bash
   # PowerShellコマンドを実行
   powershell.exe -Command "Get-Service | Where-Object{$_.Name -match 'mysql'}"
   ```

**方法3: サービス名が分かったら状態を確認**
   サービス名が「MySQL80」の場合の例：

   **コマンドプロンプト：**
   ```cmd
   sc query MySQL80
   ```

   **PowerShell：**
   ```powershell
   Get-Service MySQL80
   ```

   **bashシェル：**
   ```bash
   cmd.exe /c "sc query MySQL80"
   ```

   `STATE`が`RUNNING`なら起動中、`STOPPED`なら停止中です。

**方法4: MySQLコマンドで直接確認**
   ```bash
   mysqladmin -u root -p ping
   ```
   `mysqld is alive`と表示されれば起動しています。
   （MySQLのbinディレクトリがPATHに追加されている必要があります）

**MySQLサービスを起動する場合**
   サービス名が「MySQL80」の場合の例：

   **コマンドプロンプト（管理者権限で実行）：**
   ```cmd
   net start MySQL80
   ```

   **PowerShell（管理者権限で実行）：**
   ```powershell
   Start-Service MySQL80
   ```

   **bashシェル（管理者権限のコマンドプロンプトから実行）：**
   ```bash
   cmd.exe /c "net start MySQL80"
   ```

#### macOSの場合
```bash
# Homebrewを使用する場合
brew install mysql
brew services start mysql
```

#### Linuxの場合
```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install mysql-server
sudo systemctl start mysql
sudo systemctl enable mysql
```

### 2. データベースの作成

MySQLに接続してデータベースを作成します：

#### 方法1: mysqlコマンドがPATHに追加されている場合

```bash
mysql -u root -p
```

#### 方法2: Git Bashからフルパスで実行する場合

MySQLのbinディレクトリがPATHに追加されていない場合は、フルパスで実行します：

```bash
# MySQL 8.0の場合の例
"/c/Program Files/MySQL/MySQL Server 8.0/bin/mysql.exe" -u root -p
```

または、パスにスペースが含まれている場合はエスケープします：

```bash
/c/Program\ Files/MySQL/MySQL\ Server\ 8.0/bin/mysql.exe -u root -p
```

**注意**: 
- Git Bashでは、Windowsのパス（`C:\Program Files\...`）をUnix形式（`/c/Program Files/...`）で記述します
- バージョンによってパスが異なる場合があります（例：`MySQL Server 8.0`、`MySQL Server 5.7`など）
- 実際のインストールパスを確認するには、Windowsエクスプローラーで `C:\Program Files\MySQL\` を確認してください

#### MySQLに接続できたら

MySQLプロンプトで以下を実行：

```sql
CREATE DATABASE IF NOT EXISTS qr_order_db;
EXIT;
```

### 3. 環境変数の設定（オプション）

デフォルト値を使用する場合は、この手順はスキップできます。

デフォルト値：
- `DB_HOST=localhost`
- `DB_PORT=3306`
- `DB_USER=root`
- `DB_PASSWORD=password`
- `DB_NAME=qr_order_db`

#### Windowsの場合（PowerShell）
```powershell
$env:DB_HOST="localhost"
$env:DB_PORT="3306"
$env:DB_USER="root"
$env:DB_PASSWORD="your_password"
$env:DB_NAME="qr_order_db"
```

#### Windowsの場合（コマンドプロンプト）
```cmd
set DB_HOST=localhost
set DB_PORT=3306
set DB_USER=root
set DB_PASSWORD=your_password
set DB_NAME=qr_order_db
```

#### macOS/Linuxの場合
```bash
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=your_password
export DB_NAME=qr_order_db
```

または、`.env`ファイルを作成して使用することもできます（その場合は環境変数読み込みライブラリが必要）。

### 4. Goの依存関係をインストール

プロジェクトのルートディレクトリで以下を実行：

```bash
go mod download
```

### 5. データベースマイグレーションの実行

マイグレーションを最新の状態に反映するには、以下のいずれかの方法を使用します：

#### 方法1: 専用コマンドを使用（推奨）

```bash
go run cmd/migrate/main.go
```

#### 方法2: サーバー起動時に自動実行

サーバーを起動すると、自動的にマイグレーションが実行されます：

```bash
go run main.go
```

#### 方法3: golang-migrate CLIツールを使用

`golang-migrate`のCLIツールをインストールして使用することもできます：

```bash
# Windows (Scoopを使用する場合)
scoop install migrate

# または、Goから直接インストール
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

インストール後、以下のコマンドでマイグレーションを実行：

```bash
# デフォルト設定を使用する場合
migrate -path migrations -database "mysql://root:password@tcp(localhost:3306)/qr_order_db" up

# 環境変数を使用する場合
migrate -path migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up
```

**注意**: 
- サーバー起動時（`go run main.go`）にも自動的にマイグレーションが実行されます
- マイグレーションは冪等性があるため、既に適用済みのマイグレーションはスキップされます
- 環境変数を設定している場合は、それらの値が使用されます

### 6. サーバーの起動

#### 通常の起動
```bash
go run main.go
```

#### Airを使用したホットリロード起動（推奨）
```bash
# Airがインストールされていない場合
go install github.com/air-verse/air@latest

# Airで起動
air -c .air.toml
```

### 7. 動作確認

サーバーが起動したら、以下のエンドポイントにアクセスして確認できます：

```bash
curl http://localhost:8080/qr_code/1
```

## トラブルシューティング

### MySQLに接続できない場合

#### 1. MySQLサービスが起動しているか確認

**Windowsの場合：**

   **コマンドプロンプト：**
   ```cmd
   # サービス一覧からMySQLを検索
   sc query type= service | findstr /I "mysql"
   
   # サービス名が分かったら状態確認（例：MySQL80）
   sc query MySQL80
   
   # 停止している場合は起動（管理者権限が必要）
   net start MySQL80
   ```

   **PowerShell：**
   ```powershell
   # サービス一覧からMySQLを検索
   Get-Service | Where-Object{$_.Name -match "mysql"}
   
   # サービス名が分かったら状態確認（例：MySQL80）
   Get-Service MySQL80
   
   # 停止している場合は起動（管理者権限が必要）
   Start-Service MySQL80
   ```

   **bashシェル（Git Bashなど）：**
   ```bash
   # サービス一覧からMySQLを検索
   cmd.exe /c "sc query type= service | findstr /I mysql"
   
   # サービス名が分かったら状態確認（例：MySQL80）
   cmd.exe /c "sc query MySQL80"
   
   # 停止している場合は起動（管理者権限が必要）
   cmd.exe /c "net start MySQL80"
   ```

**macOS/Linuxの場合：**
   ```bash
   sudo systemctl status mysql
   ```

#### 2. MySQLコマンドが使えない場合

`mysql`コマンドが見つからない場合は、MySQLのbinディレクトリがPATHに追加されていない可能性があります。

**Windowsの場合：**

##### 方法1: GUIでPATHに追加（推奨）

1. **MySQLのインストールパスを確認**
   - 通常は以下のいずれかです：
     - `C:\Program Files\MySQL\MySQL Server 8.0\bin`
     - `C:\Program Files\MySQL\MySQL Server 8.1\bin`
     - `C:\Program Files\MySQL\MySQL Server 5.7\bin`
   - Windowsエクスプローラーで `C:\Program Files\MySQL\` を開いて確認してください

2. **システム環境変数のPATHに追加**
   - Windowsキー + R を押して「ファイル名を指定して実行」を開く
   - `sysdm.cpl` と入力してEnter
   - 「詳細設定」タブをクリック
   - 「環境変数」ボタンをクリック
   - 「システム環境変数」セクションで「Path」を選択して「編集」をクリック
   - 「新規」をクリック
   - MySQLのbinディレクトリのパスを入力（例：`C:\Program Files\MySQL\MySQL Server 8.0\bin`）
   - 「OK」をクリックしてすべてのダイアログを閉じる
   - **新しいコマンドプロンプトまたはGit Bashを開き直す**（重要：既に開いているターミナルには反映されません）

##### 方法2: PowerShellでPATHに追加（管理者権限が必要）

```powershell
# 現在のシステムPATHを取得
$currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")

# MySQLのbinディレクトリのパス（実際のパスに置き換えてください）
$mysqlPath = "C:\Program Files\MySQL\MySQL Server 8.0\bin"

# PATHに追加（既に存在する場合は追加しない）
if ($currentPath -notlike "*$mysqlPath*") {
    $newPath = $currentPath + ";" + $mysqlPath
    [Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
    Write-Host "PATHに追加しました: $mysqlPath"
} else {
    Write-Host "既にPATHに追加されています: $mysqlPath"
}

# 新しいコマンドプロンプトまたはGit Bashを開き直してください
```

##### 方法3: 一時的にPATHに追加（現在のセッションのみ）

**コマンドプロンプト：**
```cmd
set PATH=%PATH%;C:\Program Files\MySQL\MySQL Server 8.0\bin
```

**PowerShell：**
```powershell
$env:PATH += ";C:\Program Files\MySQL\MySQL Server 8.0\bin"
```

**Git Bash：**
```bash
export PATH="$PATH:/c/Program Files/MySQL/MySQL Server 8.0/bin"
```

##### 確認方法

新しいターミナルを開いて以下を実行：

```bash
mysql --version
```

バージョン情報が表示されれば成功です。

##### フルパスで実行する場合（一時的な解決策）

PATHに追加する前に、フルパスで実行することもできます：

```cmd
"C:\Program Files\MySQL\MySQL Server 8.0\bin\mysql.exe" -u root -p
```

#### 3. ポート3306が使用可能か確認

   ```cmd
   # Windows
   netstat -an | findstr 3306
   
   # macOS/Linux
   netstat -an | grep 3306
   ```
   何も表示されない場合は、MySQLが起動していない可能性があります。

#### 4. ユーザー名とパスワードが正しいか確認

   ```cmd
   mysql -u root -p
   ```
   またはフルパスで：
   ```cmd
   "C:\Program Files\MySQL\MySQL Server 8.0\bin\mysql.exe" -u root -p
   ```

#### 5. MySQLのエラーログを確認

   ```cmd
   # エラーログの場所を確認（通常は以下のいずれか）
   # C:\ProgramData\MySQL\MySQL Server 8.0\Data\*.err
   # またはインストール時に指定したデータディレクトリ
   ```

#### 6. よくある問題と解決方法

- **サービスが見つからない**: MySQLがサービスとしてインストールされていない可能性があります。MySQLインストーラーで再インストールするか、手動でサービスを登録してください。
- **ポート3306が既に使用されている**: 他のアプリケーションがポート3306を使用している可能性があります。`netstat -ano | findstr 3306`でプロセスIDを確認し、必要に応じて終了してください。
- **パスワードを忘れた**: MySQLのパスワードリセット手順に従ってください。

### ポート8080が既に使用されている場合

`main.go`のポート番号を変更するか、使用中のプロセスを終了してください。

### DockerのMySQLサーバーに入る方法

Docker Composeを使用している場合、以下のコマンドでMySQLコンテナに接続できます：

#### 方法1: docker-compose execを使用（推奨）

```bash
# MySQLクライアントで直接接続（パスワード: password）
docker-compose exec mysql mysql -u root -ppassword qr_order_db

# または、パスワードを対話的に入力
docker-compose exec mysql mysql -u root -p qr_order_db
```

#### 方法2: コンテナのシェルに入ってからMySQLに接続

```bash
# コンテナのシェルに入る
docker-compose exec mysql bash

# コンテナ内でMySQLに接続
mysql -u root -ppassword qr_order_db
# または
mysql -u root -p qr_order_db
```

#### 方法3: dockerコマンドを直接使用

```bash
# コンテナ名を確認
docker ps

# コンテナ名を使用して接続（例：qr-order-app-server-mysql-1）
docker exec -it qr-order-app-server-mysql-1 mysql -u root -ppassword qr_order_db

# または、パスワードを対話的に入力
docker exec -it qr-order-app-server-mysql-1 mysql -u root -p qr_order_db
```

#### 接続情報

- **ユーザー名**: `root`
- **パスワード**: `password`（docker-compose.ymlで設定）
- **データベース名**: `qr_order_db`
- **ホスト**: `localhost`（ホストマシンから接続する場合、ポート3307を使用）
- **ポート**: `3307`（ホストマシンから接続する場合、docker-compose.ymlで`3307:3306`にマッピング）

#### ホストマシンから直接接続する場合

ホストマシンにMySQLクライアントがインストールされている場合：

```bash
mysql -h 127.0.0.1 -P 3307 -u root -ppassword qr_order_db
```

**注意**: 
- Docker Composeを使用している場合、コンテナ名は通常`<プロジェクト名>-mysql-1`の形式になります
- プロジェクト名は`docker-compose.yml`があるディレクトリ名から取得されます
- `docker ps`コマンドで実行中のコンテナ名を確認できます

## 注意事項

- デフォルトのパスワード（`password`）は本番環境では使用しないでください
- 環境変数を設定する場合は、セキュリティに注意してください
- マイグレーションは自動的に実行されますが、エラーが発生した場合はログを確認してください



