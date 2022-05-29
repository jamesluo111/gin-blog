FROM scratch

WORKDIR $GOPATH/src/github.com/jamesluo111/gin-blog
COPY . $GOPATH/src/github.com/jamesluo111/gin-blog

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]