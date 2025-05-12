# fib_Api
## 概要
今回使用した技術
- 言語 : Go
- 使用した主要ライブラリ:
  - gin-gonic/gin v1.10.0 - Webフレームワーク
  - stretchr/testify v1.10.0 - テストツールキット
- 使用したデプロイ環境 : [render.com](https://render.com/)
- サーバーエンドポイントURL : https://fib-api-yn1x.onrender.com ※最初のリクエストからサーバーが起動するまで1~2分掛かります。 2回目以降は時間はあまり掛かりません。
  - 正常系
    - [フィボナッチ数 1](https://fib-api-yn1x.onrender.com/fib/1)
    - [フィボナッチ数 50](https://fib-api-yn1x.onrender.com/fib/50)
    - [フィボナッチ数 100](https://fib-api-yn1x.onrender.com/fib/100)
    - [フィボナッチ数 150](https://fib-api-yn1x.onrender.com/fib/150)
  
  - 異常系
    - [フィボナッチ数 0](https://fib-api-yn1x.onrender.com/fib/0)
    - [フィボナッチ数 -1](https://fib-api-yn1x.onrender.com/fib/-1)
    - [不正な入力: 文字列](https://fib-api-yn1x.onrender.com/fib/helloworld)
    - [範囲外の値: 9999](https://fib-api-yn1x.onrender.com/fib/9999)
   
## フォルダ構成
```
Directory structure:
└── fib_api/
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── Makefile
    ├── lib/
    │   └── features/
    │       └── fibonacci_rest_api/
    │           ├── entities/
    │           │   └── fibonacci.go
    │           └── rest/
    │               └── fibonacci_rest_controller.go
    ├── test/
    │   └── features/
    │       └── fibocci_rest_api/
    │           ├── entities/
    │           │   └── fibonacci_test.go
    │           └── rest/
    │               └── fibonacci_rest_controller_test.go
    └── .github/
        ├── actions/
        │   └── go-ut/
        │       └── action.yml
        └── workflows/
            └── ci.yml
```
### フォルダの説明
- lib : ソースコード本体が格納されているフォルダ
- features : 機能別でコードを格納するフォルダ
- fibonacci_rest_api : 今回の課題であるフィボナッチ数APIに関連するコードを格納するフォルダ
- entities : Domain層におけるエンティティクラス
- rest : 今回はフィーチャードアーキテクチャを採用し、明確なレイヤードアーキテクチャではないので、restRepositoryImplのような実装本体をこちらに実装。ここではAPIの値を定義

- test : テスト本体を格納
※test以下についてはlibフォルダ内と同じ構成で作成

