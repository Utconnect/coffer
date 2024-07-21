#!/usr/bin/env bash

NEW_VERSION=$(npx standard-version --dry-run | grep 'tagging release' | tail -n1 | awk '{print $4}')
SONAR_CLOUD_FILE=.sonarcloud.properties

# Check if NEW_VERSION is set and update sonar-project.properties
if [[ -n "$NEW_VERSION" ]]; then
  echo -e "sonar.projectVersion=$NEW_VERSION" > $SONAR_CLOUD_FILE
else
  echo -e "sonar.projectVersion=latest" > $SONAR_CLOUD_FILE
fi

# Add sonar-project.properties to git
git add $SONAR_CLOUD_FILE