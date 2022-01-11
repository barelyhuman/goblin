FROM golang:1.16
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go mod download github.com/smartystreets/goconvey
RUN go build -o /server cmd/goblin-api/main.go
CMD [ "/server" ]