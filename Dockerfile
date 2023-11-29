FROM scratch

WORKDIR /app
COPY ./main ./

CMD ["./main"]
