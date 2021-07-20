# Go Template Online Parser
https://go-template-debugger.herokuapp.com/

# Description
GoTemplate online parser + helm parser online to make it easier to debug go-templates and helm templates

# Local Development

## UI Only
```bash
cd webui-vue
yarn serve
```

## Backend
```bash
./build.sh
docker run -it -e PORT=3000 -p 3000 docker.io/library/gotplonline
```
or 
```bash
./build.sh
go run main.go
```

# Deploy Heroku
```bash
heroku container:login
git add -A
git commit -m "" -a
git push heroku main
```


