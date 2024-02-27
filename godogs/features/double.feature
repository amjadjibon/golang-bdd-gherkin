Feature:
    As a gopher
    I was wandering around the internet
    I met a wizard
    He told me if I have even number of godogs
    Then He can double them

    Scenario: Doubled the number of godogs
        Given there are 8 godogs
        When I ask the wizard to double them
        Then I should have 16 godogs

    Scenario: Not doubled the number of godogs
        Given there are 7 godogs
        When I ask the wizard to double them
        Then I should have 7 godogs
