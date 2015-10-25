var badUser = {
  name: "Hacker",
  password: "HackerPassword"
},
user = {
  name: "Jameson Henry",
  email: "henry@oxford.ac.uk",
  password: "my-strong-password"
},
attack = {
  sql : "' OR 1=1 -- a",
  xss: "<marquee>"
};

function expectErrors(text) {
  expect(text).toContain("Sorry... Your request was not successful");
}

function expectSuccess(text) {
  expect(text).toContain("Success! Your account is ready!");
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

    // creates an ccount
    describe('Create account', function() {
      it('completes the register form and sends it', function(done) {
          browser
              .url("/auth/register")
              .setValue('#register-name', user.name)
              .setValue('#register-password', user.password)
              .setValue('#register-password-confirm', user.password)
              .setValue('#register-email', user.email)
              .submitForm("#register-form")
              .pause(1000)
              .isExisting(".alert-success").then(function(isExisting){
                expect(isExisting).toBe(true);
              })
              .call(done);
      });
      it('goes back to the login page', function(done) {
          browser
          .click("a=Go back")
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe('Login to Wisply');
          })
          .call(done);
      });
      it('connects the account', function(done) {
          browser
          .setValue('#login-email', user.email)
          .setValue('#login-password', user.password)
          .submitForm("#login-form")
          .pause(1000)
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe('Wisply - Building the hive of education');
          })
          .call(done);
      });
      it('checks if the account is connected', function(done) {
          browser
          .isExisting('#menu-logout-button').then(function(isExisting){
              expect(isExisting).toBe(true);
          })
          .call(done);
      });
      it('disconnects the account', function(done) {
          browser
          .click('#menu-logout-button')
          .pause(3500)
          .isExisting('#menu-logout-button').then(function(isExisting){
              expect(isExisting).toBeFalsy();
          })
          .call(done);
      });
    });


    //
    describe('rejects wrong data formats', function() {
          it('rejects SQL injection attack', function(done) {
              browser
                  .url("/auth/register")
                  .setValue('#register-name', attack.sql)
                  .setValue('#register-password', attack.sql)
                  .setValue('#register-password-confirm', attack.sql)
                  .setValue('#register-email', attack.sql)
                  .submitForm("#register-form")
                  .pause(1000)
                  .getText(".alert-warning").then(expectErrors)
                  .call(done);
          });
          it('rejects XSS attack', function(done) {
              browser
                  .url("/auth/register")
                  .setValue('#register-name', attack.xss)
                  .setValue('#register-password', attack.xss)
                  .setValue('#register-password-confirm', attack.xss)
                  .setValue('#register-email', attack.xss)
                  .submitForm("#register-form")
                  .pause(1000)
                  .getText(".alert-warning").then(expectErrors)
                  .call(done);
          });
          it('rejects empty form', function(done) {
              browser
                  .url("/auth/register")
                  .submitForm("#register-form")
                  .pause(1000)
                  .getText(".alert-warning").then(expectErrors)
                  .call(done);
          });

        describe('rejects wrong name field', function() {
           it('rejects a form without name', function(done) {
               browser
                 .url("/auth/register")
                 .setValue('#register-password', user.password)
                 .setValue('#register-password-confirm', user.password)
                 .setValue('#register-email', user.email)
                 .submitForm("#register-form")
                 .pause(1000)
                 .getText(".alert-warning").then(expectErrors)
                 .call(done);
           });
           it('rejects an invalid name', function(done) {
               browser
                 .url("/auth/register")
                 .setValue('#register-name', "<>-348_23mnsn0")
                 .setValue('#register-password', user.password)
                 .setValue('#register-password-confirm', user.password)
                 .setValue('#register-email', user.email)
                 .submitForm("#register-form")
                 .pause(1000)
                 .getText(".alert-warning").then(expectErrors)
                 .call(done);
           });
           it('rejects a name too short (2 characters)', function(done) {
               browser
                 .url("/auth/register")
                 .setValue('#register-name', "aa")
                 .setValue('#register-password', user.password)
                 .setValue('#register-password-confirm', user.password)
                 .setValue('#register-email', user.email)
                 .submitForm("#register-form")
                 .pause(1000)
                 .getText(".alert-warning").then(expectErrors)
                 .call(done);
           });
           it('rejects a too big name (240 characters)', function(done) {
               browser
                 .url("/auth/register")
                 .setValue('#register-name', "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
                 .setValue('#register-password', user.password)
                 .setValue('#register-password-confirm', user.password)
                 .setValue('#register-email', user.email)
                 .submitForm("#register-form")
                 .pause(1000)
                 .getText(".alert-warning").then(expectErrors)
                 .call(done);
           });
         });

         describe('rejects invalid password confirmation field', function() {
           it('rejects a form without the password confirmation', function(done) {
               browser
                 .url("/auth/register")
                 .setValue('#register-name', user.name)
                 .setValue('#register-password', user.password)
                 .setValue('#register-email', user.email)
                 .submitForm("#register-form")
                 .pause(1000)
                 .isExisting("h4=I got a problem").then(function(element){
                   expect(element).toBe(true);
                 })
                 .call(done);
           });
           it('rejects a form which has the confirmation password field different from password', function(done) {
               browser
                 .url("/auth/register")
                 .setValue('#register-name', user.name)
                 .setValue('#register-password', user.password)
                 .setValue('#register-password-confirm', "different")
                 .setValue('#register-email', user.email)
                 .submitForm("#register-form")
                 .pause(1000)
                 .isExisting("h4=I got a problem").then(function(element){
                   expect(element).toBe(true);
                 })
                 .call(done);
           });
         });
           describe('rejects invalid email fields', function() {
              it('rejects an empty email field', function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-password', user.password)
                    .setValue('#register-password-confirm', user.password)
                    .setValue('#register-name', user.name)
                    .submitForm("#register-form")
                    .pause(1000)
                    .getText(".alert-warning").then(expectErrors)
                    .call(done);
              });
              it('rejects an email too short (2 characters)', function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-password', user.password)
                    .setValue('#register-password-confirm', user.password)
                    .setValue('#register-name', user.name)
                    .setValue('#register-email', "aa")
                    .submitForm("#register-form")
                    .pause(1000)
                    .getText(".alert-warning").then(expectErrors)
                    .call(done);
              });
              it('rejects a valid email too long (148 characters)', function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-password', user.password)
                    .setValue('#register-password-confirm', user.password)
                    .setValue('#register-name', user.name)
                    .setValue('#register-email', "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.aaaa")
                    .submitForm("#register-form")
                    .pause(1000)
                    .getText(".alert-warning").then(expectErrors)
                    .call(done);
              });
              it('rejects an email which has invalid format', function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-password', user.password)
                    .setValue('#register-password-confirm', user.password)
                    .setValue('#register-name', user.name)
                    .setValue('#register-email', attack.sql)
                    .submitForm("#register-form")
                    .pause(1000)
                    .getText(".alert-warning").then(expectErrors)
                    .call(done);
              });

              it('rejects an email which has been already used', function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-name', user.name)
                    .setValue('#register-password', user.password)
                    .setValue('#register-password-confirm', user.password)
                    .setValue('#register-email', user.email)
                    .submitForm("#register-form")
                    .pause(1000)
                    .isExisting(".alert-warning").then(function(isExisting){
                      expect(isExisting).toBe(true);
                    })
                    .call(done);
              });
           });
           describe('rejects the form with invalid password field', function() {
              it('rejects the form which does not contain the password field', function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-name', user.name)
                    .setValue('#register-email', user.email)
                    .submitForm("#register-form")
                    .pause(1000)
                    .getText(".alert-warning").then(expectErrors)
                    .call(done);
              });
              it('rejects the form with a password too short (3 characters)', function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-name', user.name)
                    .setValue('#register-password', "pass")
                    .setValue('#register-password-confirm', "pass")
                    .setValue('#register-email', user.email)
                    .submitForm("#register-form")
                    .pause(1000)
                    .getText(".alert-warning").then(expectErrors)
                    .call(done);
              });
              it('rejects a form which contains a valid password, but too long (164 characters)', function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-name', user.name)
                    .setValue('#register-password', "passpasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspass")
                    .setValue('#register-password-confirm', "passpasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspasspass")
                    .setValue('#register-email', user.email)
                    .submitForm("#register-form")
                    .pause(1000)
                    .getText(".alert-warning").then(expectErrors)
                    .call(done);
              });
              it("rejects a form which does not contain a valid password format (-'1/.,`0=23-4`)", function(done) {
                  browser
                    .url("/auth/register")
                    .setValue('#register-name', user.name)
                    .setValue('#register-password', "-'1/.,`0=23-4`'")
                    .setValue('#register-password-confirm', "-'1/.,`0=23-4`'")
                    .setValue('#register-email', user.email)
                    .submitForm("#register-form")
                    .pause(1000)
                    .getText(".alert-warning").then(expectErrors)
                    .call(done);
              });
          });
      });
});
