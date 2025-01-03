FROM node:22

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

EXPOSE 3000

CMD ["npm", "run", "start:dev"]
