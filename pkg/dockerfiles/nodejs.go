package dockerfiles

const NodeJS = `FROM node:latest

RUN mkdir /app && chown appuser:appuser /app
WORKDIR /app

COPY package*.json

RUN npm install

COPY . .

EXPOSE 3000

USER appuser

CMD ["npm", "start"]
`
