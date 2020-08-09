#syntax=docker/dockerfile:experimental
FROM sqitch/sqitch

COPY . /sqitch
WORKDIR /sqitch

ENTRYPOINT [ "sqitch" , "deploy", "db:pg://root:root@house_finder_db:5432/house_finder"]
