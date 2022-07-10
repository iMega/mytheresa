ARG GO_IMG

FROM $GO_IMG as builder
ENV GOFLAGS=-mod=mod
ARG CWD
WORKDIR $CWD
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-s -w" -o $CWD/app .


FROM scratch
ARG CWD
COPY --from=builder $CWD/app .
CMD ["/app"]
