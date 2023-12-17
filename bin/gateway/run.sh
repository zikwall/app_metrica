#!/bin/bash

systemctl stop gateway.service
rm -rf /tmp/gateway.sock
systemctl start gateway.service
systemctl status gateway.service