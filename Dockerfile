FROM alpine
ENV TZ=Asia/Shanghai

ENV API_HOST_LS :80
RUN apk add --no-cache tzdata && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY test-abci-linux-amd64 /test-abci
RUN chmod +x /test-abci
EXPOSE 80

CMD ["/test-abci"]
