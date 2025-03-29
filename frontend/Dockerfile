FROM node:21-alpine

# project dir
WORKDIR /app

COPY package*.json ./

# install dependencies
RUN npm install

# copy all other folders and files
COPY . .

EXPOSE 3000

CMD ["npm", "start"]