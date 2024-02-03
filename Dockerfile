FROM golang:1.19
WORKDIR /app

COPY go.* ./
RUN go mod download

ENV GOBLIN_ORIGIN_URL=""
ENV ORIGIN_URL=""

COPY . ./
RUN cd www \
    ;make installLinux \
    ;make build \
    ;cd .. \
    ;rm -rf ./static \
    ;ln -sf ./www/dist ./static

RUN go build -o ./goblin-api ./cmd/goblin-api 

EXPOSE 3000

CMD [ "./goblin-api" ]