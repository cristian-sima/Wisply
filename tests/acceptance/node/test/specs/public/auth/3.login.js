var userModule = require("../../general/user.js"),
  user = userModule.user,
  admin = {
    name: "Cristian Sima",
    email: "cristian.sima93@yahoo.com",
    password: "password"
  };

function expectErrors(text) {
  expect(text).toContain("Sorry... Your request was not successful");
}

describe('Login', function() {

  it('goes to the login page from menu', function(done) {
      browser
        .url("/")
        .windowHandleMaximize()
        .pause(1000)
        .click("a=Login")
        .pause(500)
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe('Login to Wisply');
        })
        .call(done);
  });
  it('submits an empty form and receives errors', function(done) {
      browser
        .submitForm("#login-form")
        .getText(".alert-warning").then(expectErrors)
        .call(done);
  });
  it('submits a form with just the email field completed', function(done) {
      browser
        .url("/auth/login")
        .setValue('#login-email', "admin@admin.com")
        .pause(1000)
        .submitForm("#login-form")
        .getText(".alert-warning").then(expectErrors)
        .call(done);
  });
  describe('Wrong data formats', function() {
      it('rejects a login form with wrong details', function(done) {
          browser
            .url("/auth/login")
            .setValue('#login-email', "wrong@wrong.com")
            .setValue('#login-password', "wrong")
            .pause(1000)
            .submitForm("#login-form")
            .getText(".alert-warning").then(expectErrors)
            .call(done);
      });
      it('rejects the SQL injection attack using the login form', function(done) {
          browser
            .url("/auth/login")
            .setValue('#login-email', "' OR 1=1 -- a")
            .setValue('#login-password', "' OR 1=1 -- a")
            .pause(1000)
            .submitForm("#login-form")
            .getText(".alert-warning").then(expectErrors)
            .call(done);
      });
      it('rejects the XSS attack using the login form', function(done) {
          browser
            .url("/auth/login")
            .setValue('#login-email', "<marquee>")
            .setValue('#login-password', "<marquee>")
            .pause(1000)
            .submitForm("#login-form")
            .getText(".alert-warning").then(expectErrors)
            .call(done);
      });
    describe('Rejects invalid email field', function() {
      it('rejects an email with a wrong format (wrongwrong.com)', function(done) {
          browser
            .url("/auth/login")
            .setValue('#login-email', "wrongwrong.com")
            .submitForm("#login-form")
            .getText(".alert-warning").then(function(text){
                expect(text).toContain("It must be a valid EMAIL type");
            })
            .call(done);
      });
      it('rejects an email field which is too short (a@a)', function(done) {
          browser
            .url("/auth/login")
            .setValue('#login-email', "a@a")
            .submitForm("#login-form")
            .getText(".alert-warning").then(expectErrors)
            .call(done);
      });
      it('rejects an email field which is valid, but too long (163 characters)', function(done) {
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
      describe('Password field', function() {
        it('rejects a password field which is valid, but too long (162 characters)', function(done) {
            browser
              .url("/auth/login")
              .setValue('#login-email', "good@email.com")
              .setValue('#login-password', "aaaasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasaasasasa.com")
              .submitForm("#login-form")
              .getText(".alert-warning").then(expectErrors)
              .call(done);
        });
        it('rejects an password field which is valid, but too short (aa)', function(done) {
            browser
              .url("/auth/login")
              .setValue('#login-email', "good@email.com")
              .setValue('#login-password', "aa")
              .submitForm("#login-form")
              .getText(".alert-warning").then(expectErrors)
              .call(done);
        });
        it("rejects a password which has an invalid format (-12=./'[]=-1)", function(done) {
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
  describe('Redirecting an account using the "sendMe" parameter', function() {
    it('redirects to home page a login request without sendMe paramter', function(done) {
        browser
          .url("/auth/login")
          .setValue('#login-email', user.email)
          .setValue('#login-password', user.password)
          .submitForm("#login-form")
          .pause(1000)
          .url(function(err, res){
              expect(res.value).toEqual("http://localhost:8081/");
          })
          .click('#menu-logout-button')
          .pause(3000)
          .call(done);
    });
    it('redirects to home page an account without privileges for a restricted area', function(done) {
        browser
          .url("/auth/login?sendMe=/admin")
          .pause(1000)
          .setValue('#login-email', user.email)
          .setValue('#login-password', user.password)
          .submitForm("#login-form")
          .pause(1500)
          .url(function(err, res){
              expect(res.value).toEqual("http://localhost:8081/");
          })
          .click('#menu-logout-button')
          .pause(3000)
          .call(done);
    });
    it('redirects to the home page when the value of sendMe is invalid (../../etc/pass)', function(done) {
        browser
          .url("/auth/login?sendMe=../../etc/pass")
          .setValue('#login-email', user.email)
          .setValue('#login-password', user.password)
          .submitForm("#login-form")
          .pause(1000)
          .url(function(err, res){
              expect(res.value).toEqual("http://localhost:8081/");
          })
          .click('#menu-logout-button')
          .pause(3000)
          .call(done);
    });
    it('redirects to the home page, when sendMe is an external address (http://google.com)', function(done) {
        browser
          .url("/auth/login?sendMe=http://google.com")
          .setValue('#login-email', user.email)
          .setValue('#login-password', user.password)
          .submitForm("#login-form")
          .pause(1000)
          .url(function(err, res){
              expect(res.value).toEqual("http://localhost:8081/");
          })
          .click('#menu-logout-button')
          .pause(3000)
          .call(done);
    });
    it('redirects to restricted area an account with enough privileges and a valid "sendMe" paramter)', function(done) {
        browser
          .url("/auth/login?sendMe=/admin")
          .setValue('#login-email', admin.email)
          .setValue('#login-password', admin.password)
          .submitForm("#login-form")
          .pause(1000)
          .url(function(err, res){
              expect(res.value).toEqual("http://localhost:8081/admin");
          })
          .click('#menu-logout-button')
          .pause(3000)
          .call(done);
    });
  });
});
