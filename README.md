# todolist-go-api
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kzuabe/todolist-go-api)
![GitHub](https://img.shields.io/github/license/kzuabe/todolist-go-api)
[![CI](https://github.com/kzuabe/todolist-go-api/actions/workflows/ci.yml/badge.svg)](https://github.com/kzuabe/todolist-go-api/actions/workflows/ci.yml)

Todoリストアプリ API Go実装

主な利用OSS
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [firebase/firebase-admin-go: Firebase Admin Go SDK](https://github.com/firebase/firebase-admin-go)


## 使い方


```bash
$ go mod download

$ export API_ENV=<develop|production> DSN="<user>:<password>:3306)/<dbname>" GOOGLE_APPLICATION_CREDENTIALS="<Path to service-account-file.json>"

$ make run
```

## ディレクトリ構成

```
.
├── Makefile
├── app # アプリケーションの実装
│   ├── config
│   ├── controller
│   ├── model
│   ├── repository
│   ├── router
│   └── usecase
├── cmd
│   └── todolist
│       └── main.go
├── docs
├── go.mod
├── go.sum
└── test # テストで利用するモック・データ
    ├── mocks
    └── testdata
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
