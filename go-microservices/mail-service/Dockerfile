FROM golang:1.20-alpine AS Builder

WORKDIR /app
COPY . .

RUN go build -o ./dist . 

FROM alpine:latest AS Runner

WORKDIR /app

COPY --from=Builder /app/dist ./
COPY --from=Builder /app/smtp/emails ./emails
COPY --from=Builder /app/smtp/templates ./templates

CMD [ "./dist" ]
