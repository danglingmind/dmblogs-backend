FROM golang:1.15.6-alpine

ADD app /home
        
WORKDIR /home

CMD ["./app"]