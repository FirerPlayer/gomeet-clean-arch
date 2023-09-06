FROM oven/bun

WORKDIR /home/bun/app
RUN ln -sf /bin/bash /bin/sh
COPY . .
RUN bun i
RUN bun --bun run build
# RUN cd ./build

CMD [ "bun", "./build/index.js"]