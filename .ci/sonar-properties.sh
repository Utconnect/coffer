#!/usr/bin/env bash

NEW_VERSION=$(npx standard-version --dry-run | grep 'tagging release' | tail -n1 | awk '{print $4}')
SONAR_CLOUD_FILE=.sonarcloud.properties

cp .ci/$SONAR_CLOUD_FILE.tmp $SONAR_CLOUD_FILE

if [[ -n "$NEW_VERSION" ]]; then
  echo -e "sonar.projectVersion=$NEW_VERSION" >> $SONAR_CLOUD_FILE
else
  echo -e "sonar.projectVersion=latest" >> $SONAR_CLOUD_FILE
fi

git add $SONAR_CLOUD_FILE