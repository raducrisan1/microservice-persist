FROM alpine:3.8
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY ./microservice-persist /app/
WORKDIR /app
EXPOSE 3050
CMD ["./microservice-persist"]