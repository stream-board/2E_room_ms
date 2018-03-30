FROM        golang:1.9

   
# Setting up working directory
WORKDIR     /go/src/gin-container
ADD         . /go/src/gin-container

RUN     go get github.com/gin-gonic/gin
RUN     go get github.com/jinzhu/gorm
RUN     go get github.com/mattn/go-sqlite3

# RUN     go install github.com/gin-gonic/gin
# RUN     go install github.com/jinzhu/gorm
# RUN     go install github.com/mattn/go-sqlite3

# Restore godep dependencies
#RUN godep restore

EXPOSE 4004
ENTRYPOINT  ["/usr/local/go/bin/go"]
CMD     ["run", "src/main.go"]

#https://github.com/EarvinKayonga/gin-container/blob/master/Dockerfile