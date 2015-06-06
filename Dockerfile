FROM google/golang
MAINTAINER 3GGS "3ggs.me"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app/team-3GGS/mate-go
ADD . /gopath/app/

RUN git clone https://github.com/Go-Docker-Hackathon/team-3GGS.git
RUN cd ./team-3GGS/mate-go/
ENTRYPOINT ["build.sh"]

EXPOSE 8083
CMD ["/gopath/app/team-3GGS/mate-go/bin/main"]
