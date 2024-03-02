Feature: Book management
    In order to use book API
    As a Librarian
    I need to be able to manage books

    Scenario: then user try to insert one book, one created book should be displayed by the system
        When I send "POST" request to "/v1/books" with payload:
        """
        {
            "title": "Dune",
            "author": "Frank Herbert"
        }   
        """
        Then the response code should be 201
        And the response payload should match json:
        """
        {
            "id": {id},
            "title": "Dune",
            "author": "Frank Herbert"
        }
        """

    Scenario: then user try to insert one book with title and author, the system should fail ```400``` response
        When I send "POST" request to "/v1/books" with payload:
        """
        {}   
        """
        Then the response code should be 400

    Scenario: then user try to insert one book with only title, the book should be created
        When I send "POST" request to "/v1/books" with payload:
        """
        {
            "title": "Dune"
        }   
        """
        Then the response code should be 201
        And the response payload should match json:
        """
        {
            "id": {id},
            "title": "Dune",
            "author": null
        }
        """

    Scenario: then user try to insert one book with only author, the book should be created
        When I send "POST" request to "/v1/books" with payload:
        """
        {
            "author": "Frank Herbert"
        }   
        """
        Then the response code should be 201
        And the response payload should match json:
        """
        {
            "id": {id},
            "title": null,
            "author": "Frank Herbert"
        }
        """
