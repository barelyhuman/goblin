FROM node:16-alpine3.15 as website-builder
WORKDIR /www
COPY ./www .
RUN npm i -g pnpm
RUN pnpm install
RUN pnpm build

FROM golang:1.17
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN mkdir -p ./static
COPY --from="website-builder" /www/build ./static
RUN go build -o /server cmd/goblin-api/main.go

ENV PORT=3000
EXPOSE $PORT

CMD [ "/server" ]