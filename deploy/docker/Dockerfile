# Licensed Materials - Property of tenxcloud.com
# (C) Copyright 2018 Dreamxos. All Rights Reserved.
# 2018-10-08  @author gaozh

FROM alpine:3.7

RUN apk add --no-cache ca-certificates && mkdir -p /daas/
ADD controller /daas/
RUN chmod 755 /daas/controller
