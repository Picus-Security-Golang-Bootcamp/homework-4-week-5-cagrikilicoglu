# Assignment | Week 5

> Bookstore simulating app

This is a bookstore application providing a rest API.

The app contains a database that has two tables, one for top-selling books of all time and one for the authors.
The database are created by the app. The book and author data is read from csv file of which the user can specify the path.

## Endpoints and Requests

#### Home screen

    `GET /`

#### Get all the books currently in the database.

    `GET /books/`

#### Get all the books including those deleted before.

    `GET /books/all`

#### Get only the books that are in stock.

    `GET /books/stock`

#### Get books under a certain price of your preferance.

     `GET /price/{priceunder}`

        Example Request: (get the books under price 32)

        `GET /price/32`

#### Get a book by its ID.

     `GET /books?id={id}`

        Example Request: (get the book with id 3)

        `GET /books?id=3`

#### Get a book by its ISBN.

     `GET /books?isbn={isbn}`

        Example Request: (get the book with isbn 9781128355898)

        `GET /books?isbn=9781128355898`

#### Get books by its name. (elastic search)

     `GET /books?name={name}`

        Example Request: (get the books with the name containing "the")

        `GET /books?name=the`

#### Delete a book from database. (soft-delete)

     `DELETE /books/delete?id={id}`

        Example Request: (delete the book with the id 4)

        `DELETE /books/delete?id=4`

### Order books from the database by their ID and of preferred quantity.

    `PATCH /books/order?id={id}&quantity={quantity}`

        Example Request: (order the book with the id 5 of quantity 2)

        `PATCH /books/order?id=5&quantity=2`

#### Add a new book to the database. (create the book on the database)

    `POST /books/add`

        Example Request Body:

        {"ID":"11","name":"Utopia","pageNumber":182,"stockNumber":20,"stockId":"11SF","price":14.7,"isbn":"9781128355898","authorID":"909","Author":{"ID":"909","name":"Thomas Moore"}}

#### Get all the authors in the database, with the books of the authors.

    `GET /authors/`

#### Get all the authors in the database, without the books of the authors.

    `GET /authors/*`

#### Get an author with his/her ID.

    `GET /authors?id={id}`

        Example Request: (get the author with id 101)

        `GET /authors?id=101`

#### Get authors by their name. (elastic search)

     `GET /authors?name={name}`

        Example Request: (get the authors with the name containing "j.")

        `GET /authors?name=j.`

#### Get the books of authors by their name. (elastic search)

     `GET /authors/books?name={name}`

        Example Request: (get the books of author with the name containing "antoine")

        `GET /author/books?name=antoine`

## Links

- Project repository: https://github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-cagrikilicoglu

## License

The project has no license.
