FROM scratch

USER 1000

WORKDIR /app
COPY ./main ./

CMD ["./main"]
