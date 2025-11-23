FROM node:25-alpine

ENV LANG C.UTF-8
ENV HOST 0.0.0.0


WORKDIR /app

COPY ./package*.json ./

RUN npm install

EXPOSE 3010

CMD ["npm", "run", "dev"]
