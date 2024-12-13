Feature: Delete user
  As a user, I want to delete a user, so I can remove the user info.

  Background:
    Given there is a clean "postgres" database

  Scenario: Delete user successfully
    Given these rows are stored in table "users" of database "postgres":
      | id                                   | first_name | last_name | nickname | password_hash                                                    | email         | country |
      | 26ef0140-c436-4838-a271-32652c72f6f2 | Alice      | Bob       |          | f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d | alice@bob.com | UK      |

    When I request HTTP endpoint with method "DELETE" and URI "/v1/users/26ef0140-c436-4838-a271-32652c72f6f2"

    Then I should have response with status "No Content"
    And I should have response with header "Content-Type: application/json"
    And no rows in table "users" of database "postgres"


  Scenario: Delete user failed, invalid argument
    When I request HTTP endpoint with method "DELETE" and URI "/v1/users/26ef0140-c436-4838-a271"

    Then I should have response with status "Bad Request"
    And I should have response with body like
    """
    {
      "code": 400,
      "message": "validation error",
      "error": "<ignore-diff>",
      "details": [
          {"field": "id", "description": "value must be a valid UUID"}
      ]
    }
    """
