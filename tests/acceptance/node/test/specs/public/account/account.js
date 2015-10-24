
var  user = {
      name: "Jameson Henry",
      email: "henry@oxford.ac.uk",
      password: "my-strong-password"
    },
    settings = {
      separator: "::"
    },
     connectionCookie = {},
   hackedCookie = "";

// creates an ccount
describe('Account', function() {

      it('connect', function(done) {
        browser
        .url("/auth/login")
        .deleteCookie("session")
        .deleteCookie("connection")
        .pause(1000)
        .url("/auth/login")
        .pause(1500)
        .setValue('#login-email', user.email)
        .setValue('#login-password', user.password)
        .submitForm("#login-form")
        .pause(3000)
        .call(done);
      });

      it('go to dashboard', function(done){
        browser.
            url('/account')
            .pause(1000)
            .getTitle(function(err, title) {
                expect(err).toBe(undefined);
                expect(title).toBe("Account - Dashboard");
            })
            .call(done);
      });

      it('go to searches', function(done){
        browser.
            url('/account/search')
            .pause(1000)
            .getTitle(function(err, title) {
                expect(err).toBe(undefined);
                expect(title).toBe("Account - Activity");
            })
            .call(done);
      });

    });
