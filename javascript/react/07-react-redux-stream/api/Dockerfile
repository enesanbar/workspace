FROM node:alpine
WORKDIR /usr/src/app
COPY package*.json ./
COPY . ./
RUN npm install

EXPOSE 3001
CMD ["npm", "start"]
