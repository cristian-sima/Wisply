var badUser = {
  name: "Hacker",
  password: "HackerPassword"
},
    user = {
      name: "Jameson Henry",
      email: "henry@oxford.ac.uk",
      password: "my-strong-password"
    };

function expectErrors(text) {
  expect(text).toContain("Sorry... Your request was not successful");
}

describe('Register', function() {

  it('goes to the register page from menu', function(done) {
      browser
      .url("/")
      .windowHandleMaximize()
      .element("#navbar-main")
      .click("a=Register")
      .pause(1000)
      .getTitle(function(err, title) {
          expect(err).toBe(undefined);
          expect(title).toBe('Create a new account');
      })
      .call(done);
    });
});
