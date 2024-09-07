FROM node:latest

ENV APP_ROOT=/home/src

RUN mkdir -p $APP_ROOT
WORKDIR $APP_ROOT

COPY ./src/ $APP_ROOT/

CMD ["npm", "start"]

