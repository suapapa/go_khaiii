FROM ubuntu:20.04 AS khaiii-builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    wget \
    libboost-all-dev

WORKDIR /src
RUN wget https://github.com/kakao/khaiii/archive/refs/tags/v0.4.tar.gz && \
    tar -xzf v0.4.tar.gz && \
    rm v0.4.tar.gz

WORKDIR /src/khaiii-0.4
RUN mkdir build
WORKDIR /src/khaiii-0.4/build
RUN cmake -E env CXXFLAGS="-w" cmake ..
RUN mkdir /usr_local
RUN cmake -E env CXXFLAGS="-w" cmake -DCMAKE_INSTALL_PREFIX=/usr_local -DCMAKE_BUILD_TYPE=Release  ..
# CMD ["/bin/bash"]
RUN make -j$(nproc)
# RUN make large_resource
RUN make resource
RUN make install

FROM golang:1.24

COPY --from=khaiii-builder /usr_local/ /usr/local/
RUN ldconfig

COPY examples/analyze /app
WORKDIR /app
RUN go mod init example
RUN go get
RUN go build 

RUN apt update && apt install -y locales
RUN locale-gen en_US.UTF-8
RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8

CMD ["/bin/bash"]

