#!/bin/bash

export EPHEMERAL_IP_ADDR=$(hostname -i)

exec kamailio \
    -E \
    -DD \
    --substdefs="/EPHEMERAL_IP_ADDR/$EPHEMERAL_IP_ADDR/" \
    -f /etc/kamailio/kamailio.cfg 