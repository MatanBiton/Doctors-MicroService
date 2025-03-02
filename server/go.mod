module github.com/TekClinic/Doctors-MicroService/server

go 1.22.0

require (
	github.com/TekClinic/Doctors-MicroService/doctors_protobuf v0.100.0-integrated
	github.com/TekClinic/MicroService-Lib v0.1.3
	github.com/go-playground/validator/v10 v10.22.0
	github.com/mattn/go-sqlite3 v1.14.24
	github.com/stretchr/testify v1.10.0
	github.com/uptrace/bun v1.2.10
	github.com/uptrace/bun/dialect/pgdialect v1.2.10
	github.com/uptrace/bun/dialect/sqlitedialect v1.2.10
	github.com/uptrace/bun/driver/pgdriver v1.2.1
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.65.0
	k8s.io/apimachinery v0.31.0
)

require (
	github.com/alexlast/bunzap v0.1.0 // indirect
	github.com/coreos/go-oidc/v3 v3.11.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gabriel-vasile/mimetype v1.4.5 // indirect
	github.com/go-jose/go-jose/v4 v4.0.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/puzpuzpuz/xsync/v3 v3.5.1 // indirect
	github.com/sa-/slicefunk v0.1.4 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240812133136-8ffd90a71988 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	mellium.im/sasl v0.3.1 // indirect
)

replace github.com/TekClinic/Doctors-MicroService/doctors_protobuf => ./../doctors_protobuf
