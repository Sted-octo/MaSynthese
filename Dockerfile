FROM alpine as build
COPY . /
RUN chmod +x /out/octoptimist
RUN mkdir -m777 -p /out/static
RUN mkdir -m777 -p /out/private


ENV PORT="9090"
ENV CLIENT_ID=""
ENV CLIENT_SECRET=""
ENV REDIRECT_URL=""
ENV OCTOPOD_DOMAIN=""

VOLUME /out/static
VOLUME /out/private

EXPOSE 9090
WORKDIR /out
ENTRYPOINT ["/out/octoptimist"]
