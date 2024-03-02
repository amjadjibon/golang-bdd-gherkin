Feature: Read Books
    In order to use book API
    As a Librarian
    I need to be able to read books

    Scenario: get all books
        # make a book table in the database with title and author
        Given the following books exist:
            | title | author |
            | Book1 | Author1 |
            | Book2 | Author2 |
        When I send "GET" request to "/v1/books" with payload:
        """
        """
        Then the response code should be 200
        And the response should be a JSON array with the books
