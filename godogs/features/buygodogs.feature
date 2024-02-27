Feature: buy some godogs
    In order to have some fun
    As a godog fan
    I need to be able to buy some godogs

    Scenario: Filling the cart
        Given there are 5 godogs
        When I eat 5
        Then there should be none remaining
        Then Then I buy 4
        Then there should be 4 remaining
