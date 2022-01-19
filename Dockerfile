FROM node:16-alpine3.15 as website-builder
WORKDIR /www
COPY ./www .
RUN yarn 
RUN yarn build 




FROM golang:1.16
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN mkdir -p ./static
COPY --from="website-builder" /www/build ./static
RUN go build -o /server cmd/goblin-api/main.go
CMD [ "/server" ]