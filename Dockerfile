FROM debian:jessie

RUN apt-get update && \
    apt-get install -y mongodb

RUN mkdir -p /data/db /etc/ssl/certs /app

COPY run.sh /app/oauth.sh
RUN chmod +x /app/oauth.sh
COPY components /app/components
COPY oauth /app/oauth
RUN chmod +x /app/oauth


ENV HOST ":8080"
ENV SEC_KEY ""
ENV REDIRECT_URL ""
ENV GOOGLE_CLIENT_ID ""
ENV GOOGLE_SECRET ""
ENV GOOGLE_REDIRECT_URL ""

WORKDIR /app/
CMD ["/app/oauth.sh"]