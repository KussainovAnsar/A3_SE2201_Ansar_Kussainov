## Running Tests:

Navigate to the directory containing the test files. Run the following command in your terminal:

go test

Integration tests:
TestCreateMovie: Creates a new movie record in the database.
TestCreateMovie_InvalidTitle: Attempts to create a movie record with an empty title field. Expects an error response.
TestCreateMovie_InvalidYear: Attempts to create a movie record with an invalid year value. Expects an error response.
TestDeleteMovie: Deletes a movie record from the database.

Unit Tests:
TestCreateUser: Creates a new user account.
TestCreateUser_InvalidEmail: Attempts to create a user account with an invalid email format. Expects an error response.
TestGetAuthenticationToken: Retrieves an authentication token for a user.
TestGetAuthenticationToken_InvalidBody: Attempts to retrieve an authentication token with an invalid request body. Expects an error response.
TestActivateAccount: Activates a user account using a valid activation token.
TestActivateAccount_InvalidToken: Attempts to activate a user account using an invalid activation token. Expects an error response.
