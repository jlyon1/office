FROM node:8 as npmenv

ADD /office /office

WORKDIR /office

# Install npm dependencies and build
RUN npm install
RUN npm run generate

FROM golang:1.11 as go

WORKDIR /go/src/github.com/jlyon1/office/
COPY . /go/src/github.com/jlyon1/office

RUN go get github.com/go-chi/chi
RUN go get github.com/sahilm/fuzzy
RUN go build cmd/server/main.go
RUN ls

FROM golang:1.11
COPY --from=npmenv /office/dist/ office/dist
COPY --from=go /go/src/github.com/jlyon1/office/main main
COPY --from=go /go/src/github.com/jlyon1/office/script/ script/

RUN ls office/dist

CMD ["./main"]