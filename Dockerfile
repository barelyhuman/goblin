FROM node:20-alpine AS site_builder
WORKDIR /app
COPY ./www . 

RUN npm i \
    ; npm run build \
    ; rm -rf ./static \ 
    ; ln -sf ./www/_site ./static 


FROM golang:1.19
WORKDIR /app

COPY go.* ./
RUN go mod download

ENV GOBLIN_ORIGIN_URL="http://goblin.run"
ENV ORIGIN_URL="http://goblin.run"

COPY . ./
RUN rm -rf ./static ./www
COPY --from=site_builder /app/_site ./static

RUN go build -o ./goblin-api ./cmd/goblin-api 

EXPOSE 3000

CMD [ "./goblin-api" ]