FROM golang:1.20

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

 #Estamos mantendo o processo sendo executado.
CMD ["tail", "-f", "/dev/null"]
