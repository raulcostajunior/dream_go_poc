FROM golang:buster

WORKDIR /root/go_web_app

VOLUME /root/go_web_app/src

CMD ["/bin/bash"]