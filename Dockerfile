# The FROM instruction sets the Base Image for subsequent instructions.
# Using Nginx as Base Image
FROM golang:latest
MAINTAINER <284077318@qq.com> 
COPY . $GOPATH/src/github/jicg/myblog
# The RUN instruction will execute any commands
# Adding HelloWorld page into Nginx server
WORKDIR $GOPATH/src/github/jicg/myblog
RUN go get myblog & go install -a myblog
# The EXPOSE instruction informs Docker that the container listens on the specified network ports at runtime
EXPOSE 8080
# The CMD instruction provides default execution command for an container
# Start Nginx and keep it from running background
CMD myblog