# go-deamon

#Local Configuration

Install Glide (Package Manager)
``` shell
curl -L https://glide.sh/get -o $GOPATH/bin/glide | sh
chmod u+x $GOPATH/bin/glide
```


Intall Wercker CI
``` shell
curl -L https://s3.amazonaws.com/downloads.wercker.com/cli/stable/linux_amd64/wercker -o $GOPATH/bin/wercker
chmod u+x $GOPATH/bin/wercker
```

Install Postgresql
``` shell
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib
```
or read https://www.digitalocean.com/community/tutorials/postgresql-ubuntu-16-04-ru

Install Redis
``` shell
sudo apt-get install build-essential tcl
curl -O http://download.redis.io/redis-stable.tar.gz
tar xzvf redis-stable.tar.gz
make
sudo make install
```
or read https://www.digitalocean.com/community/tutorials/how-to-install-and-configure-redis-on-ubuntu-16-04

Get project
``` shell
go get github.com/anabiozz/go-daemon
```

#For Local Run (without Docker)
in command line:
``` shell
glide install 
glide up
redis-server
go run main.go
```

#Run in Docker container
in command line:
``` shell
runlocal
```

