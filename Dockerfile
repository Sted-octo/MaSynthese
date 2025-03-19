FROM alpine as build
COPY . /
RUN chmod +x /out/octoptimist
RUN mkdir -m777 -p /out/static
RUN mkdir -m777 -p /out/static/css
RUN mkdir -m777 -p /out/private

COPY /static/. /out/static/
COPY /static/css/. /out/static/css
COPY /private/. /out/private/

ENV CLIENT_ID=""
ENV CLIENT_SECRET=""
ENV REDIRECT_URL=""
ENV OCTOPOD_DOMAIN=""

EXPOSE 9090
WORKDIR /out
ENTRYPOINT ["/out/octoptimist"]
