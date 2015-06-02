FROM debian:jessie

RUN apt-get update && \
    apt-get install -y mongodb

RUN mkdir -p /data/db /etc/ssl/certs

COPY oauth /usr/bin/oauth
RUN chmod +x /usr/bin/oauth
COPY run.sh /usr/bin/oauth.sh
RUN chmod +x /usr/bin/oauth.sh

ENV HOST ":8080"
ENV SEC_KEY ""
ENV REDIRECT_URL ""
ENV GOOGLE_CLIENT_ID ""
ENV GOOGLE_SECRET ""
ENV GOOGLE_REDIRECT_URL ""

CMD ["oauth.sh"]