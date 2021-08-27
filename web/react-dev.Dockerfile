FROM node:16.8.0-alpine
RUN mkdir /web
ADD . /web/
WORKDIR /web
# COPY package.json ./package.json
RUN yarn install
CMD ["yarn", "start"]