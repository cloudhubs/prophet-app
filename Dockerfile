# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o goapp
# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
EXPOSE 8082
ENTRYPOINT ./goapp

#FROM alpine
#ADD main main
#RUN pwd
#RUN ls -l
#EXPOSE 8082
#ENTRYPOINT ["./main"]