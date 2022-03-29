SET MODULE_NAME="go-demo"

rm %MODULE_NAME%
rm %MODULE_NAME%.zip

SET GOOS=linux
SET GOARCH=amd64
go build ./main.go
SET GOOS=windows
SET GOARCH=amd64

zip %MODULE_NAME%.zip ./%MODULE_NAME% ./assets/