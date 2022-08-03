FROM ubuntu:20.04 
# RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY daemon ./daemon
COPY scion ./scion
COPY dispatcher ./dispatcher
COPY run.sh ./run.sh
RUN chmod +x run.sh
CMD ["./daemon"]