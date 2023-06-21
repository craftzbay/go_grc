#!/usr/bin/env bash

go env -w GOPRIVATE=https://gitlab.com/yourname/gomath

git config \
  --global \
  url."https://<yourname>:<personal access token from gitlab>@gitlab.com".insteadOf \
  "https://gitlab.com"