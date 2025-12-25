# TomcatKit

**[English](README.md)** | **[한국어](README_KR.md)** | **[日本語](README_JP.md)**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/playok/TomcatKit)](https://goreportcard.com/report/github.com/playok/TomcatKit)
[![Tomcat](https://img.shields.io/badge/Tomcat-9.0-F8DC75?style=flat&logo=apache-tomcat)](https://tomcat.apache.org/)
[![AI Generated](https://img.shields.io/badge/AI%20Generated-Claude%20Code-blueviolet?style=flat&logo=anthropic)](https://claude.ai/claude-code)

Apache Tomcat 9.0設定管理のためのCLIベースTUI（テキストユーザーインターフェース）ユーティリティです。

## デモ

![TomcatKit Demo](docs/assets/demo.gif)

## 機能

- **インタラクティブTUI**: [tview](https://github.com/rivo/tview)を使用したncursesスタイルのターミナルインターフェース
- **包括的な設定**: Tomcat 9.0のすべての主要な設定領域をカバー
- **自動検出**: 環境変数、一般的なパス、実行中のプロセスからTomcatインストールを自動検出
- **安全な編集**: 設定ファイル変更前に自動バックアップを作成
- **マルチインスタンス対応**: 最近使用したTomcatインスタンスを記憶
- **多言語対応**: 英語、韓国語、日本語（F2で切り替え）
- **カラーUI**: 意味に基づいた直感的なボタンカラー
  - 緑: 保存、追加、適用
  - 赤: 削除、除去、戻る
  - 黄: キャンセル
  - 青: ナビゲーション（コンテキスト、パラメータ）
- **コンテキストヘルプ**: 各設定フィールドのプロパティヘルプパネル
- **リアルタイムXMLプレビュー**: 設定変更のリアルタイムプレビュー

## 対応設定モジュール

| モジュール | 状態 | 説明 |
|------------|------|------|
| Server | 完了 | server.xmlコア設定（Server、Service、Engine、Host） |
| Connector | 完了 | HTTP、AJP、SSL/TLSコネクタとスレッドプール |
| Security/Realm | 完了 | 認証Realmとtomcat-users.xml管理 |
| JNDI Resources | 完了 | DataSource、Mail Session、Environment、Resource Links |
| Virtual Hosts | 完了 | Host、Context、Parameters、Session Manager設定 |
| Valves | 完了 | AccessLog、RemoteAddr、RemoteIp、ErrorReport、SSO、StuckThreadバルブ |
| Clustering | 完了 | セッションレプリケーション、メンバーシップ、インターセプター、ファームデプロイヤー |
| Logging | 完了 | JULI logging.properties、ファイルハンドラー、ロガー |
| Context | 完了 | context.xml設定、リソース、クッキー、セッションマネージャー |
| Web | 完了 | web.xmlサーブレット、フィルター、セッション、セキュリティ制約 |
| Quick Templates | 完了 | 仮想スレッド、HTTPS、コネクションプール、Gzip、セキュリティ |

## インストール

### ソースからビルド

```bash
# リポジトリをクローン
git clone https://github.com/playok/tomcatkit.git
cd tomcatkit

# ビルド
make build

# またはgoを直接使用
go build -o bin/tomcatkit ./cmd/tomcatkit
```

### 必要条件

- Go 1.21以降
- Apache Tomcat 9.0インストール

## 使用方法

### 基本的な使用方法

```bash
# 自動検出で実行
./bin/tomcatkit

# Tomcatパスを明示的に指定
./bin/tomcatkit -home /opt/tomcat -base /var/tomcat/instance1

# バージョン表示
./bin/tomcatkit -version

# ヘルプ表示
./bin/tomcatkit -help
```

### CLIオプション

| オプション | 説明 |
|------------|------|
| `-home` | CATALINA_HOMEパス（Tomcatインストールディレクトリ） |
| `-base` | CATALINA_BASEパス（インスタンスディレクトリ、デフォルト：CATALINA_HOME） |
| `-version` | バージョン情報を表示 |
| `-help` | ヘルプメッセージを表示 |

### ナビゲーション

| キー | アクション |
|------|------------|
| 矢印キー | メニューとリストをナビゲート |
| Enter | アイテム選択または確認 |
| Escape | 一つ前に戻る |
| Tab | フォームフィールド間を移動 |
| F2 | 言語切り替え（EN/KR/JP） |
| q | アプリケーション終了 |
| Ctrl+C | 強制終了 |

## プロジェクト構造

```
tomcatkit/
├── cmd/
│   └── tomcatkit/
│       └── main.go           # アプリケーションエントリーポイント
├── internal/
│   ├── config/
│   │   ├── tomcat.go         # Tomcatインスタンス設定
│   │   ├── settings.go       # アプリケーション設定の永続化
│   │   ├── server/           # server.xmlタイプと操作
│   │   ├── connector/        # コネクタプロトコルとデフォルト
│   │   ├── realm/            # Realmタイプとtomcat-users.xml
│   │   ├── jndi/             # JNDIリソースタイプとcontext.xml
│   │   ├── logging/          # ロギング設定
│   │   └── web/              # web.xmlタイプと操作
│   ├── detector/             # Tomcat自動検出
│   ├── i18n/                 # 国際化（EN/KR/JP）
│   ├── parser/               # XMLパースユーティリティ
│   └── tui/
│       ├── app.go            # メインTUIアプリケーション
│       └── views/            # 設定ビュー
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 設定ファイル

TomcatKitが管理するTomcat設定ファイル：

- `$CATALINA_BASE/conf/server.xml` - メインサーバー設定
- `$CATALINA_BASE/conf/tomcat-users.xml` - ユーザーとロール定義
- `$CATALINA_BASE/conf/context.xml` - デフォルトコンテキスト設定
- `$CATALINA_BASE/conf/web.xml` - デフォルトWebアプリケーション設定
- `$CATALINA_BASE/conf/logging.properties` - JULIロギング設定

## 設定保存場所

アプリケーション設定の保存場所：
- Linux/macOS: `~/.config/tomcatkit/settings.json`
- Windows: `%APPDATA%\tomcatkit\settings.json`

保存される設定：
- 最後に使用したTomcatインスタンス
- 最近のインスタンスパス
- 優先言語

## このプロジェクトについて

このプロジェクトは楽しみと学習を目的として作成された**趣味プロジェクト**です。AI支援開発を探求し、Tomcat管理者に役立つツールを提供するために作られました。

### AI生成

このプロジェクトは**[Claude Code](https://claude.ai/claude-code)**（AnthropicのClaude）を使用してAIによって完全に生成されました。

- **AIモデル**: Claude Opus 4.5 (`claude-opus-4-5-20251101`)
- **開発ツール**: Claude Code CLI
- **人間の役割**: プロジェクト方針、要件仕様、レビュー

このリポジトリのすべてのコード、ドキュメント、設定はAI支援開発を通じて生成されました。AIがアーキテクチャ設計、実装、デバッグ、ドキュメント作成を担当し、人間がガイダンスと検証を提供しました。

> **注意**: これは個人の趣味プロジェクトであり、Apache Software Foundationとは関係ありません。

## ライセンス

MITライセンス - 詳細は[LICENSE](LICENSE)ファイルを参照してください。

## コントリビューション

コントリビューションを歓迎します！お気軽にPull Requestを提出してください。

## 作者

[playok](https://github.com/playok)

---

<p align="center">
  <sub>Claude CodeのAI支援で構築</sub>
</p>
