FROM fuzzers/go-fuzz:1.2.0

#Update go

RUN rm -rf /usr/local/go
RUN curl -OL https://golang.org/dl/go1.20.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin


#Install go-fuzz-build 

RUN go install github.com/dvyukov/go-fuzz/go-fuzz-build@latest

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

WORKDIR /tmp/app/fuzz
RUN go-fuzz-build -libfuzzer -o fuzz.a
RUN clang -fsanitize=fuzzer fuzz.a -o fuzz.libfuzzer

ENTRYPOINT []
CMD ["/tmp/app/fuzz/fuzz.libfuzzer"]
