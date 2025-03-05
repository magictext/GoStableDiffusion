FROM debian:bookworm-slim AS downloader

RUN apt-get update && \
    apt-get install -y curl

WORKDIR /models
RUN curl -L -O https://huggingface.co/CompVis/stable-diffusion-v-1-4-original/resolve/main/sd-v1-4.ckpt

FROM golang:1.24-bookworm AS builder

RUN apt-get update && \
    apt-get install -y cmake git

WORKDIR /src
COPY . .

RUN cd stable-diffusion.cpp && rm -rf build

RUN go generate
RUN cp stable-diffusion.cpp/build/bin/libstable-diffusion.so /lib/
RUN ldconfig
RUN go build -o main cmd/main/main.go

FROM debian:bookworm-slim

LABEL org.opencontainers.image.source = "https://github.com/Binozo/GoStableDiffusion"

RUN apt-get update && \
    apt-get install -y --no-install-recommends libgomp1

WORKDIR /app
COPY --from=builder /src/main .
COPY --from=builder /lib/libstable-diffusion.so /lib/libstable-diffusion.so
COPY --from=downloader /models/sd-v1-4.ckpt /app/models/sd-v1-4.ckpt

CMD ["/app/main"]