FROM golang:1.20

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify \
  && go install github.com/pressly/goose/v3/cmd/goose@latest \
  && go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

COPY . .

# RUN go build -o bin

EXPOSE 8000

# CMD ["/app/bin"]

# run container with CMD ["/app/bin"]
# anonymous volume "-v /app/bin" prevents bin file from being
# overritten by bind mount "-v D:/code/rss:/app" 
# docker run --name rss_server --rm -p 8000:8000 -v D:/code/rss:/app -v /app/bin rss

# run container with terminal open to user
# docker run --name rss_server --rm --network rss_net -it -p 8000:8000 -v D:/code/rss:/app rss

# run shell inside container
# docker exec -it rss_server sh

# run postgres container
# docker run --name rss_db --rm -d --network rss_net -e POSTGRES_PASSWORD=postgres postgres

# copy file from container
# docker cp CONTANER_NAME:app/file.txt .