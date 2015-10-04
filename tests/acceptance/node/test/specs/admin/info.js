
var account = {
  name : "Cristian Sima",
  email: "cristian.sima93@yahoo.com",
  password: "password",
},
dashboard =  {
  title: "Admin - Dashboard"
};

exports.info = {
    account: account,
    dashboard: dashboard,
};

exports.init = function (browser) {
  it('disconnects the current account', function (done) {
    browser
        .windowHandleMaximize()
        .url("/")
        .pause(1000)
        .isExisting('#menu-logout-button').then(function(isExisting){
            if(isExisting) {
                this.click("#menu-logout-button");
                this.pause(3000);
            }
        })
        .call(done);
  });
  it('connects as Administrator', function (done) {
    browser
        .pause(1000)
        .url("/auth/login?sendMe=/admin")
        .setValue('#login-email', account.email)
        .setValue('#login-password', account.password)
        .click('#login-remember-me')
        .pause(1000)
        .submitForm("#login-form")
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe(dashboard.title);
        })
        .call(done);
  });
};
