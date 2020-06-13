FROM golang:latest

RUN git clone https://github.com/mattrout92/games.git
RUN cd games/api && go install

CMD /go/bin/api
EXPOSE 8080
