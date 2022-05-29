go build -v -ldflags="-X 'covid-19-api/build.Version=$(./version.sh)' -X 'covid-19-api/build.User=$(id -u -n)' -X 'covid-19-api/build.Time=$(date)'" -o apiserver ./cmd/server
