
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

/**
 * It splits a connection cookie and returns the id and the token
 * @param {string} string The cookie in the string format
 * @return {object} The elements of the connection cookie
 */
function splitCookie(string) {
  var array = string.split(settings.separator),
    obj = {
    id: array[0],
    token: array[1],
  };
  return obj;
}

/**
 * It creates the value for the connection cookie
 * @param  {object} cookie An object which contains the id and the token of the cookie
 * @return {string} The string for the connection cookie
 */
function createCookie(cookie) {
  var string =  cookie.id + settings.separator + cookie.token;
  return string;
}


browser.addCommand("isAccountConnected", function() {
    return this.isExisting('#menu-logout-button');
});
browser.addCommand("isShortConnection", function() {
  return this.isExisting('#menu-logout-button').then(function(isConnected) {
         return this.getCookie("session").then(function(session) {
            return this.getCookie("connection").then(function(connectionCookie) {
                var isShort = false;
                if(session && !connectionCookie) {
                  isShort = true;
                }
                return isConnected && isShort;
            });
         });
     });
});
browser.addCommand("isLongConnection", function() {
  return this.isExisting('#menu-logout-button').then(function(isConnected) {
         return this.getCookie("connection").then(function(connectionCookie) {
          var isLong = false;
          if(connectionCookie) {
            isLong = true;
          }
          return isConnected && isLong;
         });
     });
});

// creates an ccount
describe('Cookies', function() {

    describe('Short connection', function() {
      it('disconnects any account', function(done) {
        browser
        .pause(1500)
        .url("/")
        .isExisting('#menu-logout-button').then(function(isExisting){
          console.log("Exista butonul");
            if(isExisting) {
                this.pause(1000);
                this.click("#menu-logout-button");
                this.pause(3500);
            }
        })
        .call(done);
      });
      it('connects and disconnects', function(done) {
        browser
        .url("/auth/login")
        .pause(1500)
        .setValue('#login-email', user.email)
        .setValue('#login-password', user.password)
        .submitForm("#login-form")
        .isShortConnection(function(err, result){
          expect(result).toBe(true);
        })
        .click('#menu-logout-button')
        .pause(3000)
        .call(done);
      });
      it('rejects the modification of session cookie in order to connect', function(done) {
        browser
        .url("/auth/login")
        .deleteCookie("session")
        .setCookie({
          name: "session",
          value: "340cbe187716d64c316g3cf7f7a10193",
        })
        .refresh()
        .isAccountConnected().then(function(err, isAccountConnected){
            expect(isAccountConnected).toBeFalsy();
        })
        .call(done);
      });
      it('rejects the modification of session cookie when the account has been connected', function(done) {
        browser
        .url("/auth/login")
        .setValue('#login-email', user.email)
        .setValue('#login-password', user.password)
        .submitForm("#login-form")
        .isShortConnection(function(err, result){
          expect(result).toBe(true);
        })
        .setCookie({
          name: "session",
          value: "340cbe187716d64c316g3cf7f7a10193",
        })
        .refresh()
        .isAccountConnected().then(function(err, isAccountConnected){
            expect(isAccountConnected).toBeFalsy();
        })
        .call(done);
      });
      it('disconnects the account if there is no session cookie', function(done) {
        browser
        .url("/auth/login")
        .setValue('#login-email', user.email)
        .setValue('#login-password', user.password)
        .submitForm("#login-form")
        .isShortConnection(function(err, result){
          expect(result).toBe(true);
        })
        .deleteCookie("session")
        .refresh()
        .isAccountConnected().then(function(err, isAccountConnected){
            expect(isAccountConnected).toBeFalsy();
        })
        .call(done);
      });
    });
    describe('Long term connection', function() {
      it('connects the account and disconnects it', function(done) {
        browser
        .url("/auth/login")
        .pause(1000)
        .setValue('#login-email', user.email)
        .setValue('#login-password', user.password)
        .click('#login-remember-me')
        .submitForm("#login-form")
        .isLongConnection(function(err, result){
          expect(result).toBe(true);
        })
        .click('#menu-logout-button')
        .pause(3500)
        .call(done);
      });
      it('rejects if the connection cookie is added by attacker', function(done) {
        browser
        .url("/auth/login")
        .deleteCookie("connection")
        .pause(200)
        .setCookie({
          name: "connection",
          value: "15::w8vkgcawi88xlrwzTJbKJ7bbun8=",
        })
        .url("/auth/login")
        .pause(500)
        .isLongConnection(function(err, result){
          expect(result).toBeFalsy();
        })
        .call(done);
      });
        describe('Modifies the cookie', function() {
          it('rejects the modification of id', function(done) {
            browser
            .url("/auth/login")
            .pause(1000)
            .setValue('#login-email', user.email)
            .setValue('#login-password', user.password)
            .click('#login-remember-me')
            .submitForm("#login-form")
            .isLongConnection(function(err, result){
              expect(result).toBe(true);
            })
            .getCookie("connection").then(function(cookie){
                connectionCookie = splitCookie(cookie.value);
                hacked = createCookie({
                  id: "666",
                  token: connectionCookie.token,
                });
                this.setCookie({
                  name: "connection",
                  value: hacked
                });
            })
            .deleteCookie("session")
            .url("/auth/login")
            .pause(500)
            .isLongConnection(function(err, result){
              expect(result).toBeFalsy();
            })
            .call(done);
          });
          it('rejects the modification of token', function(done) {
            browser
            .url("/auth/login")
            .pause(1000)
            .setValue('#login-email', user.email)
            .setValue('#login-password', user.password)
            .click('#login-remember-me')
            .submitForm("#login-form")
            .isLongConnection(function(err, result){
              expect(result).toBe(true);
            })
            .getCookie("connection").then(function(cookie){
                connectionCookie = splitCookie(cookie.value);
                hacked = createCookie({
                  id: connectionCookie.id,
                  token: "-AZ8DP995S1ecA_99YQirr_wULs=",
                });
                this.setCookie({
                  name: "connection",
                  value: hacked
                });
            })
            .deleteCookie("session")
            .url("/auth/login")
            .pause(500)
            .isLongConnection(function(err, result){
              expect(result).toBeFalsy();
            })
            .call(done);
          });
        });
        describe('Delete', function() {
          it('rejects the connection if the cookies are deleted', function(done) {
            browser
            .url("/auth/login")
            .pause(1000)
            .setValue('#login-email', user.email)
            .setValue('#login-password', user.password)
            .click('#login-remember-me')
            .submitForm("#login-form")
            .isLongConnection(function(err, result){
              expect(result).toBe(true);
            })
            .deleteCookie("connection")
            .deleteCookie("session")
            .url("/auth/login")
            .pause(500)
            .isLongConnection(function(err, result){
              expect(result).toBeFalsy();
            })
            .call(done);
          });
          it('keeps connected if the connection session cookie is deleted', function(done) {
            browser
            .url("/auth/login")
            .pause(1000)
            .setValue('#login-email', user.email)
            .setValue('#login-password', user.password)
            .click('#login-remember-me')
            .submitForm("#login-form")
            .isLongConnection(function(err, result){
              expect(result).toBe(true);
            })
            .deleteCookie("session")
            .url("/auth/login")
            .pause(500)
            .isLongConnection(function(err, result){
              expect(result).toBe(true);
            })
            .click('#menu-logout-button')
            .pause(3500)
            .call(done);
          });
        });
    });
});
