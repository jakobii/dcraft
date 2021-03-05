#!/bin/bash

export client_id=""
export state=""
export scope=""
export response_type="token"

curl -X POST "https://discord.com/api/oauth2/authorize?response_type=${response_type}&client_id=${client_id}&state=${state}&scope=${scope}"
