FROM ubuntu:20.04  
# RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY control ./control
COPY dispatcher ./dispatcher
COPY run.sh ./run.sh
RUN chmod +x run.sh
CMD ["./control"]