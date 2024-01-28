Golang practice creating CRUD operations following tutorial on youtube

- [go crud](https://www.youtube.com/watch?v=lf_kiH_NPvM)
- [go auth](https://www.youtube.com/watch?v=ma7rUS_vW9M)

## Starting

- run `make dev`

## Feature

- CRUD
- Auth

## Endpoint list

- signup
  - type: POST
  - path: `/signup`
  - body JSON: email, password
- login
  - type: POST
  - path: `/login`
  - body JSON: email, password
- get post
  - type: GET
  - path: `/posts/:id`
- get all posts
  - type: GET
  - path: `/posts`
- create post
  - type: POST
  - path: `/posts`
  - body JSON: title, body
- edit post
  - type: PUT
  - path: `/posts/:id`
  - body JSON: title, body
- create post
  - type: POST
  - path: `/posts/:id`
  - body JSON: title, body
