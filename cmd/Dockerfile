FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY . /
CMD ["./dynapi"]
EXPOSE 443
VOLUME ["/certs"]


### Build ###
# CGO_ENABLED=0 GOOS=linux go build -a -o dynapi .
# docker build -t dynapi .

### Deploy ###
# docker tag dynapi codingconcepts/dynapi
# docker push codingconcepts/dynapi

### Run ###
# docker pull codingconcepts/dynapi
# docker run -d -p 443:443 -e HOST=api.robreid.io -e CERTS=/certs -e PORT=443 -v /certs:/certs codingconcepts/dynapi