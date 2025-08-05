ARG NODE_VERSION=20.19.0

FROM node:${NODE_VERSION}-alpine

ENV NODE_ENV=development

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev"]