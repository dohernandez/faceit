Feature: Add new user
  As a user, I want to add a new, so later can be used for any further analysis.

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
        "password_hash": "supersecurepassword",
        "email": "alice@bob.com",
        "country": "UK"
    }
    """

    Then I should have response with status "Created"
    And I should have response with header "Content-Type: application/json"
    And I should have response with body
    """
    {
      // Ignoring id dynamic values.
      "id": "<ignore-diff>",
    }
    """
    And Then these rows are available in table "users" of database "postgres"
      | first_name | last_name | nickname | password_hash       | email         | country |
      | Alice      | Bob       | AB123    | supersecurepassword | alice@bob.com | UK      |

  Scenario: Add new user failed, already exists
    Given these rows are stored in table "users" of database "postgres":
      | first_name | last_name | nickname | password_hash       | email         | country |
      | Alice      | Bob       | AB123    | supersecurepassword | alice@bob.com | UK      |

    When I request HTTP endpoint with method "POST" and URI "/v1/users"
    And I request HTTP endpoint with body
    """
    {
      "first_name": "Alice",
        "last_name": "Bob",
        "nickname": "AB123",
        "password_hash": "supersecurepassword",
        "email": "alice@bob.com",
        "country": "UK"
    }
    """

    Then I should have response with status "Conflict"
