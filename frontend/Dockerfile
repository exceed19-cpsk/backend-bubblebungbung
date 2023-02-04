FROM node:18.13.0

ARG prot=3000
ENV PROT=${port}

COPY ./package*.json .

RUN npm install

COPY . .

EXPOSE ${port}

CMD ["npm", "start"]