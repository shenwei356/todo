# TODO list

Just a toy for learning web development.

[Demo](http://app.shenwei.me:8080)

## Architecture

Frontend

- HTML5
- CSS3, [Bootstrap4](https://getbootstrap.com/docs/4.1/)
- Javascript, [JQuery3](https://api.jquery.com/)

Backend

- Router: [chi](https://github.com/go-chi/chi)
- DB: [bbolt](https://github.com/etcd-io/bbolt)([storm](https://github.com/asdine/storm))

## RESTful API

- `/items/`
    - `GET`
    - `PUT`
- `/items/{id}/`
    - `GET`
    - `PUT`
    - `DELETE`
- `/items/search`
    - `GET`

## Try

[download](https://github.com/shenwei356/todo/releases), decompress and run. For example:

    $ tar -zxvf todo_linux_amd64.tar.gz
    $ cd todo
    $ ./todo

Visit [http://127.0.0.1:8080](http://127.0.0.1:8080) .

