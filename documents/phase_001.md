# Phase 1 – Logging Feature and Clean Refactor

## Feature Overview

Feature name: File-based Logging.

Purpose: Logging is a necessity in a system, no matter how small.

## Problem Statement

Logging in the existing system is too unreliable as its not captured (only when you're running it from Terminal), which can only be viewed
in the development phase. We want to be able to check the logs later in the future.

## Goals and Objectives

1. Create a file-based logging.
2. Refactor to make a cleaner code base.

## Feature Scope

1. Logging Scope:

    - Log configuration should be centralized in one file.
    - Put all the logs inside a file. Don't forget to secure the file location.

2. Clean Refactor:

    - Separate server creation logic in main.go.
    - Separate common config creation logic i main.go.
    - Rename current usages of "main" and "short" on feature level.
    - Replace all lambda functions with named ones.

## High-Level Timeline

1. Logging: will take probably 2-3 hours top, testing included.
2. Clean Refactor: will take around a day.
