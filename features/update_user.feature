Feature: Update user
  As a user, I want to update a user, so I can fix/modify some info.

  Background:
    Given there is a clean "postgres" database

  Scenario: Update user successfully
    Given these rows are stored in table "users" of database "postgres":
      | id                                   | first_name | last_name | nickname | password_hash                                                    | email         | country |
      | 26ef0140-c436-4838-a271-32652c72f6f2 | Alice      | Bob       |          | f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d | alice@bob.com | UK      |

    When I request HTTP endpoint with method "PATCH" and URI "/v1/users/26ef0140-c436-4838-a271-32652c72f6f2"
    And I request HTTP endpoint with body
    """
    {
      "nickname": "AB123",
      "country": "DE"
    }
    """

    Then I should have response with status "No Content"
    And I should have response with header "Content-Type: application/json"
    And Then these rows are available in table "users" of database "postgres"
      | first_name | last_name | nickname | password_hash                                                    | email         | country |
      | Alice      | Bob       | AB123    | f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d | alice@bob.com | DE      |
