env GOOS=linux GOARCH=amd64 go build -v github.com/vorpalhex/discordia
mkdir discordia_linux
mv discordia discordia_linux/
cp ../config.sample.json discordia_linux/
cp ../readme.md discordia_linux/
