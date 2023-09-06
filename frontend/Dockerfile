FROM oven/bun

WORKDIR /home/bun/app
RUN ln -sf /bin/bash /bin/sh
COPY . .
RUN bun i && bun --bun run build
CMD [ "bun","./build/index.js"]