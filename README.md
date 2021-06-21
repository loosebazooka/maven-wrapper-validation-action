# Maven Wrapper Validation Action

A simple action to validate a `maven-wrapper.jar` binary checked into source control against maven central.
A simple search for `filename:maven-wrapper.jar` on GitHub reveals over half a million instances of this filename checked in. Let's try to make it less dangerous.

Inspired by https://github.com/gradle/wrapper-validation-action

## What this does
- Designed for https://github.com/takari/maven-wrapper
- Checks if the `maven-wrapper.jar` checked into the repo matches the file on [maven.org](https://search.maven.org/artifact/io.takari/maven-wrapper) by comparing `sha256` hashes
  - NOTE: The action downloads the maven-wrapper artifact to verify the hash since maven.org currently only store cryptographically insecure md5 and sha1 hashes.

## What this does NOT do
- Does NOT verify pgp signatures or signatures of any kind
- Does NOT verify binaries from other sources
- Does NOT ensure that your `mvnw` script is actually using this `maven-wrapper.jar`
- Does NOT prevent against attacks on maven.org (this action assumes it is safe)

## How to use
- Create a new action with the following configuration
```
name: Validate maven wrapper 

on: [push, pull_request]

jobs:
  validation:
    name: "Validation"
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: loosebazooka/maven-wrapper-validation-action@<tbd>
```
