# Phase 1 – Air and Authentication Features

## Feature Overview

Feature name: Authentication Feature

Purpose: Authentication flow will give users the incentive of not losing their saved links.

## Problem Statement

*make believe mode: on*.

Ever since our user count reached 1 trillion, we have been facing difficulties maintaining storage space as the number of saved links has grown to approximately 6 x 10^128, significantly increasing our server costs. We need to implement a mechanism to detect and remove unused links while retaining our users. After holding hundreds of discussion groups last week, it was decided that we are in dire need of an authentication flow. This flow will help us determine which links are actively used and provide our real users with more incentive to continue using our product.

## Goals and Objectives

1. Implement the authentication flow.
2. Implement Air, a Golang module, to make our development process faster.
3. Implement a parallel worker to delete unused links.

## Feature Scope

1. Authentication scope:

    - Add sign-in flow.
    - Add sign-up flow.
    - Add sign-out flow.
    - Add token mechanisms for authorized users.

2. Implement Air scope:

    - Add basic configuration to run Air smoothly.

3. Parallel worker scope:

    - Add new service to automatically delete unused links a week after their creation date.

## High-Level Timeline

1. Authentication Flow: 2 MD.
2. Air Implementation: 0.5 MD.
3. Parallel Worker: 2 MD.
