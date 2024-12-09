Feature: Add new user
  As a user, I want to add a new user, so later can be used for any further analysis.

  Background:
    Given there is a clean "postgres" database

  Scenario: Add new user successfully
    When I request HTTP endpoint with method "POST" and URI "/v1/users"
    And I request HTTP endpoint with body
    """
    {
      "first_name": "Alice",
      "last_name": "Bob",
      "nickname": "AB123",
      "password_hash": "f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d",
      "email": "alice@bob.com",
      "country": "UK"
    }
    """

    Then I should have response with status "Created"
    And I should have response with header "Content-Type: application/json"
    And I should have response with body like
    """
    {
      // Ignoring id dynamic values.
      "id": "<ignore-diff>",
    }
    """
    And Then these rows are available in table "users" of database "postgres"
      | first_name | last_name | nickname | password_hash                                                    | email         | country |
      | Alice      | Bob       | AB123    | f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d | alice@bob.com | UK      |


  Scenario: Add new user failed, already exists
    Given these rows are stored in table "users" of database "postgres":
      | first_name | last_name | nickname | password_hash                                                    | email         | country |
      | Alice      | Bob       | AB123    | f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d | alice@bob.com | UK      |

    When I request HTTP endpoint with method "POST" and URI "/v1/users"
    And I request HTTP endpoint with body
    """
    {
      "first_name": "Alice",
      "last_name": "Bob",
      "nickname": "AB123",
      "password_hash": "f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d",
      "email": "alice@bob.com",
      "country": "UK"
    }
    """

    Then I should have response with status "Conflict"
    And I should have response with body like
    """
    {
      "code": 409,
      "message": "user already exists",
      "error": "<ignore-diff>"
    }
    """


  Scenario: Add new user failed, invalid argument
    When I request HTTP endpoint with method "POST" and URI "/v1/users"
    And I request HTTP endpoint with body
    """
    {
      "first_name": "",
      "last_name": "",
      "nickname": "",
      "password_hash": "f6b7e19e0d867de6c039187905",
      "email": "alice",
      "country": "U"
    }
    """

    Then I should have response with status "Bad Request"
    And I should have response with body like
    """
    {
      "code": 400,
      "message": "Bad Request",
      "error": "<ignore-diff>",
      "details": [
          {"field": "first_name", "description": "must not be empty"},
          {"field": "last_name", "description": "must not be empty"},
          {"field": "password_hash", "description": "invalid hash"},
          {"field": "email", "description": "must be a valid email"},
          {"field": "country", "description": "must have 2 characters"}
      ]
    }
    """