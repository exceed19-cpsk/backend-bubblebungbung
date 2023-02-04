FROM node:18.13.0

ARG port=8080
ENV PORT=${port}

COPY ./package*.json .

RUN npm install

COPY . .

EXPOSE ${port}

CMD ["npm", "start"]