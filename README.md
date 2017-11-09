## snakehack-go

A basic snakehack starter kit written in go.


### How to start

1) [Fork this repo](https://github.com/stair-ch/snakehack-go/fork).

2) Clone repo to your development environment:
```
git clone git@github.com:USERNAME/snakehack-go.git $GOPATH/github.com/USERNAME/snakehack-go
cd $GOPATH/github.com/USERNAME/snakehack-go
```

3) Compile the snakehack-go server.
```
go build
```
This will create a `snakehack-go` executable.

4) Modify your configuration file

5) Run the server.
```
./snakehack-go
```

6) Test the client in your browser: [http://127.0.0.1:4242](http://127.0.0.1:4242)

### Heroku
[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

#### Heroku-cli commands
create new heroku app, `--region eu` is important for fast enough response times
```
heroku create --region eu
```
push to heroku
```
git push heroku master
```
delete heroku git
```
git remote rm heroku
```

### [govendor](https://github.com/user/repo/blob/branch/other_file.md) 
govendor is used for dependency management

1) install or update govendor
```
go get -u github.com/kardianos/govendor
```

2) add to project
```
govendor init
```
This creates a `vendor/` directory and a `vendor.json` file in that directory.

3) adding dependencies
```
govendor fetch <package>
```

4) list status
```
govendor list
```
