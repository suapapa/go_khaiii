# khaiii: go package (cgo wrapper) of khaiii, Kakao Hangul Analyzer III

## Requirement: build & install kahiii

Download(clone) [kahiii](https://github.com/kakao/khaiii) source, build and install;

    $ mkdir build && cd $_
    $ cmake -E env CXXFLAGS="-w" cmake .. # I need turn off -Wall in Ubuntu 20.04 (gcc 9.3.0)
    $ make
    $ sudo make install
    $ sudo ld-config

Install locations (default prefix):

* header: `/usr/local/include/khaiii/khaiii_api.h`
* library: `/usr/local/lib/libkhaiii.so*`
* resources: `/usr/local/share/khaiii/`

## Example

check out `_example/main.go`

## reference

* [21세기 세종계획 말뭉치 구축지침](https://ithub.korean.go.kr/user/total/referenceView.do?boardSeq=5&articleSeq=103&boardGb=T&isInsUpd=&boardType=CORPUS)
