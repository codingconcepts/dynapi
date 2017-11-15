FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY . /
CMD ["./dynoapi"]
EXPOSE 443
VOLUME ["/certs"]


### Build ###
# CGO_ENABLED=0 GOOS=linux go build -a -o dynoapi .
# docker build -t dynoapi .
# rm dynoapi

### Deploy ###
# docker tag dynoapi codingconcepts/dynoapi
# docker push codingconcepts/dynoapi

### Run ###
# docker pull codingconcepts/dynoapi
# docker run -d -p 443:443 -e HOST=api.robreid.io -e CERTS=/certs -e PORT=443 -v /certs:/certs codingconcepts/dynoapi
# docker run -d -p 1234:1234 -e HOST=api.robreid.io -e SSL=false -e PORT=1234 dynoapi