FROM golang:1.17
WORKDIR /work
ADD . .
ENTRYPOINT ["go", "run", "github.com/cosmtrek/air", "-c", ".air.toml"]