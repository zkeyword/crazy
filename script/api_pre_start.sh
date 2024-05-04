#!/bin/bash
ulimit -n 20000
nohup ./pre_api start -c pre_api.json &
