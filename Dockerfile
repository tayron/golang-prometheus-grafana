FROM ubuntu
#FROM gcr.io/distroless/base-debian10
MAINTAINER Hospeda App <ti@hospeda.app>

WORKDIR /aplicacao
EXPOSE 80

ENTRYPOINT ["./web_application"]