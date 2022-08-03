FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY router ./router
RUN chmod +x ./router
CMD ["./router"]