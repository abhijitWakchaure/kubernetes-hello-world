FROM alpine
LABEL maintainer="Abhijit Wakchaure"
# RUN apt-get update && apt-get install ca-certificates -y
WORKDIR /app
ADD ./app-server ./app-server
ENTRYPOINT [ "./app-server" ]
EXPOSE 8080