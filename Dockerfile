FROM golang:1.15-alpine

RUN mkdir /app
ADD zc-image-builder /app
ADD images /app/images
ADD metadata /app/metadata
RUN ["chmod", "+x", "/app/zc-image-builder"]
WORKDIR /app
CMD ["/app/zc-image-builder"]
