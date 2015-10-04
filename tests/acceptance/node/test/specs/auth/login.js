
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

describe('Login', function() {

  it('goes to the login page from menu', function(done) {
      browser
          .url("/")
          .windowHandleMaximize()
          .element("#navbar-main")
          .click("a=Login")
          .pause(500)
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe('Login to Wisply');
          })
          .call(done);
  });

  it('submits empty form and receives errors', function(done) {
      browser
          .submitForm("#login-form")
          .getText(".alert-warning").then(expectErrors)
          .call(done);
  });

  it('submits form with just email completed', function(done) {
      browser
          .url("/auth/login")
          .setValue('#login-email', "admin@admin.com")
          .pause(1000)
          .submitForm("#login-form")
          .getText(".alert-warning").then(expectErrors)
          .call(done);
  });


  describe('Rejecting wrong data formats', function() {
      it('rejects wrong details', function(done) {
          browser
              .url("/auth/login")
              .setValue('#login-email', "wrong@wrong.com")
              .setValue('#login-password', "wrong")
              .pause(1000)
              .submitForm("#login-form")
              .getText(".alert-warning").then(expectErrors)
              .call(done);
      });
      it('rejects SQL injection', function(done) {
          browser
              .url("/auth/login")
              .setValue('#login-email', "' OR 1=1 -- a")
              .setValue('#login-password', "' OR 1=1 -- a")
              .pause(1000)
              .submitForm("#login-form")
              .getText(".alert-warning").then(expectErrors)
              .call(done);
      });
      it('rejects XSS', function(done) {
          browser
              .url("/auth/login")
              .setValue('#login-email', "<marquee>")
              .setValue('#login-password', "<marquee>")
              .pause(1000)
              .submitForm("#login-form")
              .getText(".alert-warning").then(expectErrors)
              .call(done);
      });

    describe('Email', function() {
      it('rejects an email with a wrong format', function(done) {
          browser
              .url("/auth/login")
              .setValue('#login-email', "wrongwrong.com")
              .submitForm("#login-form")
              .getText(".alert-warning").then(function(text){
                  expect(text).toContain("It must be a valid EMAIL type");
              })
              .call(done);
      });
      it('rejects an email too short', function(done) {
          browser
              .url("/auth/login")
              .setValue('#login-email', "a@a")
              .submitForm("#login-form")
              .getText(".alert-warning").then(expectErrors)
              .call(done);
      });
      it('rejects an email too big', function(done) {
          browser
              .url("/auth/login")
              .setValue('#login-email', "aaa@asaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasasa.com")
              .submitForm("#login-form")
              .getText(".alert-warning").then(function(text){
                  expect(text).toContain("should be between 3 and 60 characters");
              })
              .call(done);
      });
    });
      describe('Password', function() {
        it('rejects too big', function(done) {
            browser
                .url("/auth/login")
                .setValue('#login-email', "good@email.com")
                .setValue('#login-password', "aaaasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasasa.com")
                .submitForm("#login-form")
                .getText(".alert-warning").then(expectErrors)
                .call(done);
        });
        it('rejects too small', function(done) {
            browser
                .url("/auth/login")
                .setValue('#login-email', "good@email.com")
                .setValue('#login-password', "aa")
                .submitForm("#login-form")
                .getText(".alert-warning").then(expectErrors)
                .call(done);
        });
        it('rejects an invalid format', function(done) {
            browser
                .url("/auth/login")
                .setValue('#login-email', "good@email.com")
                .setValue('#login-password', "-12=./'[]=-1")
                .submitForm("#login-form")
                .getText(".alert-warning").then(expectErrors)
                .call(done);
        });
      });
  });
});
