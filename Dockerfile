FROM mhart/alpine-iojs:latest

MAINTAINER Christoph Witzko <docker@christophwitzko.com>

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package.json /usr/src/app/
RUN npm install --production
COPY . /usr/src/app

EXPOSE 5000

CMD ["npm", "start"]
