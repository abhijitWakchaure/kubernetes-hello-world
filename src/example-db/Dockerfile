FROM alpine
LABEL maintainer="abhijitwakchaure.2014@gmail.com"
# RUN apt-get update && apt-get install ca-certificates -y
WORKDIR /app
ADD ./db-server ./db-server
ENTRYPOINT [ "./db-server" ]
EXPOSE 8081