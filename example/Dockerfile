FROM golang:alpine AS base

FROM base AS development

WORKDIR /sqlboiler-erg
COPY go.mod go.sum ./
RUN go mod download

WORKDIR /sqlboiler-erg/example

# RUN go install github.com/wearepointers/sqlboiler-erg@latest
RUN go install github.com/cespare/reflex@latest
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
