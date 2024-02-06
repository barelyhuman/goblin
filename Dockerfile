FROM node:lts as node-builder
WORKDIR /app 
COPY . . 
RUN npm i; npm run build 

FROM golang:1.19
WORKDIR /app
COPY go.* ./
RUN go mod download

ENV ORIGIN_URL="http://goblin.run"

COPY . ./

COPY --from=node-builder /app/www/assets /app/www/assets

RUN go build -o ./goblin-api ./cmd/goblin-api 

EXPOSE 3000

CMD [ "./goblin-api" ]