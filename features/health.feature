Feature: Service health
  As a user, I want to whether the service is up running and healthy.


  Scenario: Up and running
    When I check server is up and running
    Then It should be up and running

  Scenario: Healthy
    When I check server is healthy
    Then It should be healthy