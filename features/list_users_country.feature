Feature: List users by country
  As a user, I want to know which users belongs to a country, so I can analyze the data.

  Background:
    Given there is a clean "postgres" database

  Scenario: List users successfully
    Given these rows are stored in table "users" of database "postgres":
      | id                                   | first_name | last_name | nickname | password_hash                                                    | email                       | country |
      | 26ef0140-c436-4838-a271-32652c72f6f2 | Alice      | Bob       |          | f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d | alice@bob.com               | UK      |
      | 207a6329-ad70-4294-bf27-5d37cf6fc8cf | Jan        | Watkins   | anim     | 8c1b486e26464ebecb042095cae3d251148f8006e35397409e3902ce78d82a13 | janwatkins@beadzza.com      | MR      |
      | 29d7fe1d-6d03-4c52-9880-d39788f9c227 | Lina       | Lowe      | magna    | 41eeaa061fa11f084957d4522cb4b408dbe4b16f446c513883d8c81e66da33f6 | linalowe@beadzza.com        | UK      |
      | 87c1eb37-aca4-4842-904b-f82c720f2f86 | Stuart     | Lancaster | laboris  | a080aaa8a868f6cf92593478bd9a6a8fb53b772a42ba163bb6d38765bde918bd | stuartlancaster@beadzza.com | ZW      |
      | 1f762b7e-680c-4e7c-b617-84a62d364444 | Kelli      | Herring   | elit     | 234ca9ff96989baf042f59e11ad53adce2488484aabbd0a890fde266c6d8ca5c | kelliherring@beadzza.com    | VA      |
      | 8276758c-0256-4978-9903-cd8924b77b97 | Amelia     | Clements  |          | e4288b26ddd516a83bfaee5f9ae8224010a327286eb5531d00859ea6ba00b5f6 | ameliaclements@beadzza.com  | PK      |
      | f1ec4c49-2166-45d2-988f-cb632bd380f9 | Roman      | Keith     | dolor    | 80e967e6c166120fc14badb021298fdb9ae5f20224d4c6c416d9898cfcc3b7e7 | romankeith@beadzza.com      | UK      |

    When I request HTTP endpoint with method "GET" and URI "/v1/users?country=UK"
    Then I should have response with status "OK"
    And I should have response with header "Content-Type: application/json"
    And I should have response with body
    """
    {
      "users":[
        {
          "id": "26ef0140-c436-4838-a271-32652c72f6f2",
          "first_name": "Alice",
          "last_name": "Bob",
          "nickname": "",
          "password_hash": "f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d",
          "email": "alice@bob.com",
          "country": "UK"
        },{
          "id": "29d7fe1d-6d03-4c52-9880-d39788f9c227",
          "first_name": "Lina",
          "last_name": "Lowe",
          "nickname": "magna",
          "password_hash": "41eeaa061fa11f084957d4522cb4b408dbe4b16f446c513883d8c81e66da33f6",
          "email": "linalowe@beadzza.com",
          "country": "UK"
        },{
          "id": "f1ec4c49-2166-45d2-988f-cb632bd380f9",
          "first_name": "Roman",
          "last_name": "Keith",
          "nickname": "dolor",
          "password_hash": "80e967e6c166120fc14badb021298fdb9ae5f20224d4c6c416d9898cfcc3b7e7",
          "email": "romankeith@beadzza.com",
          "country": "UK"
        }
      ],
      "nextPageToken":""
    }
    """