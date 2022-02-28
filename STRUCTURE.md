# `demoapp` Project Structure

This document outlines the Go project structure for each demo app implementation. All web frameworks covered in this repository will follow the same project structure with only some minor framework-specific modifications.

**Status:** *Work in progress*

## Go Project Layout

    - cmd
        - api
            main.go
        - website
            main.go
    - internal
        - app
            - entities
        - pkg
            - http
                - entities
            - storage
                - entities
        - website.go
        - api.go
